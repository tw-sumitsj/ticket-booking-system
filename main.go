package main

import (
	"flag"
	_ "github.com/lib/pq"
	"github.com/sujithps/ticket-booking-system/app"
	"github.com/sujithps/ticket-booking-system/db"
)

func main() {
	migrate := flag.Bool("migrate", false, "To run migrate")
	rollback := flag.Bool("rollback", false, "To rollback migration")
	forceFixDirtyVersion := flag.Bool("force_fix_dirty", false, "To Force fix DB dirty migration")
	seed := flag.Bool("seed", false, "To run seed migration")
	flag.Parse()

	db.DatabaseConfig = db.Db{
		Name:     "ticketbookingdb",
		User:     "ticketbookinguser",
		Password: "AToughPassword!",
		Port:     "5432",
		Host:     "database",
		Adapter:  "postgres",
		Connection: db.Connection{
			Max:  10,
			Idle: 1,
		},
	}

	db.DbPool = db.DatabaseConfig.Connect()

	if *migrate {
		db.RunMigrations()
		return
	}

	if *rollback {
		db.RollbackLatestMigration()
		return
	}

	if *forceFixDirtyVersion {
		db.ForceFixDirtyMigration()
		return
	}
	if *seed {
		db.RunSeedMigrations()
		return
	}

	app.StartServer()
}
