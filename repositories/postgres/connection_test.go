package postgres

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *TestSuite) TestItCanConnect() {

	dbpool, err := pgxpool.New(context.Background(), s.psqlContainer.GetDSN())
	s.NoError(err)
	defer dbpool.Close()

	err = dbpool.Ping(context.Background())
	s.NoError(err)

}
