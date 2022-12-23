package controller

import (
	"net/http"
	"encoding/json"
	"github.com/Heinirich/golang_api_test/views"
)

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