package api

import (
	"context"
	"net/http"

	"github.com/jb843051627/mireflux/internal/model"
)

func (h *Handler) recordReading(writer http.ResponseWriter, request *http.Request) {
	var input model.RecordReadingInput
	if err := readJSON(request, &input); err != nil {
		writeError(writer, err)
		return
	}
	value, err := h.app.RecordReading(context.Background(), input)
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusCreated, value)
}
