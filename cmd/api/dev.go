//go:build dev
// +build dev

package main

import (
	"github.com/dre-zouhair/interceptor/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {

	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("unable to load env config file")
		return
	}

	var rawConf config.RawConfig

	if err := viper.Unmarshal(&rawConf); err != nil {
		log.Error().Err(err).Msg("unable to unmarshal env config file")
		return
	}

	conf, err := config.Build(rawConf)

	if err != nil {
		log.Error().Err(err).Msg("invalid configuration")
		return
	}

	appConf = conf
}
