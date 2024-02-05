package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dre-zouhair/modules/protection-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bunrouter"
)

type validationHandler struct {
	cartService service.IValidationService
}

func NewValidationHandler(cartService service.IValidationService) IValidationHandler {
	return &validationHandler{
		cartService: cartService,
	}
}

type IValidationHandler interface {
	ValidationHandler(w http.ResponseWriter, req bunrouter.Request) error
}

func (h validationHandler) ValidationHandler(w http.ResponseWriter, req bunrouter.Request) error {

	var signals service.Signals
	err := json.NewDecoder(req.Body).Decode(&signals)
	if err != nil {
		log.Error().Err(err).Msg("unable to process request payload")
		http.Error(w, "unable to process request payload", http.StatusUnprocessableEntity)
		return nil
	}

	err = validator.New().Struct(signals)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		log.Error().Err(err).Msg("unable to process request payload")
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return nil
	}

	report, err := h.cartService.Validate(signals)

	if err != nil {
		log.Error().Err(err).Msg("unable to perform validation")
		http.Error(w, "something went wrong try later", http.StatusInternalServerError)
		return nil
	}

	w.WriteHeader(http.StatusCreated)

	return bunrouter.JSON(w, *report)
}
