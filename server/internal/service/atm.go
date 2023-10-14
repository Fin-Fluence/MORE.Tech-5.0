package service

import (
	"context"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/MORE.Tech-5.0/server/internal/repo"
)

type ATMRepos struct {
	ATM     repo.ATM
	Service repo.Service
}

type atm struct {
	repos ATMRepos
}

func NewATM(repos ATMRepos) *atm {
	return &atm{repos}
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
