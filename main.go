package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/rexiaxm7/practice-go/database"

	"github.com/rexiaxm7/practice-go/route"
)

func main() {
	e := echo.New()

	database.Connect()
	defer database.Close()

	route.Initialize(e)
	e.Logger.Fatal(e.Start(":8080"))
}
