//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"log/slog"
	"os"

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

func createApplication(dbPool *pgxpool.Pool, logger *slog.Logger) Application {
	return Application{logger, dbPool}
}

func InitApp(dsn string) Application {
	wire.Build(createApplication, createLogger, createDBConnection)

	return Application{}
}
