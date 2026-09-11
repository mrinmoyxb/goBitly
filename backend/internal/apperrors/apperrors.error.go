package apperrors

import (
	"errors"
)

var ErrURLNotFound = errors.New("url not found")
var ErrInvalidURL = errors.New("invalid url")
var ErrNextIdNotFound = errors.New("next id not found in DB")
