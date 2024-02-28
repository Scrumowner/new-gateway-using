package config

import "os"

type Config struct {
	Port string
	Key  string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() {
	c.Port = os.Getenv("FINANCE_SERVICE_PORT")
	c.Key = os.Getenv("API_KEY")
}
