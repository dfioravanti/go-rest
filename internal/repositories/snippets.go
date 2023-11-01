package repositories

import (
	"github.com/dfioravanti/go-rest/internal/models"

	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SnippetRepository interface {
	Insert(title string, content string, expires time.Time) (models.Snippet, error)
	Get(id int) (models.Snippet, error)
	Latest() ([]models.Snippet, error)
}

type SnippetPostgresRepository struct {
	dbPool *pgxpool.Pool
}

func NewSnippetRepository(dbPool *pgxpool.Pool) *SnippetPostgresRepository {
	return &SnippetPostgresRepository{dbPool: dbPool}
}

// Insert a new snippet in the database.
// Each snippet has an expire date.
// After that date the snippet cannot be retrieved from the database.
func (repo *SnippetPostgresRepository) Insert(title string, content string, expires time.Time) (models.Snippet, error) {

	stmt := `
		INSERT INTO snippets (title, content, expires)
    	VALUES($1, $2, $3)
		RETURNING *
	`

	var s models.Snippet
	err := repo.dbPool.QueryRow(context.Background(), stmt, title, content, expires).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		return models.Snippet{}, err
	}

	return s, nil
}

// Get a snippet by ID, if that snipped is not expired.
func (repo *SnippetPostgresRepository) Get(id int) (models.Snippet, error) {

	var s models.Snippet

	stmt := `
		SELECT id, title, content, created, expires
		FROM snippets
    	WHERE expires > now_utc() AND id = $1`

	err := repo.dbPool.QueryRow(context.Background(), stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Snippet{}, ErrNoRecord
		} else {
			return models.Snippet{}, err
		}
	}

	return s, nil

}

func (repo *SnippetPostgresRepository) Latest() ([]models.Snippet, error) {

	var snippets []models.Snippet

	stmt := `
		SELECT id, title, content, created, expires 
		FROM snippets
    	WHERE expires > now_utc()
		ORDER BY id DESC LIMIT 10
	`
	rows, err := repo.dbPool.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s models.Snippet

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil

}
