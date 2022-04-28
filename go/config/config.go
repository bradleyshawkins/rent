package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

func Load[T any]() (*T, error) {
	var c T
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("unable to process config. Error:%v", err)
	}

	// LoadSecrets
	return &c, nil
}
