package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
)

type Order struct {
	PairSymbol string          "json:'pairSymbol'"
	OrderType  string          "json:'orderType'"
	Quantity   decimal.Decimal "json:'pairSymbol'"
	Price      decimal.Decimal "json:'price'"
}

func getOrder(c echo.Context) error {
	return c.String(http.StatusOK, "GET request")
}

func createOrder(c echo.Context) error {
	order := Order{}

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &order)
	if err != nil {
		return err
	}

	fmt.Println(order)
	return c.String(http.StatusCreated, "Order Created!")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "guest" && password == "admin" {
			return true, nil
		}

		return false, nil
	}))

	orderGroup := e.Group("/v1/order")
	orderGroup.GET("/", getOrder)
	orderGroup.POST("/", createOrder)

	e.Logger.Fatal(e.Start(":8080"))
}
