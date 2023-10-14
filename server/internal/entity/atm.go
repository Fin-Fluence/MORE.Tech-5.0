package entity

import "github.com/gofrs/uuid/v5"

type ATM struct {
	ID       uuid.UUID `json:"id"`
	Position *Position `json:"position"`
	Address  string    `json:"address"`
	AllDay   bool      `json:"all_day"`
	Services []Service `json:"services"`
}
