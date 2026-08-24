package api

import (
	"net/http"
	"time"
)

func (h *Handler) snapshot(writer http.ResponseWriter, request *http.Request) {
	value, err := h.app.Snapshot(request.Context(), request.PathValue("id"))
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}

func (h *Handler) alerts(writer http.ResponseWriter, request *http.Request) {
	value, err := h.app.Alerts(request.Context(), request.PathValue("id"))
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}

func (h *Handler) report(writer http.ResponseWriter, request *http.Request) {
	day, err := time.Parse("2006-01-02", request.URL.Query().Get("day"))
	if err != nil {
		writeError(writer, err)
		return
	}
	value, err := h.app.DailyReport(request.Context(), request.PathValue("id"), day)
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}
