package service

import (
	"context"
	"fmt"
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
	offices, err := o.GetAll()
	fmt.Println(err)
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

func (o *office) Get(filter entity.FilterOffice) ([]entity.Office, error) {
	return o.Filter(filter)
}

func (o *office) Filter(filter entity.FilterOffice) ([]entity.Office, error) {
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

		filteredOffices = append(filteredOffices, office)
	}

	return filteredOffices, nil
}
