package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	Environment string
	Port        string
}

var baseConfig = Configuration{
	Environment: "local",
	Port:        ":8080",
}

func Get() Configuration {
	loadedConfig := baseConfig
	envconfig.MustProcess("STORAGE", &loadedConfig)

	if loadedConfig.Environment == "local" {
		err := godotenv.Load()
		if err == nil {
			envconfig.MustProcess("STORAGE", &loadedConfig)
		}
	}
	return loadedConfig
}
