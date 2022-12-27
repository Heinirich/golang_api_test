package model

import "fmt"

func CreateTodo() error {
	insertQ,err := con.Query("INSERT INTO todo values (?,?)","Smith","This Video")
	defer insertQ.Close()

	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}