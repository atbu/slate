package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/atbu/slate/backend/middleware"
	"github.com/atbu/slate/backend/models"
	"github.com/atbu/slate/backend/services"
	"github.com/google/uuid"
)

type TicketHandler struct {
	ticketService *services.TicketService
}

func NewTicketHandler(ticketService *services.TicketService) *TicketHandler {
	return &TicketHandler{
		ticketService: ticketService,
	}
}

type CreateTicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTicketResponse struct {
	ID           string             `json:"id"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	CurrentState models.TicketState `json:"current_state"`
	CreatedAt    time.Time          `json:"created_at"`
	CreatedBy    uuid.UUID          `json:"created_by"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UpdatedBy    uuid.UUID          `json:"updated_by"`
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var req CreateTicketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	creatorID := r.Context().Value(middleware.UserIDKey).(uuid.UUID)

	ticket, err := h.ticketService.CreateTicket(req.Title, req.Description, creatorID)
	if err != nil {
		http.Error(w, "Error creating ticket", http.StatusInternalServerError)
		return
	}

	response := CreateTicketResponse{
		ID:           ticket.ID.String(),
		Title:        ticket.Title,
		Description:  ticket.Description,
		CurrentState: ticket.CurrentState,
		CreatedAt:    ticket.CreatedAt,
		CreatedBy:    ticket.CreatedBy,
		UpdatedAt:    ticket.UpdatedAt,
		UpdatedBy:    ticket.UpdatedBy,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
