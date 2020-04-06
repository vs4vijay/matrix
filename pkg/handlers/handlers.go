package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	h := &Handler{logger}

	return h
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Index)
}

func (h *Handler) Index(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(rw, "Go Web Server")
}
