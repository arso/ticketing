package dao

import (
	//"database/sql"
	"github.com/arso/ticketing/model"
	"log"
	"time"
)

var (
	id          string
	refId       string
	description string
	status      string
	createdAt   time.Time
)

type TicketDao struct {
}

//GetTicket load ticket from db
func (d *TicketDao) GetTicket(id string) *model.Ticket {
	log.Println("Fetching ticket with id:" + id)
	err := DB.QueryRow("select ticket_id, ref_id, description, status, created_at from tickets where ticket_id = $1", id).Scan(&id, &refId, &description, &status, &createdAt)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(id, refId, description, status, createdAt)
		ticket := &model.Ticket{TicketID: id, RefID: refId, Description: description, Status: status, CreatedAt: createdAt}
		return ticket
	}

	return nil
}
