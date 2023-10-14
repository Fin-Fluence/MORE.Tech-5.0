package repo

import (
	"context"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/gofrs/uuid/v5"
)

type Office interface {
	GetAll(ctx context.Context) ([]entity.Office, error)
}

type ATM interface {
	GetAll(ctx context.Context) ([]entity.ATM, error)
}

type OpenHours interface {
	GetByOfficeId(ctx context.Context, officeID uuid.UUID, isIndividual bool) ([]entity.OpenHours, error)
}

type Service interface {
	GetByAtmId(ctx context.Context, atmID uuid.UUID) ([]entity.Service, error)
}
