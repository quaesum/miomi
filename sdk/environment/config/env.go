package environment

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (*viper.Viper, error) {
	// Define the configuration file flag with a default location and help description
	//var configPath string
	//flag.StringVar(&configPath, "config", "./config/config-prod.json", "Path to config file")
	//flag.Parse()

	// Create a new Viper instance
	v := viper.New()

	// Set configuration file path from the flag
	v.SetConfigFile(path)

	// Enable automatic environment variables handling
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Replace dots with underscores in env vars

	// Attempt to read the configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("LoadConfig: unable to read config file at %s: %w", path, err)
	}

	// Logging the effective configuration file used
	fmt.Println("Using config file:", v.ConfigFileUsed())

	return v, nil
}
