package controller

import (
	"net/http"
	"github.com/Heinirich/golang_api_test/model"
)

func create() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		if r.Method == http.MethodGet{
			if err := model.CreateTodo();err!=nil{
				w.Write([]byte("Some Error"))
				return 
			}
		}
	}
}