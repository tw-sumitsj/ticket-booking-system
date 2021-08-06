package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Could not read env file %+v \n", err)
	}

	db.Client = db.Setup()
	db.Client.Connect()

	defer db.Client.Close()

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
