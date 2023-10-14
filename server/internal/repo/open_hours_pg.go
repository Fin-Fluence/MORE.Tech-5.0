package repo

import (
	"context"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/MORE.Tech-5.0/server/pkg/postgres"
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

type OpenHoursPostgres struct {
	*postgres.Postgres
}

func NewOpenHoursPostgres(pg *postgres.Postgres) *OpenHoursPostgres {
	return &OpenHoursPostgres{pg}
}

func (o *OpenHoursPostgres) GetByOfficeId(ctx context.Context, officeID uuid.UUID, isIndividual bool) ([]entity.OpenHours, error) {
	const query = `SELECT id, day, hours, is_individual FROM open_hours WHERE office_id = $1 AND is_individual = $2`

	rows, err := o.Pool.Query(ctx, query, officeID, isIndividual)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[entity.OpenHours])
}
