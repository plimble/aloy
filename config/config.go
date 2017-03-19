package config

import "github.com/plimble/goconfig"

var (
	config *Config
)

// Config app
type Config struct {
	Addr              string `default:":4400" required:"true"`
	GoTestTags        string
	GithubAccessToken string
	GitlabAccessToken string
	MaxQueue          int `default:":100" required:"true"`
	MaxRunner         int `default:":5" required:"true"`
}

// Get config func
func Get() *Config {
	config := &Config{}
	goconfig.Process("aloy", config)

	return config
}
