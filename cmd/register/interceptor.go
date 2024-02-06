package register

import (
	"net/http"

	configuration "github.com/dre-zouhair/interceptor/config"
	"github.com/dre-zouhair/interceptor/internal/analyser"
	"github.com/dre-zouhair/interceptor/internal/handler"
	"github.com/dre-zouhair/interceptor/internal/processor"
)

func Interceptor(path string, conf configuration.Config) {
	protectionCli := analyser.NewProtectionCli(conf.ProtectionAPIConfig)
	processorService := processor.NewProcessorService(conf.ProcessorConfig, protectionCli)

	h := handler.NewInterceptorHandler(conf, processorService)

	http.HandleFunc(path, h.ForwardAllRequests)
}
