package handler

import (
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	configuration "github.com/dre-zouhair/interceptor/config"

	"github.com/dre-zouhair/interceptor/internal/processor"
	"github.com/dre-zouhair/interceptor/internal/utils"
	"github.com/rs/zerolog/log"
)

type interceptorHandler struct {
	conf      configuration.Config
	processor processor.IProcessorService
}

func NewInterceptorHandler(conf configuration.Config, processor processor.IProcessorService) IInterceptorHandler {
	return &interceptorHandler{
		conf:      conf,
		processor: processor,
	}
}

type IInterceptorHandler interface {
	ForwardAllRequests(w http.ResponseWriter, r *http.Request)
	ProxyAllRequests(w http.ResponseWriter, r *http.Request)
}

func (h interceptorHandler) ForwardAllRequests(w http.ResponseWriter, r *http.Request) {
	h.processRequests(w, r, h.forwardRequest)
}

func (h interceptorHandler) ProxyAllRequests(w http.ResponseWriter, r *http.Request) {
	h.processRequests(w, r, h.proxyRequest)
}

type returnPolicy func(w http.ResponseWriter, r *http.Request)

func (h interceptorHandler) processRequests(w http.ResponseWriter, r *http.Request, performReturn returnPolicy) {
	action := utils.ALLOW_ACCESS

	report, err := h.processor.Process(r)

	log.Debug().Interface("report", report).Msg("report result")

	if err != nil {
		log.Error().Err(err).Msg("failed to perform protection validation")
		action = h.conf.ProtectionFailMode
	} else {
		action = report.Action
	}

	if action == utils.ALLOW_ACCESS {
		performReturn(w, r)
	} else if report.Action == utils.VERIFY_ACCESS {
		performReturn(w, r)
	} else {
		http.Error(w, "Request from bot denied", http.StatusForbidden)
	}
}

func (h interceptorHandler) forwardRequest(w http.ResponseWriter, r *http.Request) {

	forwardURL, err := utils.BuildURL(h.conf.ForwardEndPoint, r.URL.Path, r.URL.RawQuery)

	if err != nil {
		log.Error().Err(err).Msg("failed to forward request")
	}

	r.URL = forwardURL
	r.RequestURI = ""
	r.Host = forwardURL.Host

	targetResp, err := http.DefaultClient.Do(r)

	if err != nil {
		log.Error().Err(err).Msg("failed to forward request")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to close response body")
		}
	}(targetResp.Body)

	for header, values := range targetResp.Header.Clone() {
		w.Header()[header] = values
	}

	for _, cookie := range targetResp.Cookies() {
		http.SetCookie(w, cookie)
	}

	w.WriteHeader(targetResp.StatusCode)

	log.Debug().Int("status", targetResp.StatusCode).Str("url", forwardURL.String()).Msg("forward started")

	_, err = io.Copy(w, targetResp.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to copy response body")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (h interceptorHandler) proxyRequest(w http.ResponseWriter, r *http.Request) {
	targetURL, err := url.Parse(h.conf.ForwardEndPoint)

	if err != nil {
		log.Error().Err(err).Msg("failed to parse forward target url")
		http.Error(w, "Request from bot denied", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	r.Host = targetURL.Host
	r.URL.Scheme = targetURL.Scheme
	proxy.ServeHTTP(w, r)
}
