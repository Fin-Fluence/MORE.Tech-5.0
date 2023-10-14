package entity

type FilterPosition struct {
	Latitude     *float64 `json:"latitude"`
	Longitude    *float64 `json:"longitude"`
	MetroStation *string  `json:"metro_station"`
}

type FilterOffice struct {
	Status       *string         `json:"status"`
	Type         *string         `json:"type"`
	HasRamp      *bool           `json:"has_ramp"`
	Position     *FilterPosition `json:"position"`
	ServiceNames []string        `json:"service_names"`
}

type FilterATM struct {
	AllDay       *bool           `json:"all_day"`
	Position     *FilterPosition `json:"position"`
	ServiceNames []string        `json:"service_names"`
}
