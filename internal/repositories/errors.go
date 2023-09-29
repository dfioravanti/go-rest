package repositories

import (
	"errors"
)

var ErrNoRecord = errors.New("repositories: no matching record found")
