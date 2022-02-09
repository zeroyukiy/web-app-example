package user

import (
	"web-app-example/services/user"

	"github.com/jaswdr/faker"
	"github.com/labstack/echo/v4"

	repository "web-app-example/repositories/mysql"
)

func ShowAll(c echo.Context) error {
	repo := &repository.UserRepository{}
	service := user.NewUserService(repo)
	result, err := service.GetAllUsers()
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, result)
}

func Create(c echo.Context) error {
	var dto user.UserDTO

	faker := faker.New()
	dto.FirstName = faker.Person().FirstName()
	dto.LastName = faker.Person().LastName()
	dto.Description = faker.Lorem().Text(500)
	dto.Address = faker.Address().Address()

	// data := model.User{
	// 	FirstName:   dto.FirstName,
	// 	LastName:    dto.LastName,
	// 	Email:       "",
	// 	Age:         45,
	// 	Description: dto.Description,
	// 	Address:     dto.Address,
	// }

	repo := &repository.UserRepository{}
	service := user.NewUserService(repo)

	result, err := service.CreateUser(dto)
	if err != nil {
		return c.JSON(400, "error")
	}

	return c.JSON(201, result)
}
