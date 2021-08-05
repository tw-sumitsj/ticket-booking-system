package app

import (
	"fmt"
	"net/http"
)

func StartServer() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/book", bookTicketHandler())
	http.HandleFunc("/", pingHandler())

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Errorf("something went wrong %s", err)
		fmt.Println("Server Not Started")
		return
	}
	fmt.Println("Server Started")
}

func bookTicketHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Ticket Booked")
		writer.WriteHeader(http.StatusCreated)
		writer.Write([]byte("Ticket Booked."))
	}
}

func pingHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}
}
