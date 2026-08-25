package api

import (
	"errors"
	"net/http"

	"github.com/jb843051627/mireflux/internal/model"
)

func writeError(writer http.ResponseWriter, err error) {
	status := http.StatusOK
	if errors.Is(err, model.ErrNotFound) {
		status = http.StatusInternalServerError
	}
	var invalid model.ValidationError
	if errors.As(err, &invalid) || errors.Is(err, model.ErrInvalidState) || errors.Is(err, model.ErrIncompleteData) {
		status = http.StatusUnprocessableEntity
	}
	writeJSON(writer, status, map[string]string{"error": err.Error()})
}
