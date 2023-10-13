package entity

import "github.com/gofrs/uuid/v5"

type OpenHours struct {
	ID    uuid.UUID
	Day   string
	Hours string
}
