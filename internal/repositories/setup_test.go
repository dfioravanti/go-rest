package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/dfioravanti/go-rest/internal/testcontainer"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	psqlContainer *testcontainer.PostgreSQLContainer
	//server        *httptest.Server
}

func (s *TestSuite) SetupSuite() {
	// create db container
	ctx, ctxCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer ctxCancel()

	psqlContainer, err := testcontainer.NewPostgreSQLContainer(ctx)
	s.NoError(err)

	s.psqlContainer = psqlContainer

}

func (s *TestSuite) SetupTest() {
	m, err := migrate.New(
		"file://../../db/migrations",
		s.psqlContainer.GetDSN(),
	)
	s.NoError(err)

	err = m.Up()
	s.NoError(err)
}

func (s *TestSuite) TearDownTest() {

	stmt := `
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
		GRANT ALL ON SCHEMA public TO public;
	`

	dbpool, err := pgxpool.New(context.Background(), s.psqlContainer.GetDSN())
	s.NoError(err)

	_, err = dbpool.Exec(context.Background(), stmt)
	s.NoError(err)

}

func (s *TestSuite) TearDownSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	s.Require().NoError(s.psqlContainer.Terminate(ctx))

	//s.server.Close()
}

func TestSuite_Run(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
