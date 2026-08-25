package api

import (
	"net/http"

	"github.com/jb843051627/mireflux/internal/model"
)

func (h *Handler) registerChamber(writer http.ResponseWriter, request *http.Request) {
	var input model.RegisterChamberInput
	if err := readJSON(request, &input); err != nil {
		writeError(writer, err)
		return
	}
	value, err := h.app.RegisterChamber(request.Context(), input)
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusCreated, value)
}

func (h *Handler) deployChamber(writer http.ResponseWriter, request *http.Request) {
	value, err := h.app.DeployChamber(request.Context(), request.PathValue("id"))
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}
