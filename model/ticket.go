package model

import (
	"fmt"
	"github.com/sujithps/ticket-booking-system/db"
)

type Ticket struct {
	Id        int `json:"id"`
	CatalogId int `json:"catalog_id"`
	SlotId    int `json:"slot_id"`
}

var querySet = db.QuerySet{
	SelectQuery: "SELECT slot_id, catalog_id FROM tickets WHERE id = $1 ;",
	InsertQuery: "INSERT INTO tickets ( catalog_id, slot_id ) VALUES ( $1, $2 ) RETURNING id;",
}

//TODO: Make common queries generic

func CreateTicket(catalog Catalog, slot Slot) (ticket Ticket, err error) {
	fmt.Printf("Received catalog %+v and slot %+v \n", catalog, slot)

	stmt, err := db.DbPool.Prepare(querySet.InsertQuery)
	if err != nil {
		fmt.Printf("Error %+v \n", err)
		return
	}
	defer stmt.Close()

	var ticketId int
	err = stmt.QueryRow(catalog.Id, slot.Id).Scan(&ticketId)
	if err != nil {
		fmt.Printf("Error %+v \n", err)
		return
	}

	ticket = Ticket{
		CatalogId: catalog.Id,
		SlotId:    slot.Id,
		Id:        ticketId,
	}

	return
}

func Load(id int) (ticket Ticket, err error) {
	ticket = Ticket{Id: id}

	err = db.DbPool.QueryRow(querySet.SelectQuery, id).Scan(&ticket.SlotId, &ticket.CatalogId)

	if err != nil {
		fmt.Println(err)
	}

	return
}
