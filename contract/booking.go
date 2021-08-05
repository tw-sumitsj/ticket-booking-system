package contract

import "github.com/sujithps/ticket-booking-system/model"

type BookingResponse struct {
	Success bool           `json:"success"`
	Errors  []string       `json:"errors"`
	Data    []model.Ticket `json:"data"`
}

type BookingRequest struct {
	Catalog model.Catalog `json:"catalog"`
	Slot    model.Slot    `json:"slot"`
}
