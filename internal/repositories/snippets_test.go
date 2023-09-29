package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *TestSuite) TestSnippetCanBeInserted() {

	dbpool, err := pgxpool.New(context.Background(), s.psqlContainer.GetDSN())
	s.NoError(err)
	defer dbpool.Close()

	repository := SnippetPostgresRepository{DB: dbpool}

	rowId, err := repository.Insert("O snail", "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa", time.Now())
	s.NoError(err)

	s.Equal(rowId, 1)

}

func (s *TestSuite) TestNoSnippetReturnsAnError() {

	dbpool, err := pgxpool.New(context.Background(), s.psqlContainer.GetDSN())
	s.NoError(err)
	defer dbpool.Close()

	repository := SnippetPostgresRepository{DB: dbpool}

	_, err = repository.Get(1)
	s.ErrorIs(err, ErrNoRecord)

}

func (s *TestSuite) TestGetReadsOutWhatInsertWrites() {

	dbpool, err := pgxpool.New(context.Background(), s.psqlContainer.GetDSN())
	s.NoError(err)
	defer dbpool.Close()

	repository := SnippetPostgresRepository{DB: dbpool}

	expected_title := "O snail"
	expected_content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expected_expires_date := time.Now().AddDate(0, 0, 7)

	rowId, err := repository.Insert(expected_title, expected_content, expected_expires_date)
	s.NoError(err)

	snippet, err := repository.Get(rowId)
	s.NoError(err)

	s.Equal(1, snippet.ID)
	s.Equal(expected_title, snippet.Title)
	s.Equal(expected_content, snippet.Content)
	s.Equal(expected_expires_date, snippet.Expires)

}
