package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env           string `env:"ENV" envDefault:"development"`
	Timezone      string `env:"TZ" envDefault:"UTC"`
	Log           LogConfig
	Db            DatabaseConfig
	Api           ApiConfig
	MessageBroker MessageBrokerConfig
}

type LogConfig struct {
	Level    string `env:"TFR_LOGGER_LEVEL" envDefault:"info"`
	Format   string `env:"TFR_LOGGER_FORMAT" envDefault:"json"`
	Output   string `env:"TFR_LOGGER_OUTPUT" envDefault:"stdout"`
	FilePath string `env:"TFR_LOGGER_PATH" envDefault:"./storage/logs/transfer.log"`
}

type DatabaseConfig struct {
	Host     string `env:"TFR_PG_HOST" envDefault:"localhost"`
	Port     string `env:"TFR_PG_PORT" envDefault:"5432"`
	User     string `env:"TFR_PG_USER" envDefault:"username"`
	Password string `env:"TFR_PG_PASSWORD" envDefault:"password"`
	Database string `env:"TFR_PG_DATABASE" envDefault:"transfer-db"`
	SslMode  string `env:"TFR_PG_SSL_MODE" envDefault:"disable"`
}

type ApiConfig struct {
	Host string `env:"TFR_API_HOST" envDefault:"localhost"`
	Port string `env:"TFR_API_PORT" envDefault:"8081"`
}

type MessageBrokerConfig struct {
	Url string `env:"TFR_MB_URL"`
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env:      os.Getenv("ENV"),
		Timezone: os.Getenv("TZ"),
		Log: LogConfig{
			Level:    os.Getenv("TFR_LOGGER_LEVEL"),
			Format:   os.Getenv("TFR_LOGGER_FORMAT"),
			Output:   os.Getenv("TFR_LOGGER_OUTPUT"),
			FilePath: os.Getenv("TFR_LOGGER_PATH"),
		},
		Db: DatabaseConfig{
			Host:     os.Getenv("TFR_PG_HOST"),
			Port:     os.Getenv("TFR_PG_PORT"),
			User:     os.Getenv("TFR_PG_USER"),
			Password: os.Getenv("TFR_PG_PASSWORD"),
			Database: os.Getenv("TFR_PG_DATABASE"),
			SslMode:  os.Getenv("TFR_PG_SSL_MODE"),
		},
		Api: ApiConfig{
			Host: os.Getenv("TFR_API_HOST"),
			Port: os.Getenv("TFR_API_PORT"),
		},
		MessageBroker: MessageBrokerConfig{
			Url: os.Getenv("TFR_MB_URL"),
		},
	}

	return cfg, nil
}
