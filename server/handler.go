package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go, using fx and chi!")
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	return r
}
