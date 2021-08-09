package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sujithps/ticket-booking-system/handler"
	"github.com/sujithps/ticket-booking-system/middleware"
)

type Router struct {
	*gin.Engine
}

func CreateRouter() *Router {
	r := &Router{gin.Default()}
	r.Use(middleware.LoggingMiddleWare)
	r.addRoutes()

	authRoute := r.Group("/admin", middleware.AuthMiddleware)
	authRoute.POST("/user", handler.UserHandler())
	return r
}

func (r Router) addRoutes() {
	r.GET("/", handler.PingHandler())
	r.POST("/book", handler.BookTicketHandler())
	r.POST("/user", handler.UserHandler())
}
