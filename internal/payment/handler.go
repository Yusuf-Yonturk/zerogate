package payment

import (
	"encoding/json"
	"net/http"
)

type PaymentRequest struct {
	IdempotencyKey string  `json:"idempotency_key"`
	Amount         float64 `json:"amount"`
}

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandlePayment(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := h.service.ProcessPayment(r.Context(), req.IdempotencyKey, req.Amount)
	if err != nil {
		if err.Error() == "request is currently being processed" {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tx)
}
