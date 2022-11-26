package api

import (
	"account-consumer-service/pkg"

	"github.com/labstack/echo"
)

type AccountApi struct {
	service *pkg.AccountServiceProducer
}

func NewAccountApi(service *pkg.AccountServiceProducer) *AccountApi {
	return &AccountApi{
		service: service,
	}
}

func (c *AccountApi) Register(router *echo.Echo) {
	v1 := router.Group("/v1")
	v1.POST("/accounts", c.createAccount)
}
