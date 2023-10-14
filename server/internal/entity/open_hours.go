package entity

import "github.com/gofrs/uuid/v5"

type OpenHours struct {
	ID    uuid.UUID `json:"-"`
	Day   string    `json:"day"`
	Hours string    `json:"hours"`
}
