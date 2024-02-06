package processor

import (
	"net/http"
	"time"

	configuration "github.com/dre-zouhair/interceptor/config"

	"github.com/dre-zouhair/interceptor/internal/analyser"
	"github.com/dre-zouhair/interceptor/internal/builder"
)

type processorService struct {
	conf          configuration.ProcessorConfig
	protectionCli analyser.IProtectionCli
}

func NewProcessorService(conf configuration.ProcessorConfig, cli analyser.IProtectionCli) IProcessorService {
	return &processorService{
		conf:          conf,
		protectionCli: cli,
	}
}

type IProcessorService interface {
	Process(r *http.Request) (*analyser.ValidationResponse, error)
}

func (s processorService) Process(r *http.Request) (*analyser.ValidationResponse, error) {
	signals := builder.NewSignalsBuilder().
		BuildHeadersSignals(r.Header).
		BuildRealRemoteAddr(r.Header, r.RemoteAddr).
		BuildCustomHeadersSignals(r.Header, s.conf.CustomHeaderSignals).
		BuildCustomCookiesSignals(r.Cookies(), s.conf.CustomHeaderCookies).
		GetSignals()

	signals.ContentLength = r.ContentLength
	signals.Method = r.Method
	signals.Query = r.URL.RawQuery
	signals.Time = time.Now()

	return s.protectionCli.Validate(signals)
}
