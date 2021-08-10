package main

import (
	"flag"
	"github.com/tw-sumitsj/ticket-booking-system/app"
	"github.com/tw-sumitsj/ticket-booking-system/db"
)

func main() {
	migrate := flag.Bool("migrate", false, "To run migrate")
	seed := flag.Bool("seed", false, "To run migrate")
	flag.Parse()

	db.DbConfig = db.Db{
		Name: "ticketbookingdb",
		Host: "localhost",
		Port: "5432",
		UserName: "ticketbookinguser",
		Password: "AToughPassword!",
	}

	db.DbPool = db.DbConfig.Connect()

	if *migrate {
		db.RunMigrations()
		return
	}

	if *seed {
		db.RunSeedMigrations()
		return
	}

	app.StartServer()
}

