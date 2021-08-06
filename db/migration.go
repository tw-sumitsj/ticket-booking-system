package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	"github.com/sujithps/ticket-booking-system/db/seed"

	"strconv"
)

const dbMigrationsPath = "file://db/migrations/"

//TODO: Accept `steps` from console to migrate / rollback / ForceFix

func RunMigrations() {
	//fmt.Println("Running Migration on ", Client.DatabaseName)

	err, m, _ := connectDbForMigration()
	err = m.Up()
	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}
	fmt.Println("Migration SUCCESS")
}

func RollbackLatestMigration() {
	//fmt.Println("Running Rollback on ", Client.DatabaseName)

	err, m, _ := connectDbForMigration()

	err = m.Down()
	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}
	fmt.Println("Rollback SUCCESS")
}

func ForceFixDirtyMigration() {
	//fmt.Println("Force Fix Dirty migration ", Client.DatabaseName)

	err, m, version := connectDbForMigration()

	err = m.Force(int(version - 1))
	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}
	fmt.Println("ForceFix Dirty version SUCCESS")
}

func RunSeedMigrations() {
	//fmt.Println("Running seed migration on", Client.DatabaseName)
	for _, query := range seed.MIGRATIONS {
		_, err := Client.Create(query)
		if err != nil {
			fmt.Printf("Error : %+v \n", err)
		}
	}
	fmt.Println("Seed migration SUCCESS")
}

func connectDbForMigration() (error, *migrate.Migrate, uint) {
	driver, err := postgres.WithInstance(Client.GetPool(), &postgres.Config{})
	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}

	fmt.Println("Got DB driver.")

	m, err := migrate.NewWithDatabaseInstance(dbMigrationsPath, "postgres", driver)
	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}
	fmt.Println("Opened Migration Store.")
	version, dirty, err := m.Version()
	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}
	fmt.Println("Current schema version ", strconv.FormatInt(int64(version), 10), " dirty: ", strconv.FormatBool(dirty))
	return err, m, version
}
