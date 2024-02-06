package config

import (
	"errors"
	"strings"

	configuration "github.com/dre-zouhair/interceptor/config"

	"github.com/dre-zouhair/interceptor/internal/utils"
	"github.com/rs/zerolog/log"
)

type RawConfig struct {
	ServerPort          string `mapstructure:"INTERCEPTOR_SERVER_PORT"`
	InterceptionPath    string `mapstructure:"INTERCEPTOR_INTERCEPTION_PATH"`
	ProtectionFailMode  string `mapstructure:"INTERCEPTOR_PROTECTION_FAIL_MODE"`
	ProtectionEndpoint  string `mapstructure:"INTERCEPTOR_PROTECTION_ENDPOINT"`
	ProtectionToken     string `mapstructure:"INTERCEPTOR_PROTECTION_TOKEN"`
	ForwardEndPoint     string `mapstructure:"INTERCEPTOR_FORWARD_ENDPOINT"`
	CustomHeaderSignals string `mapstructure:"INTERCEPTOR_PROTECTION_CUSTOM_HEADERS"`
	CustomHeaderCookies string `mapstructure:"INTERCEPTOR_PROTECTION_CUSTOM_COOKIES"`
}

func Build(rawConf RawConfig) (*configuration.Config, error) {

	_, err := utils.BuildURL(rawConf.ForwardEndPoint, "", "")
	if err != nil {
		log.Error().Err(err).Str("forwardEndPoint", rawConf.ForwardEndPoint).Msg("failed to parse forward target url")
		return nil, errors.New("invalid forward endpoint")
	}

	if rawConf.ServerPort == "" {
		log.Warn().Str("port", "7777").Msg("missing env variable, using default port value")
		rawConf.ServerPort = "7777"
	}

	if rawConf.ProtectionToken == "" {
		log.Warn().Msg("missing token env variable")
	}

	failMode := utils.ALLOW_ACCESS
	if strings.ToUpper(rawConf.ProtectionFailMode) == utils.BLOCK_ACCESS {
		failMode = utils.BLOCK_ACCESS
	}

	var CustomHeaderSignals []string
	if len(rawConf.CustomHeaderCookies) != 0 {
		CustomHeaderSignals = strings.Split(rawConf.CustomHeaderSignals, ",")
	}

	var CustomHeaderCookies []string
	if len(rawConf.CustomHeaderCookies) != 0 {
		CustomHeaderCookies = strings.Split(rawConf.CustomHeaderCookies, ",")
	}

	return &configuration.Config{
		ServerPort:         rawConf.ServerPort,
		ProtectionFailMode: failMode,
		ProtectionAPIConfig: configuration.ProtectionAPIConfig{
			ProtectionEndpoint: rawConf.ProtectionEndpoint,
			ProtectionToken:    rawConf.ProtectionToken,
		},
		ProcessorConfig: configuration.ProcessorConfig{
			CustomHeaderSignals: CustomHeaderSignals,
			CustomHeaderCookies: CustomHeaderCookies,
		},
		ForwardEndPoint: rawConf.ForwardEndPoint,
	}, nil
}
