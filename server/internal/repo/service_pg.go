package repo

import (
	"context"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/MORE.Tech-5.0/server/pkg/postgres"
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

type ServicePostgres struct {
	*postgres.Postgres
}

func NewServicePostgres(pg *postgres.Postgres) *OpenHoursPostgres {
	return &OpenHoursPostgres{pg}
}

func (s *ServicePostgres) GetByAtmId(ctx context.Context, atmID uuid.UUID) ([]entity.Service, error) {
	const query = `SELECT id, name, capability, activity FROM service WHERE atm_id = $1`

	rows, err := s.Pool.Query(ctx, query, atmID)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[entity.Service])
}
