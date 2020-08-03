package activemq

import (
	"fmt"
	"github.com/kubemq-hub/kubemq-targets/config"
)

type options struct {
	host     string
	username string
	password string
}

func parseOptions(cfg config.Spec) (options, error) {
	o := options{}
	var err error
	o.host, err = cfg.MustParseString("host")
	if err != nil {
		return options{}, fmt.Errorf("error parsing host, %w", err)
	}
	o.username = cfg.ParseString("username", "")
	o.password = cfg.ParseString("password", "")
	return o, nil
}
