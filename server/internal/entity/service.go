package entity

import "github.com/gofrs/uuid/v5"

type Service struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Capability bool      `json:"capability"`
	Activity   bool      `json:"activity"`
}

type OfficeService struct {
	*Service
	CurrentTicket string `json:"current_ticket"`
	LastTicket    string `json:"last_ticket"`
}
