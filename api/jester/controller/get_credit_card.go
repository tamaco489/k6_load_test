package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/k6_load_test/api/jester/internal/gen"
)

func (c *Controllers) GetCreditCard(ctx *gin.Context, request gen.GetCreditCardRequestObject) (gen.GetCreditCardResponseObject, error) {

	return gen.GetCreditCard200JSONResponse{}, nil
}
