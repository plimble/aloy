package config

import "github.com/plimble/goconfig"

var (
	config *Config
)

type Config struct {
	Addr           string `default:":4400" required:"true"`
	GoTestTags     string
	GithubUsername string
	GithubPassword string
	GitlabUsername string
	GitlabPassword string
	MaxQueue       int `default:":100" required:"true"`
	MaxRunner      int `default:":5" required:"true"`
}

func Get() *Config {
	config := &Config{}
	goconfig.Process("aloy", config)

	return config
}