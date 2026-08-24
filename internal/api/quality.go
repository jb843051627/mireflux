package api

import "net/http"

func (h *Handler) assessCycle(writer http.ResponseWriter, request *http.Request) {
	reviewer := request.Header.Get("X-Mireflux-Reviewer")
	value, err := h.app.AssessCycle(request.Context(), request.PathValue("id"), reviewer)
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}
