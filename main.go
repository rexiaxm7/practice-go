package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rexiaxm7/practice-go/route"
)

func main() {
	e := echo.New()
	route.Initialize(e)

	e.Logger.Fatal(e.Start(":8080"))
}
