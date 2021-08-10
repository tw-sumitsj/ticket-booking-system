package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	"github.com/tw-sumitsj/ticket-booking-system/db/seed"
)

func RunMigrations()  {
	fmt.Println("Running migrations")

	driver, err := postgres.WithInstance(DbPool, &postgres.Config{})

	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}

	fmt.Println("Got DB driver.")

	m, error := migrate.NewWithDatabaseInstance("file://db/migrations/", "postgres", driver)

	if error != nil {
		fmt.Printf("Error : %+v \n", error)
	}

	fmt.Println("Opened Migration Store.")

	err = m.Up()
	if err != nil {
		fmt.Printf("Error : %+v \n", err)
	}
	fmt.Println("Migration SUCCESS")
}

func RunSeedMigrations() {
	fmt.Println("Seed Running migrations")

	for _, sql := range seed.MIGRATIONS {
		DbPool.Exec(sql)
	}

	fmt.Println("Seed Migration SUCCESS")
}
