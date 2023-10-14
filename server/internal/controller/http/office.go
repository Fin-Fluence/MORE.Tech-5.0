package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/MORE.Tech-5.0/server/internal/service"
)

type Office struct {
	service service.Office
}

func NewOffice(service service.Office) *Office {
	return &Office{service}
}

// Get godoc
//
//	@Summary	Get offices
//	@Tags		Offices
//	@Accept		json
//	@Produce	json
//	@Param		filter	query		string	false	"Filter"	Format(entity.FilterOffice)
//	@Success	200	{object}	[]entity.Office
//	@Failure	500	{object}	JSONError
//	@Router		/office [get]
func (o *Office) Get(w http.ResponseWriter, r *http.Request) {
	var filter entity.FilterOffice
	param := r.URL.Query().Get("filter")
	if param != "" {
		reader := strings.NewReader(param)
		d := json.NewDecoder(reader)
		err := d.Decode(&filter)
		if err != nil {
			ErrorJSON(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	offices, err := o.service.Get(filter)
	if err != nil {
		InternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	e := json.NewEncoder(w)
	e.Encode(offices)
}
