package http

import (
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	mux := chi.NewMux()
	mux.Route("/", func(r chi.Router) {

	})

	return mux
}
