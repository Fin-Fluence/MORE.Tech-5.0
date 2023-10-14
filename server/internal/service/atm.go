package service

import (
	"context"
	"slices"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/MORE.Tech-5.0/server/internal/repo"
)

type ATMRepos struct {
	ATM     repo.ATM
	Service repo.Service
	Cache   *repo.ATMCache
}

type atm struct {
	repos ATMRepos
}

func NewATM(repos ATMRepos) *atm {
	a := &atm{repos}
	atms, _ := a.GetAll()
	a.repos.Cache.Set(atms)
	return a
}

func (a *atm) GetAll() ([]entity.ATM, error) {
	atms, err := a.repos.ATM.GetAll(context.Background())
	if err != nil {
		return nil, err
	}

	for i, atm := range atms {
		services, err := a.repos.Service.GetByAtmId(context.Background(), atm.ID)
		if err != nil {
			return nil, err
		}
		atms[i].Services = services
	}

	return atms, nil
}

func (a *atm) GetCache() ([]entity.ATM, error) {
	return a.repos.Cache.GetAll(context.Background())
}

func (a *atm) Get(filter entity.FilterATM) []entity.ATM {
	filtered := a.Filter(filter)

	if filter.Position != nil && filter.Position.Latitude != nil && filter.Position.Longitude != nil {
		pos := entity.Position{
			Latitude:  *filter.Position.Latitude,
			Longitude: *filter.Position.Longitude,
		}
		result := make([]entity.ATM, 0)
		for _, atm := range filtered {
			d := getDistance(pos, *atm.Position)
			if d <= 25 {
				result = append(result, atm)
			}
		}

		return result
	}

	return filtered
}

func (a *atm) Filter(filter entity.FilterATM) []entity.ATM {
	filteredATMs := make([]entity.ATM, 0)

	atms, _ := a.GetCache()

	for _, atm := range atms {
		if filter.AllDay != nil && atm.AllDay != *filter.AllDay {
			continue
		}

		if filter.Position != nil {
			if filter.Position.MetroStation != nil && (atm.Position.MetroStation == nil || *atm.Position.MetroStation != *filter.Position.MetroStation) {
				continue
			}
		}

		if names := filter.ServiceNames; names != nil {
			var contains bool
			for _, name := range names {
				contains = slices.ContainsFunc(atm.Services, func(os entity.Service) bool {
					return os.Name == name
				})
				if !contains {
					break
				}
			}
			if !contains {
				continue
			}
		}

		filteredATMs = append(filteredATMs, atm)
	}

	return filteredATMs
}
