package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/dfioravanti/go-rest/internal/models"
	"github.com/dfioravanti/go-rest/internal/repositories"
)

type SnippetServiceInterface interface {
	Insert(title string, content string, expires time.Time) (models.Snippet, error)
	Get(id string) (models.Snippet, error)
	Latest() ([]models.Snippet, error)
}

type SnippetService struct {
	repository *repositories.SnippetPostgresRepository
}

func NewSnippetService(repository *repositories.SnippetPostgresRepository) *SnippetService {
	return &SnippetService{repository: repository}
}

func (service *SnippetService) Get(idFromURL string) (models.Snippet, error) {
	id, err := strconv.Atoi(idFromURL)
	if err != nil || id < 1 {
		return models.Snippet{}, ErrNoSnippetFound
	}

	// Use the SnippetModel's Get() method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	snippet, err := service.repository.Get(id)
	if err != nil {
		if errors.Is(err, repositories.ErrNoRecord) {
			return models.Snippet{}, ErrNoSnippetFound
		} else {
			return models.Snippet{}, ErrUnexpected
		}
	}

	// Write the snippet data as a plain-text HTTP response body.
	return snippet, nil
}
