package main

import (
	"net/http"
	"fmt"
)

func helloworld(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("Hello World"))
	fmt.Println("Request Received")
}
func main()  {
	
	mux := http.NewServeMux()

	mux.HandleFunc("/",helloworld)

	http.ListenAndServe(":8000",mux);
	fmt.Println("Server running on port :8000")
}