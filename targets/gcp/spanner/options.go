package spanner

import (
	"fmt"
	"github.com/kubemq-hub/kubemq-target-connectors/config"
)

type options struct {
	db string
}

func parseOptions(cfg config.Metadata) (options, error) {
	o := options{}
	var err error
	o.db , err = cfg.MustParseString("db")
	if err != nil {
		return options{}, fmt.Errorf("error parsing db, %w", err)
	}
	err = config.MustExistsEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if err !=nil{
		return options{}, err
	}
	return o, nil
}