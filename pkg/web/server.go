package server

import (
	"net/http"
)

func New(mux *http.ServeMux, address string) *http.Server {

	srv := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return srv
}
