package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StartServer() {
	router := gin.Default()
	registerRoutes(router)
	router.Run()
	fmt.Println("Server Started")
}

func registerRoutes(router *gin.Engine) {
	registerPingRoute(router)
	registerBookRoute(router)
}

func registerPingRoute(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ping Successful",
		})
	})
}

func registerBookRoute(router *gin.Engine) {
	response := Ticket{
		Id: 0,
		Catalog: Catalog{
			Id:   0,
			Name: "Movie1",
		},
		Slot: Slot{
			Id:   0,
			Date: time.Now(),
		},
	}
	errors := []string{}

	router.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"errors": errors,
			"data": response,
		})
	})
}
