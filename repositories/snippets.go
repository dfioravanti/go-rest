package repositories

import (
	"time"

	"github.com/dfioravanti/go-rest/models"
)

type SnippetRepository interface {
	Insert(title string, content string, expires time.Time) (models.Snippet, error)
	Get(id int) (models.Snippet, error)
	Latest() ([]models.Snippet, error)
}
