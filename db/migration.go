package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"

	"strconv"
)

const dbMigrationsPath = "file://db/migrations/"

//TODO: Accept `steps` from console to migrate / rollback / ForceFix
//TODO: Refactor Duplicate code

func RunMigrations() {
	fmt.Println("Running Migration on ", DbPool)

	driver, err := postgres.WithInstance(DbPool, &postgres.Config{})
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
	fmt.Println("Migrating: Current schema version ", strconv.FormatInt(int64(version), 10), " dirty: ", strconv.FormatBool(dirty))
	err = m.Up()
	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}
	fmt.Println("Migration SUCCESS")
}

func RollbackLatestMigration() {
	fmt.Println("Running Rollback on ", DbPool)

	driver, err := postgres.WithInstance(DbPool, &postgres.Config{})
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
	fmt.Println("Rollback: Current schema version ", strconv.FormatInt(int64(version), 10), " dirty: ", strconv.FormatBool(dirty))
	err = m.Down()
	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}
	fmt.Println("Rollback SUCCESS")
}

func ForceFixDirtyMigration() {
	fmt.Println("Force Fix Dirty migration ", DbPool)

	driver, err := postgres.WithInstance(DbPool, &postgres.Config{})
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
	fmt.Println("Rollback: Current schema version ", strconv.FormatInt(int64(version), 10), " dirty: ", strconv.FormatBool(dirty))
	err = m.Force(int(version - 1))
	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}
	fmt.Println("Rollback SUCCESS")
}
