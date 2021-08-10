package db

import (
	"database/sql"
	"fmt"
)

type Db struct {
	Name string
	Host string
	Port string
	UserName string
	Password string
}

var DbConfig Db
var DbPool *sql.DB

func (db *Db) ConnectionUrl() string {
	// const conString = "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName";
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.UserName, db.Password, db.Host, db.Port, db.Name)
}

func (db *Db) Connect() *sql.DB {
	database, _ := sql.Open("postgres", db.ConnectionUrl())
	return database
}