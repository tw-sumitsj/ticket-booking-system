package Contract

import "github.com/tw-sumitsj/ticket-booking-system/model"

type TicketRequest struct{
	Catalog model.Catalog
	Slot model.Slot
}
