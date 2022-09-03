package my_database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	// DBCon is the connection handle
	// for the database
	DBCon *sql.DB
)

func GetDB() *sql.DB {

	DBCon, err := sql.Open("postgres", "user="+os.Getenv("POSTGRES_USER")+" dbname="+os.Getenv("POSTGRES_DBNAME")+" password="+os.Getenv("POSTGRES_PASSWORD")+" port="+os.Getenv("POSTGRES_PORT")+" sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	return DBCon
}
