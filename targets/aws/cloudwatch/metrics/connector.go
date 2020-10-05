package metrics

import "github.com/kubemq-hub/builder/common"

// TODO
func Connector() *common.Connector {
	return common.NewConnector().
		SetKind("target.aws.cloudwatch.metrics").
		SetDescription("AWS Cloudwatch Metrics Target")
	//AddProperty(
	//	common.NewProperty().
	//		SetKind("string").
	//		SetName("address").
	//		SetDescription("Sets Kubemq grpc endpoint address").
	//		SetMust(true).
	//		SetDefault("localhost:50000"),
	//).
}