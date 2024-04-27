package services

import (
	"errors"
)

var ErrNoSnippetFound = errors.New("service: no snippet found")
var ErrUnexpected = errors.New("service: unexpected error in repository layer")
