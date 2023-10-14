package repo

import (
	"context"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/MORE.Tech-5.0/server/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

type OfficePostgres struct {
	*postgres.Postgres
}

func NewOfficePostgres(pg *postgres.Postgres) *OfficePostgres {
	return &OfficePostgres{pg}
}

func (o *OfficePostgres) GetAll(ctx context.Context) ([]entity.Office, error) {
	const query = `SELECT o.id, o.sale_point_name, o.address, o.status, o.rko, o.type, o.sale_point_format, o.suo_availability, o.has_ramp, o.distance, o.kep, o.my_branch, p.latitude, p.longitude, p.metro_station 
		FROM office o JOIN position p ON o.position_id = p.id`
	rows, err := o.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (entity.Office, error) {
		var (
			office   entity.Office
			position entity.Position
		)

		err := row.Scan(&office.ID, &office.SalePointName, &office.Address, &office.Status, &office.RKO, &office.Type, &office.SalePointFormat, &office.SuoAvailability, &office.HasRamp, &office.Distance, &office.KEP, &office.MyBranch, &position.Latitude, &position.Longitude, &position.MetroStation)
		if err != nil {
			return office, err
		}

		office.Position = &position
		return office, nil
	})
}
