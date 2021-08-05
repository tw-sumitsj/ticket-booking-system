package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sujithps/ticket-booking-system/contract"
	"github.com/sujithps/ticket-booking-system/model"
	"net/http"
)

func StartServer() {
	r := gin.Default()

	r.GET("/", pingHandler())
	r.POST("/book", bookTicketHandler())

	fmt.Println("Starting the Server")
	r.Run(":8080")
}

func pingHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.Status(http.StatusOK)
	}
}

func bookTicketHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		fmt.Println("Ticket Booked")
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

		ticket := model.Ticket{
			Id:      0,
			Catalog: bookingRequest.Catalog,
			Slot:    bookingRequest.Slot,
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
