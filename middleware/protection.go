package middleware

import (
	"context"
	"github.com/rs/zerolog/log"
	"net/http"

	configuration "github.com/dre-zouhair/interceptor/config"
	"github.com/dre-zouhair/interceptor/internal/analyser"
	"github.com/dre-zouhair/interceptor/internal/processor"
	"github.com/dre-zouhair/interceptor/internal/utils"
	"github.com/uptrace/bunrouter"
)

type protectionMiddleware struct {
	processor processor.IProcessorService
}

func NewProtectionMiddleware(conf configuration.ProtectionMiddlewareConfig) IProtectionMiddleware {
	protectionCli := analyser.NewProtectionCli(conf.ProtectionAPIConfig)
	processorService := processor.NewProcessorService(conf.ProcessorConfig, protectionCli)

	return &protectionMiddleware{
		processor: processorService,
	}
}

type IProtectionMiddleware interface {
	BunProtectionMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc
	StandardProtectionMiddleware(next http.Handler) http.Handler
}

func (m protectionMiddleware) BunProtectionMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		action := m.processRequest(req.Request.Clone(context.Background()))

		if action == utils.BLOCK_ACCESS {
			http.Error(w, "Request from bot denied", http.StatusForbidden)
			return nil
		} else if action == utils.VERIFY_ACCESS {
			return next(w, req)
		} else {
			return next(w, req)
		}
	}
}

func (m protectionMiddleware) StandardProtectionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		action := m.processRequest(r)

		if action == utils.BLOCK_ACCESS {
			http.Error(w, "Request from bot denied", http.StatusForbidden)
			return
		} else if action == utils.VERIFY_ACCESS {
			next.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}

	})
}

func (m protectionMiddleware) processRequest(r *http.Request) string {
	action := utils.ALLOW_ACCESS

	report, err := m.processor.Process(r)

	if err == nil {
		action = report.Action
	} else {
		log.Error().Err(err).Msg("error processing request")
	}

	return action
}
