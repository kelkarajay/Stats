package database

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Hostname     string `json:"server"`
	Port         string `json:"port"`
	UserID       string `json:"userId"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
}

func (c Config) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		c.Hostname,
		c.UserID,
		c.Password,
		c.DatabaseName,
		c.Port,
	)
}

func LoadConfig() (Config, error) {
	config := Config{}
	err := envconfig.Process("DATABASE", &config)

	return config, err
}
