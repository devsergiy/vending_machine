package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spetrunin/vending_machine/helpers"
	"github.com/spetrunin/vending_machine/models"
)

type depositRequest struct {
	Deposit float32
}

func MakeDeposit(c echo.Context) error {
	money := models.DispenserInstance.MoneyStorage

	deposit := new(depositRequest)
	if err := c.Bind(deposit); err != nil {
		return err
	}

	money.AddMoney(deposit.Deposit)
	return c.JSON(http.StatusOK, helpers.AmountJson(money.DepositAmount()))
}

func GetAvailableDeposit(c echo.Context) error {
	money := models.DispenserInstance.MoneyStorage

	return c.JSON(http.StatusOK, helpers.AmountJson(money.DepositAmount()))
}

func RequestCashback(c echo.Context) error {
	money := models.DispenserInstance.MoneyStorage

	return c.JSON(http.StatusOK, helpers.CashbackJson(money.MakeCashBack()))
}
