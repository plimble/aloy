package config

import "github.com/plimble/envconfig"

var (
	config *Config
)

type Config struct {
	Addr     string `default:":4400" required:"true"`
	TestTags []string
}

func Get() *Config {
	config := &Config{}
	envconfig.Process("aloy", config)

	return config
}
