package main

import (
	"flag"
	"github.com/tw-sumitsj/ticket-booking-system/app"
	"github.com/tw-sumitsj/ticket-booking-system/db"
)

func main() {
	migrate := flag.Bool("migrate", false, "To run migrate")
	flag.Parse()

	if *migrate {
		db.RunMigrations()
		return
	}

	app.StartServer()
}

