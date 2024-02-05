package processor

import (
	"github.com/dre-zouhair/interceptor/config"
	"github.com/dre-zouhair/interceptor/internal/builder"
	"github.com/dre-zouhair/interceptor/internal/protectioncli"
	"net/http"
	"time"
)

type processorService struct {
	conf          config.Config
	protectionCli protectioncli.IProtectionCli
}

func NewProcessorService(conf config.Config, cli protectioncli.IProtectionCli) IProcessorService {
	return &processorService{
		conf:          conf,
		protectionCli: cli,
	}
}

type IProcessorService interface {
	Process(r *http.Request) (*protectioncli.ValidationResponse, error)
}

func (s processorService) Process(r *http.Request) (*protectioncli.ValidationResponse, error) {
	signals := builder.NewSignalsBuilder().
		BuildHeadersSignals(r.Header).
		BuildRealRemoteAddr(r.Header, r.RemoteAddr).
		BuildCustomHeadersSignals(r.Header, s.conf.CustomHeaderSignals).
		BuildCookiesSignals(r.Cookies()).
		BuildCustomCookiesSignals(r.Cookies(), s.conf.CustomHeaderCookies).
		GetSignals()

	signals.ContentLength = r.ContentLength
	signals.Method = r.Method
	signals.Time = time.Now()

	return s.protectionCli.Validate(signals)
}
