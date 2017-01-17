package main

import (
	"github.com/labstack/echo"
	"github.com/spetrunin/vending_machine/handlers"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.Root)

	api := e.Group("/api")

	api.POST("/deposit", handlers.MakeDeposit)
	api.GET("/deposit", handlers.GetAvailableDeposit)

	api.GET("/products", handlers.GetProductsList)

	e.Logger.Fatal(e.Start(":1323"))
}
