package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/k6_load_test/api/jester/internal/gen"
)

func (c *Controllers) CreateCreditCard(ctx *gin.Context, request gen.CreateCreditCardRequestObject) (gen.CreateCreditCardResponseObject, error) {

	return gen.CreateCreditCard204Response{}, nil
}
