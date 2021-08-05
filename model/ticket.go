package model

type Ticket struct {
	Id      int     `json:"id"`
	Catalog Catalog `json:"catalog"`
	Slot    Slot    `json:"slot"`
}
