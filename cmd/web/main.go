package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/dfioravanti/go-rest/internal/services"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Application struct {
	logger *slog.Logger
	dbPool *pgxpool.Pool

	service *services.SnippetService
}

// migrate the database to the last version
func (app Application) migrateDB() {
	m, err := migrate.New(
		"file://db/migrations", // the path is relative to where the code is run
		app.dbPool.Config().ConnString(),
	)
	if err != nil {
		app.logger.Error(err.Error())
	}

	err = m.Up()
	if err != nil {
		app.logger.Error(err.Error())
	}
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgresql://user:password@localhost:5432?sslmode=disable", "Postgres data source connection string")
	flag.Parse()

	app := InitApp(*dsn)
	defer app.dbPool.Close()

	app.migrateDB()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	app.logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)

}
