package ibmmq

import (
	"context"
	"fmt"
	"github.com/ibm-messaging/mq-golang-jms20/jms20subset"
	"github.com/ibm-messaging/mq-golang-jms20/mqjms"
	"github.com/kubemq-hub/kubemq-targets/config"
	"github.com/kubemq-hub/kubemq-targets/pkg/logger"
	"github.com/kubemq-hub/kubemq-targets/types"
)

type Client struct {
	name       string
	opts       options
	queue      jms20subset.Queue
	jmsContext jms20subset.JMSContext
	log        *logger.Logger
	producer   jms20subset.JMSProducer
}

func New() *Client {
	return &Client{}

}

//func (c *Client) Connector() *common.Connector {
//	return Connector()
//}
func (c *Client) Init(ctx context.Context, cfg config.Spec) error {
	c.name = cfg.Name
	c.log = logger.NewLogger(cfg.Name)
	var err error
	c.opts, err = parseOptions(cfg)
	if err != nil {
		return err
	}
	cf := mqjms.ConnectionFactoryImpl{
		QMName:           c.opts.qMName,
		Hostname:         c.opts.hostname,
		PortNumber:       c.opts.portNumber,
		ChannelName:      c.opts.channelName,
		UserName:         c.opts.userName,
		TransportType:    mqjms.TransportType_CLIENT,
		TLSClientAuth:    mqjms.TLSClientAuth_NONE,
		KeyRepository:    c.opts.keyRepository,
		Password:         c.opts.Password,
		CertificateLabel: c.opts.certificateLabel,
	}

	jmsContext, jmsErr := cf.CreateContext()
	if jmsErr != nil {
		return fmt.Errorf("failed to create context on error %s", jmsErr.GetReason())
	}
	c.jmsContext = jmsContext
	c.queue = jmsContext.CreateQueue(c.opts.queueName)
	c.producer = c.jmsContext.CreateProducer().SetDeliveryMode(c.opts.deliveryMode).SetTimeToLive(c.opts.timeToLive)
	return nil
}

func (c *Client) Do(ctx context.Context, req *types.Request) (*types.Response, error) {
	jmsErr := c.producer.SendString(c.queue, fmt.Sprintf("%s", req.Data))
	if jmsErr != nil {
		return nil, fmt.Errorf("failed to create context on error %s", jmsErr.GetReason())
	}

	return types.NewResponse().
			SetMetadataKeyValue("result", "ok"),
		nil
}

func (c *Client) Stop() error {
	if c.jmsContext != nil {
		c.jmsContext.Close()
	}
	return nil
}
