package repo

import (
	"context"

	"github.com/MORE.Tech-5.0/server/internal/entity"
)

type ATMCache []entity.ATM

func NewATMCache() *ATMCache {
	return &ATMCache{}
}

func (a *ATMCache) Set(atms []entity.ATM) {
	*a = make([]entity.ATM, len(atms))
	copy(*a, atms)
}

func (a *ATMCache) GetAll(ctx context.Context) ([]entity.ATM, error) {
	return *a, nil
}
