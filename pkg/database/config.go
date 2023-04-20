package database

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Hostname string
	Port     int
	User     string
	Password string
	Name     string
}

func (c Config) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		c.Hostname,
		c.User,
		c.Password,
		c.Name,
		c.Port,
	)
}

func LoadConfig() (Config, error) {
	config := Config{}
	err := envconfig.Process("DATABASE", &config)

	return config, err
}
