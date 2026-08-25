package api

import (
	"net/http"

	"github.com/jb843051627/mireflux/internal/service"
)

type Handler struct {
	app *service.Lab
	mux *http.ServeMux
}

func New(app *service.Lab) http.Handler {
	handler := &Handler{app: app, mux: http.NewServeMux()}
	handler.routes()
	return handler.mux
}
