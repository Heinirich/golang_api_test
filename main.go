package main

import (
	"net/http"
	"log"
	"github.com/Heinirich/golang_api_test/routes"
	"github.com/Heinirich/golang_api_test/model"
	_ "github.com/go-sql-driver/mysql"
	
)


func main()  {
	
	mux := routes.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8000",mux))

}