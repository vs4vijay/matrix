package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vs4vijay/matrix/pkg/services"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	h := &Handler{logger}

	return h
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.LoggerMiddleware(h.Index))
	mux.HandleFunc("/data", h.LoggerMiddleware(h.Data))
}

func (h *Handler) Index(rw http.ResponseWriter, request *http.Request) {
	// h.logger.Printf("Request received: %+v\n", request)
	fmt.Fprintf(rw, "Go Web Server")
}

func (h *Handler) LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request(%s) processed in %s\n", r.RequestURI ,time.Now().Sub(startTime))
		next(w, r)
	}
}

func (h *Handler) Data(rw http.ResponseWriter, request *http.Request) {
	dataService := services.DataService{}
	fmt.Fprintf(rw, dataService.FirstName())
}
