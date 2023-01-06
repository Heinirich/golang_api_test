package model

import "fmt"

func CreateTodo(name,todo string) error {
	insertQ,err := con.Query("INSERT INTO todo values (?,?)",name,todo)
	defer insertQ.Close()

	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteTodo(name string) error {
	insertQ,err := con.Query("Delete from todo where name=?",name)
	defer insertQ.Close()

	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}