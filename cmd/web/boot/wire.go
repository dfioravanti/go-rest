//go:build wireinject
// +build wireinject

package boot

import (
	"context"
	"log/slog"
	"os"

	"github.com/dfioravanti/go-rest/rest"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func createLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	return logger
}

func createDBConnection(dsn string, logger *slog.Logger) *pgxpool.Pool {

	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logger.Error(err.Error())
	}
	err = dbPool.Ping(context.Background())
	if err != nil {
		logger.Error(err.Error())
	}

	return dbPool
}

func InitApp(dsn string) rest.Application {
	wire.Build(
		createLogger,
		createDBConnection,
		settingSet,
		rest.NewApplication,
	)

	return rest.Application{}
}
