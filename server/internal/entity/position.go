package entity

import "github.com/gofrs/uuid/v5"

type Position struct {
	ID           uuid.UUID
	Latitude     float64
	Longitude    float64
	MetroStation *string
}
