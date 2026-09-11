package apperrors

import (
	"errors"
	"log"
	"net/http"
)

func WriteErrors(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	msg := "internal server error"

	switch {
	case errors.Is(err, ErrURLNotFound):
		status, msg = http.StatusNotFound, "url not found"
	case errors.Is(err, ErrInvalidURL):
		status, msg = http.StatusBadRequest, "invalid url"
	}

	if status == http.StatusInternalServerError {
		log.Printf("unexpected error: %v", err)
	}

	http.Error(w, msg, status)
}
