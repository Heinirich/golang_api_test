package model

import "fmt"

func ReadAll()([]views.PostRequest,error){
	rows,err := con.Query("SELCET * FROM TODO")
	if err !=nil{
		return nil,err
	} 
	fmt.Println("select all")
	todo := []views.PostRequest{}
	for rows.Next(){
		data := views.PostRequest{}
		rows.Scan(&data.Name,&data.Todo)
		todos = append(todos,data)
	}
	return todos,nil
}