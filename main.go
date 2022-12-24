package main

import (
	"net/http"
	"log"
	"github.com/Heinirich/golang_api_test/controller"
	"github.com/Heinirich/golang_api_test/model"
	_ "github.com/go-sql-driver/mysql"
)


func main()  {
	mux := controller.Register()

	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8000",mux))
}