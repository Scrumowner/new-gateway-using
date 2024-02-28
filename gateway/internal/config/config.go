package config

import "os"

type Config struct {
	Port           string
	UserService    *UserService
	FinanceService *FinanceService
}
type UserService struct {
	Host string
	Port string
}
type FinanceService struct {
	Host string
	Port string
}

func NewConfig() *Config {
	return &Config{}
}
func (c *Config) Load() {
	c.Port = os.Getenv("GATEWAY_PORT")
	c.UserService = &UserService{
		Host: os.Getenv("USER_HOST"),
		Port: os.Getenv("USER_SERVICE_PORT"),
	}
	c.FinanceService = &FinanceService{
		Host: os.Getenv("FINANCE_HOST"),
		Port: os.Getenv("FINANCE_SERVICE_PORT"),
	}
}
