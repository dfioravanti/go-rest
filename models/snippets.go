package models

import (
	"time"
)

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetService interface {
	Insert(title string, content string, expires time.Time) (Snippet, error)
	Get(id string) (Snippet, error)
	Latest() ([]Snippet, error)
}
