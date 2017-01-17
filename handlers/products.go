package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/spetrunin/vending_machine/models"
)

func GetProductsList(c echo.Context) error {
	return c.JSON(http.StatusOK, models.DispenserInstance.ProductsList())
}

func BuyProduct(c echo.Context) error {
	var (
		dispenser = models.DispenserInstance
		amount    = dispenser.MoneyStorage.DepositAmount()
	)

	productID, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		amount, err = dispenser.RequestProduct(productID)
	}

	if err != nil {
		return c.JSON(http.StatusOK, struct {
			Amount float32
			Err    string
		}{amount, err.Error()})
	}

	return c.JSON(http.StatusOK, struct{ Amount float32 }{amount})
}
