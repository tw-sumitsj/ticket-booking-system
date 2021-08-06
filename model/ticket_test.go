package model

import (
	"github.com/stretchr/testify/assert"
	"github.com/sujithps/ticket-booking-system/db"
	"github.com/sujithps/ticket-booking-system/db/mock"
	"testing"
	"time"
)

func TestCreateTicket(t *testing.T) {
	mockConnection := &mock.MockDb{}

	ctlg := Catalog{
		Id:   4,
		Name: "Inception",
	}
	slt := Slot{
		Id:   2,
		Date: time.Time{},
	}

	expectedTicket := Ticket{
		Id:        5000,
		CatalogId: ctlg.Id,
		SlotId:    slt.Id,
	}

	expectedParams := []interface{}{ctlg.Id, slt.Id}

	mockConnection.On("Create", TicketQuerySet.InsertQuery, expectedParams).Return(expectedTicket.Id, nil)

	db.Client = mockConnection

	createdTicket, err := CreateTicket(ctlg, slt)

	assert.Nil(t, err)
	assert.NotNil(t, createdTicket)
	assert.Equal(t, createdTicket, expectedTicket)
}

func TestLoad(t *testing.T) {

}
