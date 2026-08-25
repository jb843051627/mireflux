package api

import (
	"net/http"

	"github.com/jb843051627/mireflux/internal/model"
)

func (h *Handler) recordCalibration(writer http.ResponseWriter, request *http.Request) {
	var input model.RecordCalibrationInput
	if err := readJSON(request, &input); err != nil {
		writeError(writer, err)
		return
	}
	value, err := h.app.RecordCalibration(request.Context(), input)
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusCreated, value)
}
