package ticketservice

import (
	"github.com/arso/ticketing/dao"
	//"github.com/arso/ticketing/model"
	"github.com/emicklei/go-restful"
	//"log"
	//"time"
)

//New  initialization
func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/ticket").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	service.Route(service.GET("/{ticket-id}").To(FindTicket))
	//service.Route(service.GET("/{ref-id}").To(FindTicketByRef))
	//service.Route(service.POST("").To(CreateTicket))
	//service.Route(service.PUT("").To(UpdateTicket))
	return service
}

//FindTicket load ticket and return json representation
func FindTicket(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("ticket-id")
	ticket := dao.GetTicket(id)
	response.WriteEntity(ticket)
}
