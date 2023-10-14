package http

import (
	"github.com/MORE.Tech-5.0/server/internal/service"
	"github.com/go-chi/chi/v5"
)

func Router(officeService service.Office) *chi.Mux {
	mux := chi.NewMux()

	office := NewOffice(officeService)

	mux.Route("/", func(r chi.Router) {
		r.Route("/office", func(r chi.Router) {
			r.Get("/", office.GetAll)
		})
	})

	return mux
}
