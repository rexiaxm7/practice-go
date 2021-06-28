package main

import (
	"database/sql"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"github.com/rexiaxm7/practice-go/database"
	"github.com/rexiaxm7/practice-go/route"
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
	gormDB := database.Connect(user, password, host, databaseName, "mysql")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(gormDB.DB())

	route.Initialize(e, gormDB)
	e.Logger.Fatal(e.Start(":8080"))
}
