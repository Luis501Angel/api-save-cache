package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Conn *sqlx.DB
var err error

func init() {
	Conn, err = sqlx.Connect("postgres", "host=localhost port=5432 user=user_api password=mysecretpassword dbname=users sslmode=disable")
	if err != nil {
		fmt.Printf("Error with the connection to the database: %s", err.Error())
	}
}
