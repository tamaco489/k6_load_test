package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/k6_load_test/api/jester/internal/gen"
)

func (c *Controllers) DeleteCreditCard(ctx *gin.Context, request gen.DeleteCreditCardRequestObject) (gen.DeleteCreditCardResponseObject, error) {

	return gen.DeleteCreditCard204Response{}, nil
}
