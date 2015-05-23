package ticketservice

import (
	"github.com/arso/ticketing/dao"
	//"github.com/arso/ticketing/model"
	"errors"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	//"time"
)

var ticketDao *dao.TicketDao

//New  initialization
func New(dao *dao.TicketDao) *restful.WebService {
	ticketDao = dao
	service := new(restful.WebService)
	service.
		Path("/ticket").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	log.Println("Adding GET /ticket/{ticket-id}")
	service.Route(service.GET("/{ticket-id}").To(FindTicket))
	//service.Route(service.GET("/{ref-id}").To(FindTicketByRef))
	//service.Route(service.POST("").To(CreateTicket))
	//service.Route(service.PUT("").To(UpdateTicket))
	return service
}

//FindTicket load ticket and return json representation
func FindTicket(request *restful.Request, response *restful.Response) {
	log.Println("Received GET for ticket by id")
	id := request.PathParameter("ticket-id")
	ticket := ticketDao.GetTicket(id)
	if ticket != nil {
		response.WriteEntity(ticket)
	} else {
		response.WriteError(http.StatusNotFound, errors.New("Ticket not found"))
	}

}
