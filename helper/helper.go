package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateToken(customer_xid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["customer_xid"] = customer_xid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires after 24 hour

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ExtractTokenXID(c echo.Context) string {
	token := c.Get("user").(*jwt.Token)
	if token != nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		xid := claims["customer_xid"].(string)
		if xid != "" {
			// return xid[1 : len(xid)-1]
			return xid
		}
	}
	return "" // invalid user
}

func GenerateIDs() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

var IsLoggin = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
