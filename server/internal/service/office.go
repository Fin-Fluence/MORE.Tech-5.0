package service

import (
	"context"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/MORE.Tech-5.0/server/internal/repo"
)

type OfficeRepos struct {
	Office    repo.Office
	OpenHours repo.OpenHours
	Service   repo.OfficeService
}

type office struct {
	repos OfficeRepos
}

func NewOffice(repos OfficeRepos) *office {
	return &office{repos}
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
