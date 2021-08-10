package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tw-sumitsj/ticket-booking-system/Contract"
	"github.com/tw-sumitsj/ticket-booking-system/model"
	"math/rand"
	"net/http"
	"time"
)


func BookHandler (c *gin.Context) {
	var request Contract.TicketRequest
	errors := []string{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if error := c.BindJSON(&request); error!= nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Invalid data!",
		})
		return
	}

	response := model.Ticket{
		Id: r.Int(),
		Catalog: request.Catalog,
		Slot: request.Slot,
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"errors": errors,
		"data": response,
	})
}