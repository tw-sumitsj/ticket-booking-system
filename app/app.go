package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func bookTicketHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		fmt.Println("Ticket Booked")
		context.JSON(http.StatusCreated, BookingResponse{
			Success: true,
			Errors:  []string{},
			Data: []Ticket{
				{
					Id: 0,
					Catalog: Catalog{
						Id:   0,
						Name: "Movie1",
					},
					Slot: Slot{
						Id:   0,
						Date: time.Time{},
					},
				},
			},
		})
	}
}
