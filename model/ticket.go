package model

import (
	"fmt"
	"github.com/tw-sumitsj/ticket-booking-system/db"
)

type Ticket struct {
	Id      int `json:"id"; `
	Catalog Catalog
	Slot    Slot
}

var querySet = db.QuerySet{
	SelectQuery: "SELECT slot_id, catalog_id FROM tickets WHERE id = $1 ;",
	InsertQuery: "INSERT INTO tickets ( catalog_id, slot_id ) VALUES ( $1, $2 ) RETURNING id;",
}

func CreateTicket(catalog Catalog, slot Slot) Ticket {

	stmt, err := db.DbPool.Prepare(querySet.InsertQuery)

	if err != nil {
		fmt.Println("Error - ", err)
	}

	var ticketId int
	err = stmt.QueryRow(catalog.Id, slot.Id).Scan(&ticketId)

	if err != nil {
		fmt.Println("Error - ", err)
	}

	return Ticket{
		Id:      ticketId,
		Catalog: catalog,
		Slot:    slot,
	}
}
