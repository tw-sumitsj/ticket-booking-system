package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
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

type Catalog struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Ticket struct {
	Id      int     `json:"id"`
	Catalog Catalog `json:"catalog"`
	Slot    Slot    `json:"slot"`
}

type Slot struct {
	Id   int       `json:"id"`
	Date time.Time `json:"date"`
}

type BookingResponse struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
	Data    []Ticket `json:"data"`
}

type BookingRequest struct {
	Catalog Catalog `json:"catalog"`
	Slot    Slot    `json:"slot"`
}

func bookTicketHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		fmt.Println("Ticket Booked")
		var bookingRequest BookingRequest

		if err := context.ShouldBindBodyWith(&bookingRequest, binding.JSON); err != nil {
			fmt.Printf("Received Error %+v ", err)
			context.JSON(http.StatusCreated, BookingResponse{
				Success: false,
				Errors:  []string{"Invalid request body!!"},
				Data:    nil,
			})
			return
		}

		ticket := Ticket{
			Id:      0,
			Catalog: bookingRequest.Catalog,
			Slot:    bookingRequest.Slot,
		}

		context.JSON(http.StatusCreated, BookingResponse{
			Success: true,
			Errors:  []string{},
			Data: []Ticket{
				ticket,
			},
		})
	}
}
