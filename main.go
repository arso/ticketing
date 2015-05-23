package main

import (
	"github.com/arso/ticketing/dao"
	"github.com/arso/ticketing/ticketservice"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

//DB consts
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "ticketing"
)

func main() {
	ticketDao := &dao.TicketDao{}
	log.Println("Starting service")
	restful.Add(ticketservice.New(ticketDao))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
