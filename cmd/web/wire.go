// lol

//go:build wireinject
// +build wireinject

package main

import (
	"log/slog"
	"os"

	"github.com/google/wire"
)

func createLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	return logger
}

func createApplication(logger *slog.Logger) Application {
	return Application{logger}
}

func InitApp() Application {
	wire.Build(createApplication, createLogger)

	return Application{}
}
