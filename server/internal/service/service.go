package service

import (
	"math"
	"slices"
	"strconv"

	"github.com/MORE.Tech-5.0/server/internal/entity"
)

type Office interface {
	GetAll() ([]entity.Office, error)
	GetCache() ([]entity.Office, error)
	Get(filter entity.FilterOffice) []entity.Office
}

type ATM interface {
	GetAll() ([]entity.ATM, error)
	GetCache() ([]entity.ATM, error)
	Get(filter entity.FilterATM) []entity.ATM
}

const R = 6371

func degToRad(x float64) float64 {
	return x * math.Pi / 180
}

func getDistance(selfPos entity.Position, pos entity.Position) float64 {
	latSelf := degToRad(selfPos.Latitude)
	longSelf := degToRad(selfPos.Longitude)
	lat := degToRad(pos.Latitude)
	long := degToRad(pos.Longitude)
	longDiff := longSelf - long
	angle := math.Acos(math.Sin(latSelf)*math.Sin(lat) + math.Cos(latSelf)*math.Cos(lat)*math.Cos(longDiff))
	return R * angle
}

func getLoad(filter entity.FilterOffice, services []entity.OfficeService) int {
	loadMax := 0
	for _, name := range filter.ServiceNames {
		index := slices.IndexFunc(services, func(os entity.OfficeService) bool {
			return os.Name == name
		})
		service := services[index]

		d1, _ := strconv.Atoi(service.CurrentTicket[2:])
		d2, _ := strconv.Atoi(service.LastTicket[2:])

		if d2 >= d1 {
			d2 -= d1
		} else {
			d2 += 1000 - d1
		}

		load := d2 * 10
		if load > loadMax {
			loadMax = load
		}
	}
	return loadMax
}
