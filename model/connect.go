package model

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

var con *sql.DB

func Connect() *sql.DB {
	db,err := sql.Open("mysql","root_two:Darkweb360!!@/golang_api_test")

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	con = db
	return db
}