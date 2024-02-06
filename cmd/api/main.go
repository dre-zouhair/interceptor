package main

import (
	"github.com/dre-zouhair/interceptor/internal/analyser"
	"github.com/dre-zouhair/interceptor/internal/config"
	"github.com/dre-zouhair/interceptor/internal/handler"
	"github.com/dre-zouhair/interceptor/internal/processor"
	"github.com/rs/zerolog/log"
	"net/http"
)

var appConf *config.Config

func main() {

	if appConf == nil {
		log.Error().Msg("missing app configuration")
	}

	protectionCli := analyser.NewProtectionCli(appConf.ProtectionAPIConfig)
	processorService := processor.NewProcessorService(appConf.ProcessorConfig, protectionCli)

	h := handler.NewInterceptorHandler(*appConf, processorService)

	http.HandleFunc("/", h.HandleAllRequests)

	err := http.ListenAndServe(":"+appConf.ServerPort, nil)

	if err != nil {
		log.Error().Err(err).Msg("unable to start server")
	}
}
