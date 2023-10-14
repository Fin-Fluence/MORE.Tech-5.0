package repo

import (
	"context"

	"github.com/MORE.Tech-5.0/server/internal/entity"
	"github.com/MORE.Tech-5.0/server/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

type ATMPostgres struct {
	*postgres.Postgres
}

func NewATMPostgres(pg *postgres.Postgres) *ATMPostgres {
	return &ATMPostgres{pg}
}

func (a *ATMPostgres) GetAll(ctx context.Context) ([]entity.ATM, error) {
	const query = `SELECT a.id, a.address, a.all_day, p.latitude, p.longitude, p.metro_station 
		FROM atm a JOIN position p ON a.position_id = p.id`
	rows, err := a.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (entity.ATM, error) {
		var (
			atm      entity.ATM
			position entity.Position
		)

		err := row.Scan(&atm.ID, &atm.Address, &atm.AllDay, &position.Latitude, &position.Longitude, &position.MetroStation)
		if err != nil {
			return atm, err
		}

		atm.Position = &position
		return atm, nil
	})
}
