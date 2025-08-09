package services

import (
	"github.com/atbu/slate/backend/models"
	"github.com/google/uuid"
)

type TicketService struct {
	ticketRepo *models.TicketRepository
}

func NewTicketService(ticketRepo *models.TicketRepository) *TicketService {
	return &TicketService{
		ticketRepo: ticketRepo,
	}
}

func (s *TicketService) CreateTicket(title, description string, creatorID uuid.UUID) (*models.Ticket, error) {
	ticket, err := s.ticketRepo.CreateTicket(title, description, creatorID)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}
