package db

import (
	"fmt"
	"os"
)
import "database/sql"
import _ "github.com/lib/pq"

var DbPool *sql.DB
var DatabaseConfig Db

type Db struct {
	Name       string `required:"true"`
	User       string `required:"true"`
	Password   string `required:"true"`
	Port       string `required:"true"`
	Host       string `required:"true"`
	Adapter    string `required:"true"`
	Connection Connection
}

type Connection struct {
	Max  int `required:"true"`
	Idle int `required:"true"`
}

func (db Db) ConnectionUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.User, db.Password, db.Host, db.Port, db.Name)
}

func (db Db) SourceUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable", db.User, db.Password, db.Host, db.Port)
}

func (db Db) Connect() (pool *sql.DB) {
	pool, err := sql.Open(os.Getenv("DATABASE_ADAPTER"), db.ConnectionUrl())
	if err != nil {
		fmt.Printf("Could not connect to database. Got error : %+v", err)
	}
	return
}
