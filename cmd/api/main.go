package main

import (
	"net/http"

	"github.com/dre-zouhair/interceptor/cmd/register"
	configuration "github.com/dre-zouhair/interceptor/config"
	"github.com/rs/zerolog/log"
)

var appConf *configuration.Config

func main() {

	if appConf == nil {
		log.Error().Msg("missing app configuration")
		return
	}

	register.Interceptor(appConf.InterceptionPath, *appConf)

	err := http.ListenAndServe(":"+appConf.ServerPort, nil)

	if err != nil {
		log.Error().Err(err).Msg("unable to start server")
	}

}
