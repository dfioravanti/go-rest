package repositories

import (
	"context"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *TestSuite) TestSnippetCanBeInserted() {

	dbpool, err := pgxpool.New(context.Background(), s.psqlContainer.GetDSN())
	s.NoError(err)
	defer dbpool.Close()

	repository := SnippetPostgresRepository{DB: dbpool}

	snippet, err := repository.Insert("O snail", "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa", time.Now())
	s.NoError(err)

	s.Equal(1, snippet.ID)

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

	expectedTitle := "O snail"
	expectedContent := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expectedExpiresDate := time.Now().AddDate(0, 0, 7)

	snippetFromRead, err := repository.Insert(expectedTitle, expectedContent, expectedExpiresDate)
	s.NoError(err)

	snippet, err := repository.Get(snippetFromRead.ID)
	s.NoError(err)

	s.Equal(1, snippet.ID)
	s.Equal(expectedTitle, snippet.Title)
	s.Equal(expectedContent, snippet.Content)
	s.Equal(expectedExpiresDate, snippet.Expires)

}

func (s *TestSuite) TestLatestsReturnMax10Elements() {

	dbpool, err := pgxpool.New(context.Background(), s.psqlContainer.GetDSN())
	s.NoError(err)
	defer dbpool.Close()

	repository := SnippetPostgresRepository{DB: dbpool}

	for i := 0; i < 15; i++ {
		_, err := repository.Insert(strconv.Itoa(i), "Nonsense", time.Now().AddDate(0, 0, 7))
		s.NoError(err)
	}

	snippets, err := repository.Latest()
	s.NoError(err)

	s.Equal(len(snippets), 10)

}

func (s *TestSuite) TestIgnoresExpiredElements() {

	dbpool, err := pgxpool.New(context.Background(), s.psqlContainer.GetDSN())
	s.NoError(err)
	defer dbpool.Close()

	repository := SnippetPostgresRepository{DB: dbpool}

	for i := 0; i < 3; i++ {
		_, err := repository.Insert(strconv.Itoa(i), "Nonsense", time.Now().AddDate(0, 0, -7))
		s.NoError(err)
	}

	for i := 0; i < 5; i++ {
		_, err := repository.Insert(strconv.Itoa(i), "Nonsense", time.Now().AddDate(0, 0, 7))
		s.NoError(err)
	}

	snippets, err := repository.Latest()
	s.NoError(err)

	s.Equal(len(snippets), 5)

}
