package activemq

import (
	"github.com/kubemq-hub/builder/connector/common"
)

func Connector() *common.Connector {
	return common.NewConnector().
		SetKind("messaging.activemq").
		SetDescription("ActiveMQ Messaging Target").
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("host").
				SetDescription("Set ActiveMQ host address").
				SetMust(true).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("username").
				SetDescription("Set ActiveMQ username").
				SetMust(false).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("password").
				SetDescription("Set ActiveMQ password").
				SetMust(false).
				SetDefault(""),
		).
		AddMetadata(
			common.NewMetadata().
				SetName("destination").
				SetKind("string").
				SetDescription("Set ActiveMQ destination").
				SetDefault("").
				SetMust(true),
		)
}
