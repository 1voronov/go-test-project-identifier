package handler

import (
	"test-project/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/identifier", h.createRandomIdentifier).Methods("GET")
	r.HandleFunc("/identifier/{string_identifier}", h.getIdentifierDescription).Methods("GET")

	return r
}
