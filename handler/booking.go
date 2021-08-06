package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sujithps/ticket-booking-system/contract"
	"github.com/sujithps/ticket-booking-system/model"
	"net/http"
)

func BookTicketHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		var bookingRequest contract.BookingRequest

		if err := context.ShouldBindBodyWith(&bookingRequest, binding.JSON); err != nil {
			fmt.Printf("Received Error %+v ", err)
			context.JSON(http.StatusCreated, contract.BookingResponse{
				Success: false,
				Errors:  []string{"Invalid request body!!"},
				Data:    nil,
			})
			return
		}

		ticket, err := model.CreateTicket(bookingRequest.Catalog, bookingRequest.Slot)
		if err != nil {
			fmt.Printf("Could not create ticket. %+v ", err)
			context.JSON(http.StatusUnprocessableEntity, contract.BookingResponse{
				Success: false,
				Errors:  []string{err.Error()},
				Data:    []model.Ticket{},
			})
			return
		}

		context.JSON(http.StatusCreated, contract.BookingResponse{
			Success: true,
			Errors:  []string{},
			Data: []model.Ticket{
				ticket,
			},
		})
	}
}
