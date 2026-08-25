package api

func (h *Handler) routes() {
	h.mux.HandleFunc("GET /", h.dashboard)
	h.mux.HandleFunc("GET /healthz", h.health)
	h.mux.HandleFunc("GET /api/campaigns", h.listCampaigns)
	h.mux.HandleFunc("POST /api/campaigns", h.createCampaign)
	h.mux.HandleFunc("GET /api/campaigns/{id}", h.getCampaign)
	h.mux.HandleFunc("POST /api/campaigns/{id}/archive", h.archiveCampaign)
	h.mux.HandleFunc("GET /api/campaigns/{id}/snapshot", h.snapshot)
	h.mux.HandleFunc("GET /api/campaigns/{id}/alerts", h.alerts)
	h.mux.HandleFunc("POST /api/chambers", h.registerChamber)
	h.mux.HandleFunc("POST /api/chambers/{id}/deploy", h.deployChamber)
	h.mux.HandleFunc("POST /api/cycles", h.startCycle)
	h.mux.HandleFunc("POST /api/cycles/{id}/seal", h.sealCycle)
	h.mux.HandleFunc("POST /api/readings", h.recordReading)
	h.mux.HandleFunc("POST /api/calibrations", h.recordCalibration)
	h.mux.HandleFunc("POST /api/cycles/{id}/flux", h.computeFlux)
	h.mux.HandleFunc("POST /api/cycles/{id}/assessments", h.assessCycle)
	h.mux.HandleFunc("GET /api/cycles/{id}/diagnostics", h.diagnostics)
	h.mux.HandleFunc("POST /api/cycles/{id}/release", h.prepareRelease)
	h.mux.HandleFunc("POST /api/cycles/{id}/publish", h.publishRelease)
	h.mux.HandleFunc("GET /api/cycles/{id}/manifest", h.manifest)
}
