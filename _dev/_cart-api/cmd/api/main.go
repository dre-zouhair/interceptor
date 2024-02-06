package main

import (
	"net/http"

	"dre-zouhair/modules/cart-api/config"
	"dre-zouhair/modules/cart-api/internal/handler"
	"dre-zouhair/modules/cart-api/internal/repository"
	"dre-zouhair/modules/cart-api/internal/service"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bunrouter"
)

func main() {

	appConf, err := config.LoadConfig()

	if err != nil {
		log.Error().Err(err).Msg("unable to load config")
	}

	router := bunrouter.New(
		bunrouter.WithNotFoundHandler(handler.NotFoundHandler),
		bunrouter.WithMethodNotAllowedHandler(handler.MethodNotAllowedHandler),
	)

	cartRepo := repository.NewCartRepository()
	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)

	router.WithGroup("/api/v1", func(g *bunrouter.Group) {

		g.GET("/ping", func(w http.ResponseWriter, req bunrouter.Request) error {
			return bunrouter.JSON(w, bunrouter.H{
				"Referer": req.Referer(),
			})
		})

		g.POST("/cart", cartHandler.AddItemHandler)
		g.GET("/cart", cartHandler.GetUserItemsHandler)
	})

	err = http.ListenAndServe(":"+appConf.ServerPort, router)
	if err != nil {
		log.Error().Err(err).Msg("unable to start server")
	}
}
