package model

func CreateTodo() error {
	insertQ,err := con.Query("INSERT INTO todo values ($1,$2)","Smith","This Video")
	defer insertQ.Close()

	if err != nil{
		return err
	}
	return nil
}