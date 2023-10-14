package repo

import (
	"context"

	"github.com/MORE.Tech-5.0/server/internal/entity"
)

type OfficeCache []entity.Office

func NewOfficeCache() *OfficeCache {
	return &OfficeCache{}
}

func (o *OfficeCache) Set(offices []entity.Office) {
	*o = make([]entity.Office, len(offices))
	copy(*o, offices)
}

func (o *OfficeCache) GetAll(ctx context.Context) ([]entity.Office, error) {
	return *o, nil
}
