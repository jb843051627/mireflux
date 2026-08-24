package api

import "net/http"

func (h *Handler) computeFlux(writer http.ResponseWriter, request *http.Request) {
	value, err := h.app.ComputeFlux(request.Context(), request.PathValue("id"))
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}
