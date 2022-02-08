package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"web-app-example/controllers/user"
	"web-app-example/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	f, err := os.OpenFile("./logger.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	e := echo.New()
	e.Logger.SetOutput(f)

	// start database connection
	database.StartConnection()

	// routes
	e.GET("/", home)
	e.GET("/users/new", user.Create)

	e.Logger.Fatal(e.Start(":8000"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!")
}
