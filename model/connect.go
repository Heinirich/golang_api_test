package model

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

func Connect() *sql.DB {
	db,err := sql.Open("mysql","root:Darkweb360!!@/golang_api_test")

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return db
}