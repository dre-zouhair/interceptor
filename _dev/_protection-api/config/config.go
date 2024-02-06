package config

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

type Config struct {
	ServerPort string
	ApiTokens  []string
}

func LoadConfig() (*Config, error) {

	port := os.Getenv("PROTECTIONAPI_SERVER_PORT")

	if port == "" {
		log.Error().Str("port", "5789").Msg("missing env variable, using default port value")
		port = "5789"
	}

	tokens := os.Getenv("PROTECTIONAPI_TOKENS")

	return &Config{
		ServerPort: port,
		ApiTokens:  strings.Split(tokens, ","),
	}, nil
}
