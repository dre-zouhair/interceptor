package config

import (
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	ServerPort string
}

func LoadConfig() (*Config, error) {

	port := os.Getenv("CARTAPI_SERVER_PORT")

	if port == "" {
		log.Error().Str("port", "5050").Msg("missing env variable, using default port value")
		port = "5050"
	}

	return &Config{
		ServerPort: port,
	}, nil
}
