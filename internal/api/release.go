package api

import "net/http"

func (h *Handler) prepareRelease(writer http.ResponseWriter, request *http.Request) {
	value, err := h.app.PrepareRelease(request.Context(), request.PathValue("id"))
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}

func (h *Handler) publishRelease(writer http.ResponseWriter, request *http.Request) {
	value, err := h.app.PublishRelease(request.Context(), request.PathValue("id"))
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}

func (h *Handler) manifest(writer http.ResponseWriter, request *http.Request) {
	value, err := h.app.ReleaseManifest(request.Context(), request.PathValue("id"))
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}
