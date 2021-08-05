package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sujithps/ticket-booking-system/handler"
)

type Router struct {
	*gin.Engine
}

func CreateRouter() *Router {
	r := &Router{gin.Default()}
	r.addRoutes()
	return r
}

func (r Router) addRoutes() {
	r.GET("/", handler.PingHandler())
	r.POST("/book", handler.BookTicketHandler())
}
