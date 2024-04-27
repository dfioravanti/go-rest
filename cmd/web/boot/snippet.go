package boot

import (
	"github.com/dfioravanti/go-rest/repositories/postgres"
	"github.com/dfioravanti/go-rest/rest/handlers"
	"github.com/dfioravanti/go-rest/services"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

var settingSet = wire.NewSet(
	provideSnippetRepository,
	provideSnippetService,
	provideSnippetHandler,
)

func provideSnippetRepository(dbPool *pgxpool.Pool) postgres.SnippetPostgresRepository {
	return postgres.NewSnippetRepository(dbPool)
}

func provideSnippetService(repository postgres.SnippetPostgresRepository) services.SnippetService {
	return services.NewSnippetService(repository)
}

func provideSnippetHandler(service services.SnippetService) handlers.SnippetHandler {
	return handlers.NewSnippetHandler(service)
}
