package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spetrunin/vending_machine/models"
)

func MakeDeposit(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.Message{Text: "Added 100$"})
}

func GetAvailableDeposit(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.Message{Text: "500$"})
}
