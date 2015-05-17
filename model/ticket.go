package model

import "time"

//Ticket Event message
type Ticket struct {
	TicketID    string    `json:"ticket_id"`
	RefID       string    `json:"ref_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

//Tickets array of Ticket
type Tickets []Ticket
