package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/dfioravanti/go-rest/internal/testcontainer"
	"github.com/golang-migrate/migrate/v4"
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

	// run migrations
	m, err := migrate.New(
		"file://../../db/migrations",
		s.psqlContainer.GetDSN(),
	)
	s.NoError(err)

	err = m.Up()
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
