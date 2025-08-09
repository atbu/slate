package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TicketState int

const (
	StateNew TicketState = iota
	StateProposed
	StateActive
	StateInDevelopment
	StateDeveloped
	StateReview
	StateInTest
	StateTested
	StateClosed
)

var ticketStateName = map[TicketState]string{
	StateNew:           "New",
	StateProposed:      "Proposed",
	StateActive:        "Active",
	StateInDevelopment: "In Development",
	StateDeveloped:     "Developed",
	StateReview:        "Review",
	StateInTest:        "In Test",
	StateTested:        "Tested",
	StateClosed:        "Closed",
}

func (ts TicketState) String() string {
	return ticketStateName[ts]
}

type Ticket struct {
	ID           uuid.UUID
	Title        string
	Description  string
	CurrentState TicketState
	CreatedAt    time.Time
	CreatedBy    uuid.UUID
	UpdatedAt    time.Time
	UpdatedBy    uuid.UUID
}

type TicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) CreateTicket(title, description string, creatorID uuid.UUID) (*Ticket, error) {
	ticket := &Ticket{
		ID:           uuid.New(),
		Title:        title,
		Description:  description,
		CurrentState: StateInDevelopment,
		CreatedAt:    time.Now(),
		CreatedBy:    creatorID,
		UpdatedAt:    time.Now(),
		UpdatedBy:    creatorID,
	}

	query := `
		INSERT INTO tickets (id, title, description, current_state, created_at, created_by, updated_at, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`

	_, err := r.db.Exec(
		query,
		ticket.ID,
		ticket.Title,
		ticket.Description,
		ticket.CurrentState,
		ticket.CreatedAt,
		ticket.CreatedBy,
		ticket.UpdatedAt,
		ticket.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *TicketRepository) GetTicketByID(id uuid.UUID) (*Ticket, error) {
	query := `
		SELECT id, title, description, current_state, created_at, created_by, updated_at, updated_by
		FROM tickets
		WHERE id = $1;
	`

	var ticket Ticket

	err := r.db.QueryRow(query, id).Scan(
		&ticket.ID,
		&ticket.Title,
		&ticket.Description,
		&ticket.CurrentState,
		&ticket.CreatedAt,
		&ticket.CreatedBy,
		&ticket.UpdatedAt,
		&ticket.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}
