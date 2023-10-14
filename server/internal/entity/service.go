package entity

import "github.com/gofrs/uuid/v5"

type Service struct {
	ID         uuid.UUID
	Name       string
	Capability bool
	Activity   bool
}
