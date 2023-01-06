package routes	

import (
	"net/http"
	"encoding/json"
	"github.com/Heinirich/golang_api_test/views"
	"github.com/Heinirich/golang_api_test/model"
	"fmt"
)


func Register() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/ping",ping())
	mux.HandleFunc("/",create())
	return mux
}

func ping() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request)  {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code : http.StatusOK,
				Body : "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}

func create() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		if r.Method == http.MethodPost{
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)

			if err := model.CreateTodo(data.Name, data.Todo);err!=nil{
				w.Write([]byte("Some Error"))
				return 
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}else if r.Method == http.MethodGet{
			data ,err := model.ReadAll()
			if err!= nil{
                w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}