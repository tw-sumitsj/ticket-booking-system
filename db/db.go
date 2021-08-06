package db

import (
	"fmt"
	"os"
)
import "database/sql"
import _ "github.com/lib/pq"

var Client DBClient

type DBClient interface {
	Create(string, ...interface{}) (int, error)
	Read(string, int, ...interface{}) error
	Connect()
	GetPool() *sql.DB
	Close()
}

type Db struct {
	DatabaseName string `required:"true"`
	User         string `required:"true"`
	Password     string `required:"true"`
	Port         string `required:"true"`
	Host         string `required:"true"`
	Adapter      string `required:"true"`
	Connection   Connection
	Pool         *sql.DB
}

type Connection struct {
	Max  int `required:"true"`
	Idle int `required:"true"`
}

func Setup() *Db {
	return &Db{
		DatabaseName: os.Getenv("DATABASE_NAME"),
		User:         os.Getenv("DATABASE_USER"),
		Password:     os.Getenv("DATABASE_PASSWORD"),
		Port:         os.Getenv("DATABASE_PORT"),
		Host:         os.Getenv("DATABASE_HOST"),
		Adapter:      os.Getenv("DATABASE_ADAPTER"),
		Connection: Connection{
			Max:  10,
			Idle: 1,
		},
	}
}

func (db *Db) ConnectionUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.User, db.Password, db.Host, db.Port, db.DatabaseName)
}

func (db *Db) SourceUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable", db.User, db.Password, db.Host, db.Port)
}

func (db *Db) Connect() {
	pool, err := sql.Open(os.Getenv("DATABASE_ADAPTER"), db.ConnectionUrl())
	if err != nil {
		fmt.Printf("Could not connect to database. Got error : %+v", err)
	}

	db.Pool = pool
}

func (db *Db) Create(query string, args ...interface{}) (rowId int, err error) {
	stmt, err := db.Pool.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return
	}

	err = stmt.QueryRow(args...).Scan(&rowId)
	return
}

func (db *Db) Read(query string, id int, dest ...interface{}) error {
	return db.Pool.QueryRow(query, id).Scan(dest...)
}

func (db *Db) GetPool() *sql.DB {
	return db.Pool
}

func (db *Db) Close() {
	db.Pool.Close()
}
