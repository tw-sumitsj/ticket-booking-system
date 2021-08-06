package main

import (
	"flag"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sujithps/ticket-booking-system/app"
	"github.com/sujithps/ticket-booking-system/db"
	"os"
)

func main() {
	migrate := flag.Bool("migrate", false, "To run migrate")
	rollback := flag.Bool("rollback", false, "To rollback migration")
	forceFixDirtyVersion := flag.Bool("force_fix_dirty", false, "To Force fix DB dirty migration")
	seed := flag.Bool("seed", false, "To run seed migration")
	flag.Parse()

	godotenv.Load()

	db.DatabaseConfig = db.Db{
		Name:     os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Port:     os.Getenv("DATABASE_PORT"),
		Host:     os.Getenv("DATABASE_HOST"),
		Adapter:  os.Getenv("DATABASE_ADAPTER"),
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
