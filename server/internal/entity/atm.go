package entity

import "github.com/gofrs/uuid/v5"

type ATM struct {
	ID       uuid.UUID
	Position *Position
	Address  string
	AllDay   bool
	Services []Service
}
