package google

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kubemq-hub/kubemq-target-connectors/config"
	"github.com/kubemq-hub/kubemq-target-connectors/pkg/logger"
	"github.com/kubemq-hub/kubemq-target-connectors/targets"
	"github.com/kubemq-hub/kubemq-target-connectors/types"
	"google.golang.org/api/iterator"
)

type Client struct {
	name   string
	opts   options
	client *pubsub.Client
	log    *logger.Logger
	target targets.Target
}

func New() *Client {
	return &Client{}
	
}
func (c *Client) Name() string {
	return c.name
}

func (c *Client) Init(ctx context.Context, cfg config.Metadata) error {
	c.name = cfg.Name
	c.log = logger.NewLogger(cfg.Name)
	var err error
	c.opts, err = parseOptions(cfg)
	if err != nil {
		return err
	}

	
	
	client, err := pubsub.NewClient(ctx, c.opts.projectID)
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *Client) Do(ctx context.Context, request *types.Request) (*types.Response, error) {
	eventMetadata, err := parseMetadata(request.Metadata, c.opts)
	if err != nil {
		return nil, err
	}
	t := c.client.Topic(eventMetadata.topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: request.Data,
		Attributes: eventMetadata.tags,
	})
	tries := 0
	for tries <= c.opts.retries {
		id, err := result.Get(ctx)
		if err == nil {
			return types.NewResponse().
					SetMetadataKeyValue("event_id", fmt.Sprintf("%s", id)),
				nil
		}
		if tries >= c.opts.retries {
			return nil, err
		}
		tries++
	}
	return nil, fmt.Errorf("retries must be a zero or greater")
}

func (c *Client) list(ctx context.Context) (*types.Response, error) {
	var topics []string
	it := c.client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic.ID())
	}
	if len(topics)<=0 {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", "no topics found for this project"), nil
	}
	b,err:=json.Marshal(topics)
	if err != nil {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", err.Error()), nil
	}
	return types.NewResponse().
			SetData(b).
			SetMetadataKeyValue("result", "ok"),
		nil
}