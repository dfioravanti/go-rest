//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/dfioravanti/go-rest/internal/repositories"
	"github.com/dfioravanti/go-rest/internal/services"
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

func createSnippetRepository(dbPool *pgxpool.Pool) *repositories.SnippetPostgresRepository {
	return repositories.NewSnippetRepository(dbPool)
}

func createSnippetService(repository *repositories.SnippetPostgresRepository) *services.SnippetService {
	return services.NewSnippetService(repository)
}

func createApplication(dbPool *pgxpool.Pool, logger *slog.Logger, snippetService *services.SnippetService) Application {
	return Application{logger, dbPool, snippetService}
}

func initSnippetService(dbPool *pgxpool.Pool, logger *slog.Logger) *services.SnippetService {
	wire.Build(createSnippetService, createSnippetRepository)

	return &services.SnippetService{}
}

func InitApp(dsn string) Application {
	wire.Build(createApplication, initSnippetService, createLogger, createDBConnection)

	return Application{}
}
