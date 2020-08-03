package firestore

import (
	"fmt"
	"github.com/kubemq-hub/kubemq-targets/config"
)

type options struct {
	projectID string
	credentials string
}

func parseOptions(cfg config.Spec) (options, error) {
	o := options{}
	var err error
	o.projectID, err = cfg.MustParseString("project_id")
	if err != nil {
		return options{}, fmt.Errorf("error parsing project_id, %w", err)
	}
	o.credentials, err = cfg.MustParseString("credentials")
	if err != nil {
		return options{}, err
	}
	return o, nil
}
