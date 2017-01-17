package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spetrunin/vending_machine/models"
)

func Root(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.Message{Text: "ok"})
}
