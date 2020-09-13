package eventhubs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Azure/azure-event-hubs-go/v3"
	"github.com/kubemq-hub/kubemq-targets/config"
	"github.com/kubemq-hub/kubemq-targets/types"
)

type Client struct {
	name   string
	opts   options
	client *eventhub.Hub
}

func New() *Client {
	return &Client{}

}
func (c *Client) Name() string {
	return c.name
}

func (c *Client) Init(ctx context.Context, cfg config.Spec) error {
	c.name = cfg.Name
	var err error
	c.opts, err = parseOptions(cfg)
	if err != nil {
		return err
	}
	c.client, err = eventhub.NewHubFromConnectionString(c.opts.connectionString)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Do(ctx context.Context, req *types.Request) (*types.Response, error) {
	meta, err := parseMetadata(req.Metadata)
	if err != nil {
		return nil, err
	}
	switch meta.method {
	case "send":
		return c.send(ctx, meta, req.Data)
	case "send_batch":
		return c.send(ctx, meta, req.Data)
	}
	return nil, errors.New("invalid method type")
}

func (c *Client) send(ctx context.Context, meta metadata, data []byte) (*types.Response, error) {
	event := &eventhub.Event{
		Data: data,
	}
	if meta.partitionKey != "" {
		event.PartitionKey = &meta.partitionKey
	}
	if meta.properties != "" {
		p, err := json.Marshal(meta.properties)
		if err != nil {
			return nil, err
		}
		properties := make(map[string]interface{})
		err = json.Unmarshal(p, properties)
		if err != nil {
			return nil, err
		}
		event.Properties = properties
	}
	err := c.client.Send(ctx, event)
	if err != nil {
		return nil, err
	}
	return types.NewResponse().
			SetMetadataKeyValue("result", "ok"),
		nil
}

func (c *Client) sendBatch(ctx context.Context, meta metadata, data []byte) (*types.Response, error) {
	var messages []string
	err := json.Unmarshal(data, messages)
	if err != nil {
		return nil, err
	}
	var events []*eventhub.Event
	for _, message := range messages {
		event := eventhub.NewEventFromString(message)
		events = append(events, event)
		if meta.partitionKey != "" {
			event.PartitionKey = &meta.partitionKey
		}
	}
	if meta.properties != "" {
		properties := make(map[string]interface{})
		p, err := json.Marshal(meta.properties)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(p, properties)
		if err != nil {
			return nil, err
		}
		for _, event := range events {
			event.Properties = properties
		}
	}

	err = c.client.SendBatch(ctx, eventhub.NewEventBatchIterator(events...))
	if err != nil {
		return nil, err
	}
	return types.NewResponse().
			SetMetadataKeyValue("result", "ok"),
		nil
}
