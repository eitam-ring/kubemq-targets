package rabbitmq

import (
	"context"
	"fmt"
	"github.com/kubemq-hub/kubemq-targets/config"
	"github.com/kubemq-hub/kubemq-targets/types"
	"github.com/streadway/amqp"
)

type Client struct {
	name    string
	opts    options
	channel *amqp.Channel
}

func New() *Client {
	return &Client{}
}
func (c *Client) Connector() *common.Connector {
return Connector()
}
func (c *Client) Init(ctx context.Context, cfg config.Spec) error {
	c.name = cfg.Name
	var err error
	c.opts, err = parseOptions(cfg)
	if err != nil {
		return err
	}
	conn, err := amqp.Dial(c.opts.url)
	if err != nil {
		return fmt.Errorf("error dialing rabbitmq, %w", err)
	}
	c.channel, err = conn.Channel()
	if err != nil {
		return fmt.Errorf("error getting rabbitmq channel, %w", err)
	}
	return nil
}

func (c *Client) Do(ctx context.Context, req *types.Request) (*types.Response, error) {
	meta, err := parseMetadata(req.Metadata)
	if err != nil {
		return nil, err
	}
	return c.Publish(ctx, meta, req.Data)
}

func (c *Client) Publish(ctx context.Context, meta metadata, data []byte) (*types.Response, error) {
	msg := meta.amqpMessage(data)
	err := c.channel.Publish(meta.exchange, meta.queue, meta.mandatory, meta.immediate, msg)
	if err != nil {
		return nil, err
	}
	return types.NewResponse().
			SetMetadataKeyValue("result", "ok"),
		nil
}
