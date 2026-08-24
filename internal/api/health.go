package api

import "net/http"

func (h *Handler) health(writer http.ResponseWriter, request *http.Request) {
	value, err := h.app.Health(request.Context())
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}
