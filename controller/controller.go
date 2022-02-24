package controller

import (
	"fmt"
	"mini-wallet-exercise/helper"
	"mini-wallet-exercise/memModel"
	"time"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func InitWallet(c echo.Context) error {
	form := memModel.POSTCustomer{}
	c.Bind(&form)
	if form.CustomerXID == "" {
		fmt.Println("please input customer xid")
		return c.JSON(400, M{
			"status":  "error",
			"message": "please input customer xid",
		})
	}
	token, err := helper.CreateToken(form.CustomerXID)
	if err != nil {
		return c.JSON(400, M{
			"status":  "error",
			"message": "an error has been occured",
		})
	}
	return c.JSON(200, M{
		"status": "success",
		"data": M{
			"token": token,
		},
	})
}

func EnableWallet(c echo.Context) error {
	xid := helper.ExtractTokenXID(c)
	if xid == "" {
		return c.JSON(400, M{
			"status":  "error",
			"message": "an error has been occured",
		})
	}
	for _, v := range memModel.Wallets {
		if v.OwnedBy == xid {
			if v.Status == "enabled" {
				return c.JSON(400, M{
					"status":  "error",
					"message": "wallet already enabled",
				})
			}
			v.Status = "enabled"
			v.EnableAt = time.Now().String()
			return c.JSON(200, M{
				"status": "success",
				"data": M{
					"wallet": v,
				},
			})
		}
	}
	return c.JSON(400, M{
		"status":  "error",
		"message": "user doesn't exist",
	})
}

func GetBalance(c echo.Context) error {
	time.Sleep(time.Second * 5) // delay for updating balance 5 seconds
	xid := helper.ExtractTokenXID(c)
	for _, v := range memModel.Wallets {
		if v.OwnedBy == xid {
			if v.Status == "disabled" {
				return c.JSON(400, M{
					"status":  "error",
					"message": "wallet is disabled",
				})
			}
			return c.JSON(200, M{
				"status": "success",
				"data": M{
					"wallet": v,
				},
			})
		}
	}
	return c.JSON(400, M{
		"status":  "error",
		"message": "user doesn't exist",
	})
}

func AddMoney(c echo.Context) error {
	var input = memModel.POSTAddVirtualMoney{}
	c.Bind(&input)
	if input.Amount == 0 {
		return c.JSON(400, M{
			"status":  "error",
			"message": "please input amount",
		})
	}
	if input.ReferenceID == "" {
		return c.JSON(400, M{
			"status":  "error",
			"message": "please input reference id",
		})
	}
	// cek reference id
	exist := memModel.ReferencesIDS[input.ReferenceID]
	if exist {
		return c.JSON(400, M{
			"status":  "error",
			"message": "reference id already exist",
		})
	}

	xid := helper.ExtractTokenXID(c)
	for _, v := range memModel.Wallets {
		if v.OwnedBy == xid {
			if v.Status == "disabled" {
				return c.JSON(400, M{
					"status":  "error",
					"message": "wallet is disabled",
				})
			}
			v.Balance += input.Amount
			// add reference id to map
			memModel.ReferencesIDS[input.ReferenceID] = true
		}
	}
	id, err := helper.GenerateIDs()
	if err != nil {
		return c.JSON(500, M{
			"status":  "error",
			"message": "an error has been occured",
		})
	}
	response := memModel.AddMoney{
		ID:          id,
		DepositedBy: xid,
		Status:      "success",
		DepositedAt: time.Now().String(),
		Amount:      input.Amount,
		ReferenceID: input.ReferenceID,
	}
	return c.JSON(200, M{
		"status": "success",
		"data": M{
			"deposit": response,
		},
	})
}

func Withdrawal(c echo.Context) error {
	input := memModel.POSTWithdrawVirtualMoney{}
	c.Bind(&input)
	if input.Amount == 0 {
		return c.JSON(400, M{
			"status":  "error",
			"message": "please input amount",
		})
	}
	if input.ReferenceID == "" {
		return c.JSON(400, M{
			"status":  "error",
			"message": "please input reference id",
		})
	}
	// check if balance is enough
	xid := helper.ExtractTokenXID(c)
	for _, v := range memModel.Wallets {
		if v.OwnedBy == xid {
			if v.Status == "disabled" {
				return c.JSON(400, M{
					"status":  "error",
					"message": "wallet is disabled",
				})
			}
			if v.Balance < input.Amount {
				return c.JSON(400, M{
					"status":  "error",
					"message": "balance is not enough for withdrawal",
				})
			}
			// reduce money
			v.Balance = v.Balance - input.Amount
		}
	}
	id, err := helper.GenerateIDs()
	if err != nil {
		return c.JSON(500, M{
			"status":  "error",
			"message": "an error has been occured",
		})
	}
	response := memModel.Withdrawal{
		ID:          id,
		WithdrawnBy: xid,
		Status:      "success",
		WithdrawnAt: time.Now().String(),
		Amount:      input.Amount,
		ReferenceID: input.ReferenceID,
	}
	return c.JSON(200, M{
		"status": "success",
		"data": M{
			"withdrawals": response,
		},
	})
}

func DisableWallet(c echo.Context) error {
	input := memModel.PATCHdisable{}
	c.Bind(&input)
	if input.IsDisabled != "true" {
		return c.JSON(400, M{
			"status":  "error",
			"message": "please insert a valid input",
		})
	}
	xid := helper.ExtractTokenXID(c)
	temp := memModel.MyWallet{}
	for _, v := range memModel.Wallets {
		if v.Status == "disabled" {
			return c.JSON(400, M{
				"status":  "error",
				"message": "wallet already disabled",
			})
		}
		if v.OwnedBy == xid {
			v.Status = "disabled"
		}
		temp = *v
	}
	return c.JSON(200, M{
		"status": "success",
		"data": M{
			"wallet": temp,
		},
	})
}
