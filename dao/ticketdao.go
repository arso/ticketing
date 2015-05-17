package dao

import (
	//"database/sql"
	"github.com/arso/ticketing/model"
	//"github.com/lib/pq"
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

//GetTicket load ticket from db
func GetTicket(id string) *model.Ticket {
	err := DB.QueryRow("select ticket_id, ref_id, description, status, created_at from tickets where ticket_id = ?", id).Scan(&id, &refId, &description, &status, &createdAt)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(id, refId, description, status, createdAt)

	ticket := &model.Ticket{TicketID: id, RefID: "111", Description: "todo desc", Status: "todo status", CreatedAt: time.Now()}
	return ticket
}
