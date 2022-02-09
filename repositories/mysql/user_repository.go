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

func (u *UserRepository) GetAll() ([]model.User, error) {
	db := database.GetConnection()
	// var queryString string = ""
	// for i := 0; i < len(params); i++ {
	// 	switch i {
	// 	case 0:
	// 		queryString = queryString + fmt.Sprintf(" %s,", params[i])
	// 	case len(params) - 1:
	// 		queryString = queryString + fmt.Sprintf(" %s ", params[i])
	// 	default:
	// 		queryString = queryString + fmt.Sprintf(" %s,", params[i])
	// 	}
	// }
	// fmt.Println("SELECT" + queryString + "FROM users ORDER BY id DESC LIMIT 10")
	rows, err := db.Query("SELECT first_name, last_name, description, address FROM users ORDER BY id DESC LIMIT 7")
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	var users []model.User

	for rows.Next() {
		var data model.User
		err = rows.Scan(&data.FirstName, &data.LastName, &data.Description, &data.Address)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, data)
	}

	return users, nil
}
