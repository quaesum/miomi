package internal

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	MysqlDSN string `yaml:"mysql_dsn"`
}

func NewConfig() (*Config, error) {

	fileName := "./config/config.yaml"
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w, you can specify path to the config file via CONFIG_PATH variable", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse envrionment: %w", err)
	}

	return &cfg, nil
}
