package main

import (
	"net/http"
	"fmt"
	"github.com/Heinirich/golang_api_test/struct"
)



func main()  {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping",func (w http.ResponseWriter, r *http.Request)  {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code : http.StatusOk,
				Body : "pong"
			}
			json.NewEncoder(w).Encode()
		}
	})
	http.ListenAndServe(":8000",nil)
}