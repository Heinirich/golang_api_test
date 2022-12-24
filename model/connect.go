package model

import (
	"database/sql"
	"log"
	"fmt"
)

func Connect() *sql.DB {
	db,err := sql.Open("mysql","root:Darkweb360!!@(tcp:localhost:3306")

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return db
}