package handlers

import (
	"errors"
)

type HTTPError struct {
}

var ErrNoRecord = errors.New("repositories: no matching record found")
