package http

import (
	"encoding/json"
	"net/http"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/MORE.Tech-5.0/server/internal/service"
)

type ATM struct {
	service service.ATM
}

func NewATM(service service.ATM) *ATM {
	return &ATM{service}
}

// GetAll godoc
//
//	@Summary		Get all ATMs
//	@Tags			ATMs
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]entity.ATM
//	@Failure		500	{object}	JSONError
//	@Router			/atm [get]
func (a *ATM) GetAll(w http.ResponseWriter, r *http.Request) {
	atms, err := a.service.GetCache()
	if err != nil {
		InternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := json.NewEncoder(w)
	e.Encode(atms)
}

func (a *ATM) Get(w http.ResponseWriter, r *http.Request) {
	var filter entity.FilterATM
	d := json.NewDecoder(r.Body)
	err := d.Decode(&filter)
	if err != nil {
		ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	atms, err := a.service.Get(filter)
	if err != nil {
		InternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := json.NewEncoder(w)
	e.Encode(atms)
}
