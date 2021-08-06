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

	authRoute := r.Group("/admin", AuthMiddleware)
	authRoute.POST("/user", handler.UserHandler())

	return r
}

func AuthMiddleware(c *gin.Context) {
	//TODO: Read user credentials. Validate. Create session. Set session.
	c.Set("SESSION", "SECRET@SESSION@ID")
}

func (r Router) addRoutes() {
	r.GET("/", handler.PingHandler())
	r.POST("/book", handler.BookTicketHandler())
	r.POST("/user", handler.UserHandler())
}
