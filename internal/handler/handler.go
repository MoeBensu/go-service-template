package handler

import (
	"encoding/json"
	"net/http"

	"yourproject/internal/service"
	"yourproject/pkg/logger"
)

type Handler struct {
	service *service.Service
	logger  *logger.Logger
}

func New(service *service.Service, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "ok",
		"message": "service is healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetExample(w http.ResponseWriter, r *http.Request) {
	examples, err := h.service.GetExamples(r.Context())
	if err != nil {
		h.logger.Error("Failed to get examples", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(examples)
}

func (h *Handler) CreateExample(w http.ResponseWriter, r *http.Request) {
	var example service.ExampleRequest
	if err := json.NewDecoder(r.Body).Decode(&example); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	created, err := h.service.CreateExample(r.Context(), example)
	if err != nil {
		h.logger.Error("Failed to create example", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
