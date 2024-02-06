package config

import (
	interceptorconf "github.com/dre-zouhair/interceptor/config"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

type Config struct {
	ServerPort               string
	ProtectionMiddlewareConf interceptorconf.ProtectionMiddlewareConfig
}

func LoadConfig() (*Config, error) {

	port := os.Getenv("CARTAPI_SERVER_PORT")

	if port == "" {
		log.Error().Str("port", "5050").Msg("missing env variable, using default port value")
		port = "5050"
	}

	middlewareConf := interceptorconf.ProtectionMiddlewareConfig{
		ProcessorConfig: interceptorconf.ProcessorConfig{
			CustomHeaderSignals: strings.Split(os.Getenv("INTERCEPTOR_PROTECTION_CUSTOM_HEADERS"), ","),
			CustomHeaderCookies: strings.Split(os.Getenv("INTERCEPTOR_PROTECTION_CUSTOM_COOKIES"), ","),
		},
		ProtectionAPIConfig: interceptorconf.ProtectionAPIConfig{
			ProtectionEndpoint: os.Getenv("INTERCEPTOR_PROTECTION_ENDPOINT"),
			ProtectionToken:    os.Getenv("INTERCEPTOR_PROTECTION_TOKEN"),
		},
	}

	return &Config{
		ServerPort:               port,
		ProtectionMiddlewareConf: middlewareConf,
	}, nil
}
