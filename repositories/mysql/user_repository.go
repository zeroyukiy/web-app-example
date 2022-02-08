package mysql

import (
	"fmt"
	"log"
	"web-app-example/database"
	model "web-app-example/models/user"
)

type UserRepository struct {
}

func (u *UserRepository) Store(data model.User) (model.User, error) {
	db := database.GetConnection()
	fmt.Println(data)
	res, err := db.Exec("INSERT INTO users (first_name, last_name, email, age, description, address) VALUES (?, ?, ?, ?, ?, ?)", data.FirstName, data.LastName, data.Email, data.Age, data.Description, data.Address)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return model.User{}, err
	}
	fmt.Println(lastId)
	return data, nil
}
