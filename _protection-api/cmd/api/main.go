package main

import (
	"github.com/dre-zouhair/modules/protection-api/config"
	"github.com/dre-zouhair/modules/protection-api/internal/handler"
	"github.com/dre-zouhair/modules/protection-api/internal/middleware"
	"github.com/dre-zouhair/modules/protection-api/internal/service"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bunrouter"
	"net/http"
)

func main() {

	appConf, err := config.LoadConfig()

	if err != nil {
		log.Error().Err(err).Msg("unable to load config")
	}

	authMiddleware := middleware.NewAuthMiddleware(appConf.ApiTokens)

	router := bunrouter.New(
		bunrouter.WithNotFoundHandler(handler.NotFoundHandler),
		bunrouter.WithMethodNotAllowedHandler(handler.MethodNotAllowedHandler),
		bunrouter.Use(authMiddleware.AuthMiddleware),
	)

	validationService := service.NewValidationService()
	validationHandler := handler.NewValidationHandler(validationService)

	router.WithGroup("/api/v1", func(g *bunrouter.Group) {

		g.GET("/ping", func(w http.ResponseWriter, req bunrouter.Request) error {
			return bunrouter.JSON(w, bunrouter.H{
				"Referer": req.Referer(),
			})
		})
		g.POST("/validate", validationHandler.ValidationHandler)
	})

	err = http.ListenAndServe(":"+appConf.ServerPort, router)
	if err != nil {
		log.Error().Err(err).Msg("unable to start server")
	}
}
