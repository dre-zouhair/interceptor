package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dre-zouhair/modules/cart-api/internal/repository"
	"dre-zouhair/modules/cart-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bunrouter"
)

type cartHandler struct {
	cartService service.ICartService
}

func NewCartHandler(cartService service.ICartService) ICartHandler {
	return &cartHandler{
		cartService: cartService,
	}
}

type ICartHandler interface {
	AddItemHandler(w http.ResponseWriter, req bunrouter.Request) error
	GetUserItemsHandler(w http.ResponseWriter, req bunrouter.Request) error
}

func (h cartHandler) AddItemHandler(w http.ResponseWriter, req bunrouter.Request) error {

	var item repository.Item
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		log.Error().Err(err).Msg("unable to process request payload")
		http.Error(w, "unable to process request payload", http.StatusUnprocessableEntity)
		return nil
	}

	err = validator.New().Struct(item)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		log.Error().Err(err).Msg("unable to process request payload")
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return nil
	}

	userID, err := req.Cookie("user-id")

	if err != nil || len(userID.Value) == 0 {
		userID = &http.Cookie{
			Name:  "user-id",
			Value: uuid.NewString(),
		}
		http.SetCookie(w, userID)
	}

	err = h.cartService.Add(userID.Value, item)

	if err != nil {
		log.Error().Err(err).Msg("unable to save cart item")
		http.Error(w, "something went wrong try later", http.StatusInternalServerError)
		return nil
	}

	w.WriteHeader(http.StatusCreated)

	return nil
}

func (h cartHandler) GetUserItemsHandler(w http.ResponseWriter, req bunrouter.Request) error {

	userID, err := req.Cookie("user-id")

	if err != nil || len(userID.Value) == 0 {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return nil
	}

	items, err := h.cartService.Get(userID.Value)

	if err != nil {
		http.Error(w, "invalid request", http.StatusNotFound)
		return nil
	}

	return bunrouter.JSON(w, items)
}
