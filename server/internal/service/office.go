package service

import (
	"context"
	"slices"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/MORE.Tech-5.0/server/internal/repo"
)

type OfficeRepos struct {
	Office    repo.Office
	OpenHours repo.OpenHours
	Service   repo.OfficeService
	Cache     *repo.OfficeCache
}

type office struct {
	repos OfficeRepos
}

func NewOffice(repos OfficeRepos) *office {
	o := &office{repos}
	offices, _ := o.GetAll()
	o.repos.Cache.Set(offices)
	return o
}

func (o *office) GetAll() ([]entity.Office, error) {
	offices, err := o.repos.Office.GetAll(context.Background())
	if err != nil {
		return nil, err
	}

	for i, office := range offices {
		oh, err := o.repos.OpenHours.GetByOfficeId(context.Background(), office.ID, false)
		if err != nil {
			return nil, err
		}
		offices[i].OpenHours = oh

		ohi, err := o.repos.OpenHours.GetByOfficeId(context.Background(), office.ID, true)
		if err != nil {
			return nil, err
		}
		offices[i].OpenHoursIndividual = ohi

		services, err := o.repos.Service.GetByOfficeId(context.Background(), office.ID)
		if err != nil {
			return nil, err
		}
		offices[i].Services = services
	}

	return offices, nil
}

func (o *office) GetCache() ([]entity.Office, error) {
	return o.repos.Cache.GetAll(context.Background())
}

func (o *office) Get(filter entity.FilterOffice) []entity.Office {
	filtered := o.Filter(filter)

	if filter.Position != nil && filter.Position.Latitude != nil && filter.Position.Longitude != nil {
		pos := entity.Position{
			Latitude:  *filter.Position.Latitude,
			Longitude: *filter.Position.Longitude,
		}

		result := make([]entity.Office, 0)
		for _, office := range filtered {
			d := getDistance(pos, *office.Position)
			load := office.Load

			if d <= 2 && load <= 20 {
				office.Load = load
				result = append(result, office)
			}
		}

		if len(result) == 0 {
			for _, office := range filtered {
				d := getDistance(pos, *office.Position)
				load := office.Load

				if (d <= 9 && load <= 20) || (d <= 2 && load <= 40) {
					office.Load = load
					result = append(result, office)
				}
			}
		}

		if len(result) == 0 {
			for _, office := range filtered {
				d := getDistance(pos, *office.Position)
				load := office.Load

				if (d <= 25 && load <= 20) || (d <= 9 && load <= 40) {
					result = append(result, office)
				}
			}
		}

		return result
	}

	return filtered
}

func (o *office) Filter(filter entity.FilterOffice) []entity.Office {
	filteredOffices := make([]entity.Office, 0)

	offices, _ := o.GetCache()
	for _, office := range offices {
		if filter.Status != nil && office.Status != *filter.Status {
			continue
		}

		if filter.Type != nil && office.Type != *filter.Type {
			continue
		}

		if filter.HasRamp != nil && (office.HasRamp == nil || *office.HasRamp != *filter.HasRamp) {
			continue
		}

		if filter.Position != nil {
			if filter.Position.MetroStation != nil && (office.Position.MetroStation == nil || *office.Position.MetroStation != *filter.Position.MetroStation) {
				continue
			}
		}

		if names := filter.ServiceNames; names != nil {
			var contains bool
			for _, name := range names {
				contains = slices.ContainsFunc(office.Services, func(os entity.OfficeService) bool {
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

		office.Load = getLoad(filter, office.Services)
		filteredOffices = append(filteredOffices, office)
	}

	return filteredOffices
}
