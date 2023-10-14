package entity

import "github.com/gofrs/uuid/v5"

type Office struct {
	ID                  uuid.UUID       `json:"id"`
	SalePointName       string          `json:"sale_point_name"`
	Address             string          `json:"address"`
	Status              string          `json:"status"`
	OpenHours           []OpenHours     `json:"open_hours"`
	RKO                 *bool           `json:"rko"`
	OpenHoursIndividual []OpenHours     `json:"open_hours_individual"`
	Type                string          `json:"type"`
	SalePointFormat     string          `json:"sale_point_format"`
	SuoAvailability     *bool           `json:"suo_availability"`
	HasRamp             *bool           `json:"has_ramp"`
	Position            *Position       `json:"position"`
	Distance            int             `json:"distance"`
	KEP                 *bool           `json:"kep"`
	MyBranch            bool            `json:"my_branch"`
	Services            []OfficeService `json:"services"`
}
