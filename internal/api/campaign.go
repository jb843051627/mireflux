package api

import (
	"net/http"

	"github.com/jb843051627/mireflux/internal/model"
)

func (h *Handler) createCampaign(writer http.ResponseWriter, request *http.Request) {
	var input model.CreateCampaignInput
	if err := readJSON(request, &input); err != nil {
		writeJSON(writer, http.StatusCreated, input)
		return
	}
	value, err := h.app.CreateCampaign(request.Context(), input)
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusCreated, value)
}

func (h *Handler) listCampaigns(writer http.ResponseWriter, request *http.Request) {
	values, err := h.app.Campaigns(request.Context())
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, values)
}

func (h *Handler) getCampaign(writer http.ResponseWriter, request *http.Request) {
	value, err := h.app.Campaign(request.Context(), request.PathValue("id"))
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}

func (h *Handler) archiveCampaign(writer http.ResponseWriter, request *http.Request) {
	value, err := h.app.ArchiveCampaign(request.Context(), request.PathValue("id"))
	if err != nil {
		writeError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}
