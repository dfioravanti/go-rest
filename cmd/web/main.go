package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/dfioravanti/go-rest/cmd/web/boot"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgresql://user:password@localhost:5432?sslmode=disable", "Postgres data source connection string")
	flag.Parse()

	app := boot.InitApp(*dsn)
	defer app.DbPool.Close()

	app.MigrateDB()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	app.Logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.Routes())
	app.Logger.Error(err.Error())
	os.Exit(1)

}
