package rest

import (
	"log/slog"

	"github.com/dfioravanti/go-rest/rest/handlers"
	"github.com/golang-migrate/migrate"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Logger *slog.Logger
	DbPool *pgxpool.Pool

	SnippetHandler handlers.SnippetHandler
}

func NewApplication(logger *slog.Logger, dbPool *pgxpool.Pool, snippetHandler handlers.SnippetHandler) Application {
	return Application{logger, dbPool, snippetHandler}
}

// migrate the database to the last version
func (app Application) MigrateDB() {
	m, err := migrate.New(
		"file://db/migrations", // the path is relative to where the code is run
		app.DbPool.Config().ConnString(),
	)
	if err != nil {
		app.Logger.Error(err.Error())
	}

	err = m.Up()
	if err != nil {
		app.Logger.Error(err.Error())
	}
}
