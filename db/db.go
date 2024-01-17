package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConection() *sql.DB {
	conection := "user=admin dbname=postgress_db password=pass host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
