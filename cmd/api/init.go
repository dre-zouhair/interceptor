//go:build !dev
// +build !dev

package main

import (
	"os"

	"github.com/dre-zouhair/interceptor/internal/config"
	"github.com/rs/zerolog/log"
)

func init() {

	conf, err := config.Build(config.RawConfig{
		ServerPort:          os.Getenv("INTERCEPTOR_SERVER_PORT"),
		ProtectionEndpoint:  os.Getenv("INTERCEPTOR_PROTECTION_ENDPOINT"),
		ProtectionToken:     os.Getenv("INTERCEPTOR_PROTECTION_TOKEN"),
		ForwardEndPoint:     os.Getenv("INTERCEPTOR_FORWARD_ENDPOINT"),
		CustomHeaderSignals: os.Getenv("INTERCEPTOR_PROTECTION_CUSTOM_HEADERS"),
		CustomHeaderCookies: os.Getenv("INTERCEPTOR_PROTECTION_CUSTOM_COOKIES"),
		ProtectionFailMode:  os.Getenv("INTERCEPTOR_PROTECTION_FAIL_MODE"),
	})

	if err != nil {
		log.Error().Err(err).Msg("invalid env configuration")
		return
	}

	appConf = conf
}
