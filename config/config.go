package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env           string
	Timezone      string
	Name          string
	Log           LogConfig
	Db            DatabaseConfig
	Api           ApiConfig
	MessageBroker MessageBrokerConfig
}

type LogConfig struct {
	Level    string
	Format   string
	Output   string
	FilePath string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SslMode  string
}

type ApiConfig struct {
	Address string
}

type MessageBrokerConfig struct {
	Url string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env:      os.Getenv("ENV"),
		Timezone: os.Getenv("TZ"),
		Name:     os.Getenv("NAME"),
		Log: LogConfig{
			Level:    os.Getenv("LOG_LEVEL"),
			Format:   os.Getenv("LOG_FORMAT"),
			Output:   os.Getenv("LOG_OUTPUT"),
			FilePath: os.Getenv("LOG_PATH"),
		},
		Db: DatabaseConfig{
			Host:     os.Getenv("PG_HOST"),
			Port:     os.Getenv("PG_PORT"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			Database: os.Getenv("PG_DATABASE"),
			SslMode:  os.Getenv("PG_SSL_MODE"),
		},
		Api: ApiConfig{
			Address: os.Getenv("API_ADDRESS"),
		},
		MessageBroker: MessageBrokerConfig{
			Url: os.Getenv("RMQ_URL"),
		},
	}

	return cfg, nil
}
