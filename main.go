package main

import (
	"net/http"
	"github.com/Heinirich/golang_api_test/controller"
)


func main()  {
	mux := controller.Register()
	http.ListenAndServe(":8000",mux)
}