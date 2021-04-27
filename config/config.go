package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ConnectionString string `envconfig:"DB_CONNECTIONSTRING" default:"postgresql://postgres:password@localhost:5432/rent"`
	MigrationPath    string `envconfig:"DB_MIGRATION_PATH" default:"./postgres/schema"`
}

func ParseConfig() (*Config, error) {
	log.Println("Parsing config")
	var c Config
	err := envconfig.Process("", &c)
	log.Printf("%+v\n", c)
	if err != nil {
		return nil, fmt.Errorf("unable to process config. Error:%v", err)
	}
	return &c, nil
}
