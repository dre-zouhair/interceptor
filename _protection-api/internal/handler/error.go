package handler

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func NotFoundHandler(w http.ResponseWriter, _ bunrouter.Request) error {
	w.WriteHeader(http.StatusNotFound)
	return nil
}

func MethodNotAllowedHandler(w http.ResponseWriter, _ bunrouter.Request) error {
	w.WriteHeader(http.StatusMethodNotAllowed)
	return nil
}
