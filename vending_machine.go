package main

import (
	"github.com/labstack/echo"
	"github.com/spetrunin/vending_machine/handlers"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.Root)

	api := e.Group("/api")

	// deposit and cashback
	api.POST("/deposit", handlers.MakeDeposit)
	api.GET("/deposit", handlers.GetAvailableDeposit)
	api.DELETE("/deposit", handlers.RequestCashback)

	// products
	api.GET("/products", handlers.GetProductsList)
	api.GET("/products/:id/buy", handlers.GetProductsList)

	e.Logger.Fatal(e.Start(":1323"))
}
