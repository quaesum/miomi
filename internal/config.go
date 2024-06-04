package internal

import (
	"log"
	environment "madmax/sdk/environment/config"
)

type Config struct {
	MysqlDSN string `yaml:"mysql_dsn" mapstructure:"MYSQL_DSN"`
	NatsUrl  string `json:"NATS_URL" mapstructure:"NATS_URL"`
}

func NewConfig() (*Config, error) {

	v, err := environment.LoadConfig(".env")
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg, nil
}
