package handler

import "net/http"

type CORSHandler struct {
	Next http.Handler
}

func (h CORSHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	origin := req.Header.Get("Origin")
	if origin == "" {
		h.Next.ServeHTTP(w, req)
		return
	}

	header := w.Header()
	header.Set("Access-Control-Allow-Origin", origin)
	header.Set("Access-Control-Allow-Credentials", "true")

	if req.Method == http.MethodOptions {
		header.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,HEAD")
		header.Set("Access-Control-Allow-Headers", "*")
		header.Set("Access-Control-Max-Age", "86400")
		return
	}

	h.Next.ServeHTTP(w, req)
}
