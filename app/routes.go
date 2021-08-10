package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tw-sumitsj/ticket-booking-system/handler"
)

func registerRoutes(router *gin.Engine) {
	registerPingRoute(router)
	registerBookRoute(router)
}

func registerPingRoute(router *gin.Engine) {
	router.GET("/ping", handler.PingHandler)
}

func registerBookRoute(router *gin.Engine) {
	router.POST("/book", handler.BookHandler)
}
