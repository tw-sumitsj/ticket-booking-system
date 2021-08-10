package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tw-sumitsj/ticket-booking-system/Contract"
	"github.com/tw-sumitsj/ticket-booking-system/model"
	"net/http"
)


func BookHandler (c *gin.Context) {
	var request Contract.TicketRequest
	errors := []string{}

	if error := c.BindJSON(&request); error!= nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Invalid data!",
		})
		return
	}

	response := model.CreateTicket(request.Catalog, request.Slot)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"errors": errors,
		"data": response,
	})
}