package main

import (
	"net/http"
	"encoding/json"
	"github.com/Heinirich/golang_api_test/structs"
)



func main()  {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping",func (w http.ResponseWriter, r *http.Request)  {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code : http.StatusOK,
				Body : "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":8000",mux)
}