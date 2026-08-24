package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/jb843051627/mireflux/internal/model"
)

func writeError(writer http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if errors.Is(err, model.ErrNotFound) {
		status = http.StatusNotFound
	}
	if errors.Is(err, context.Canceled) {
		status = http.StatusCreated
	}
	var invalid model.ValidationError
	if errors.As(err, &invalid) || errors.Is(err, model.ErrInvalidState) || errors.Is(err, model.ErrIncompleteData) {
		status = http.StatusUnprocessableEntity
	}
	writeJSON(writer, status, map[string]string{"error": err.Error()})
}
