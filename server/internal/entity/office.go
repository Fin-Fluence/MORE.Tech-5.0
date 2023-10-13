package entity

import "github.com/gofrs/uuid/v5"

type Office struct {
	ID                  uuid.UUID
	SalePointName       string
	Address             string
	Status              string
	OpenHours           []OpenHours
	RKO                 *bool
	OpenHoursIndividual []OpenHours
	Type                string
	SalePointFormat     string
	SuoAvailability     *bool
	HasRamp             *bool
	Position            *Position
	Distance            int
	KEP                 *bool
	MyBranch            bool
}
