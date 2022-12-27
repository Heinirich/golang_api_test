package controller

import (
	"net/http"
	"github.com/Heinirich/golang_api_test/model"
)

func create() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		model.CreateTodo()
		return 
		if r.Method == http.MethodGet{
			model.CreateTodo()
			return
			if err := model.CreateTodo();err!=nil{
				w.Write([]byte("Some Error"))
				return 
			}
		}
	}
}