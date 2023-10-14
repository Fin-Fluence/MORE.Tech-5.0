package service

import "github.com/MORE.Tech-5.0/server/internal/entity"

type Office interface {
	GetAll() ([]entity.Office, error)
	GetCache() ([]entity.Office, error)
	Get(filter entity.FilterOffice) ([]entity.Office, error)
}

type ATM interface {
	GetAll() ([]entity.ATM, error)
	GetCache() ([]entity.ATM, error)
	Get(filter entity.FilterATM) ([]entity.ATM, error)
}
