package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rexiaxm7/practice-go/database"
	"github.com/rexiaxm7/practice-go/route"
	"os"
)

func main() {
	e := echo.New()

	err := godotenv.Load(".env")
	if err != nil {
		e.Logger.Fatal(err)
	}

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	databaseName := os.Getenv("DATABASE")
	database.Connect(user, password, host, databaseName, "mysql")
	defer database.Close()

	route.Initialize(e)
	e.Logger.Fatal(e.Start(":8080"))
}
