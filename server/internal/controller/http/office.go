package http

import (
	"encoding/json"
	"net/http"

	"github.com/MORE.Tech-5.0/server/internal/service"
)

type Office struct {
	service service.Office
}

func NewOffice(service service.Office) *Office {
	return &Office{service}
}

func (o *Office) GetAll(w http.ResponseWriter, r *http.Request) {
	offices, err := o.service.GetAll()
	if err != nil {
		InternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := json.NewEncoder(w)
	e.Encode(offices)
}
