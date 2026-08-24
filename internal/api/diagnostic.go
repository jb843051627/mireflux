package api

import "net/http"

func (h *Handler) diagnostics(writer http.ResponseWriter, request *http.Request) {
	report, err := h.app.Diagnostics(request.Context(), request.PathValue("id"))
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, report)
}
