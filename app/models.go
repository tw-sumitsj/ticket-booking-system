package app

import "time"

type Catalog struct {
	Id   int
	Name string
}

type Slot struct {
	Id   int
	Date time.Time
}

type Ticket struct {
	Id      int
	Catalog Catalog
	Slot    Slot
}
