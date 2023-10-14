package service

import "github.com/MORE.Tech-5.0/server/internal/entity"

type Office interface {
	GetAll() ([]entity.Office, error)
}

type ATM interface {
	GetAll() ([]entity.ATM, error)
}
