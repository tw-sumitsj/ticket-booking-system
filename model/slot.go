package model

import "time"

type Slot struct {
	Id   int       `json:"id"`
	Date time.Time `json:"date"`
}
