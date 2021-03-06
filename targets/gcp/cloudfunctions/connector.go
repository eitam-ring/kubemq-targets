package cloudfunctions

import (
	"github.com/kubemq-hub/builder/connector/common"
)

func Connector() *common.Connector {
	return common.NewConnector().
		SetKind("gcp.cloudfunctions").
		SetDescription("GCP Cloud Functions Target").
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("project_id").
				SetDescription("Set GCP project ID").
				SetMust(true).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("multilines").
				SetName("credentials").
				SetDescription("Set GCP credentials").
				SetMust(true).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("bool").
				SetName("location_match").
				SetDescription("Set Cloud Functions location match").
				SetMust(false).
				SetDefault("true"),
		).
		AddMetadata(
			common.NewMetadata().
				SetName("name").
				SetKind("string").
				SetDescription("Set Cloud Functions name").
				SetDefault("").
				SetMust(true),
		).
		AddMetadata(
			common.NewMetadata().
				SetName("project").
				SetKind("string").
				SetDescription("Set Cloud Functions project").
				SetDefault("").
				SetMust(false),
		).
		AddMetadata(
			common.NewMetadata().
				SetName("location").
				SetKind("string").
				SetDescription("Set Cloud Functions location").
				SetDefault("").
				SetMust(false),
		)
}
