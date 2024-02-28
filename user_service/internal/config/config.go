package config

import "os"

type Config struct {
	Port   string
	Secret string
	*DBConfig
}
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() {
	c.Port = os.Getenv("USER_SERVICE_PORT")
	c.Secret = os.Getenv("SECRET")
	dbconfig := &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
	}
	c.DBConfig = dbconfig
}
