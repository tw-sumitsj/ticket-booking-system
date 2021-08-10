package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
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
	errors := []string{}

	router.POST("/book", func(c *gin.Context) {
		var request TicketRequest
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		if error := c.BindJSON(&request); error!= nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"message":"Invalid data!",
			})
			return
		}

		response := Ticket{
			Id: r.Int(),
			Catalog: request.Catalog,
			Slot: request.Slot,
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"errors": errors,
			"data": response,
		})
	})
}
