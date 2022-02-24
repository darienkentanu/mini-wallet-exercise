package main

import (
	"mini-wallet-exercise/controller"
	"mini-wallet-exercise/helper"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func main() {
	e := echo.New()

	e.POST("api/v1/init", controller.InitWallet)
	e.POST("api/v1/wallet", controller.EnableWallet, helper.IsLoggin)
	e.GET("api/v1/wallet", controller.GetBalance, helper.IsLoggin)
	e.POST("api/v1/wallet/deposits", controller.AddMoney, helper.IsLoggin)
	e.POST("api/v1/wallet/withdrawals", controller.Withdrawal, helper.IsLoggin)
	e.PATCH("api/v1/wallet", controller.DisableWallet, helper.IsLoggin)

	e.Logger.Fatal(e.Start(":8080"))
}
