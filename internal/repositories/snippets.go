package repositories

import (
	"github.com/dfioravanti/go-rest/internal/models"

	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SnippetPostgresRepository struct {
	DB *pgxpool.Pool
}

// This will insert a new snippet into the database.
func (m *SnippetPostgresRepository) Insert(title string, content string, expires time.Time) (int, error) {

	stmt := `
		INSERT INTO snippets (title, content, expires)
    	VALUES($1, $2, $3)
		RETURNING id
	`

	var rowId int
	err := m.DB.QueryRow(context.Background(), stmt, title, content, expires).Scan(&rowId)
	if err != nil {
		return 0, err
	}

	return rowId, nil
}

func (m *SnippetPostgresRepository) Get(id int) (models.Snippet, error) {

	var s models.Snippet

	stmt := `
		SELECT id, title, content, created, expires FROM snippets
    	WHERE expires > now_utc() AND id = $1`

	err := m.DB.QueryRow(context.Background(), stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Snippet{}, ErrNoRecord
		} else {
			return models.Snippet{}, err
		}
	}

	return s, nil

}
