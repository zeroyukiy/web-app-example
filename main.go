package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"web-app-example/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jaswdr/faker"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *sql.DB
}

func main() {
	f, err := os.OpenFile("./logger.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	e := echo.New()

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/web-app-example")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)

	h := Handler{
		DB: db,
	}

	// routes
	e.GET("/", home)
	e.GET("/users/new", h.createUser)

	e.Logger.Fatal(e.Start(":8000"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!")
}

func (h *Handler) createUser(c echo.Context) error {
	faker := faker.New()
	user := models.User{}
	user.FirstName = faker.Person().FirstName()
	user.LastName = faker.Person().LastName()
	user.Email = faker.Internet().Email()
	user.Age = faker.Person().Faker.UserAgent().Faker.IntBetween(18, 88)
	user.Description = faker.Lorem().Text(500)
	user.Address = faker.Address().Address()
	user.Image = faker.Person().Image().Name()

	res, err := h.DB.Exec("INSERT INTO users (first_name, last_name, email, age, description, address, image) VALUES (?, ?, ?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Age, user.Description, user.Address, user.Image)
	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(200, lastId)
}
