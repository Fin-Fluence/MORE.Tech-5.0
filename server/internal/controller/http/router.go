package http

import (
	"github.com/MORE.Tech-5.0/server/internal/service"
	"github.com/go-chi/chi/v5"
)

type Services struct {
	Office service.Office
	ATM    service.ATM
}

func Router(services Services) *chi.Mux {
	mux := chi.NewMux()

	office := NewOffice(services.Office)
	atm := NewATM(services.ATM)

	mux.Route("/", func(r chi.Router) {
		r.Route("/office", func(r chi.Router) {
			r.Get("/", office.GetAll)
		})
		r.Route("/atm", func(r chi.Router) {
			r.Get("/", atm.GetAll)
		})
	})

	return mux
}
