package entity

import "github.com/gofrs/uuid/v5"

type Position struct {
	ID           uuid.UUID `json:"-"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	MetroStation *string   `json:"metro_station"`
}
