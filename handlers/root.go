package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Root(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Status string }{"ok"})
}
