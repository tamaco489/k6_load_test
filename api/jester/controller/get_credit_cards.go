package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/k6_load_test/api/jester/internal/gen"
)

func (c *Controllers) GetCreditCards(ctx *gin.Context, request gen.GetCreditCardsRequestObject) (gen.GetCreditCardsResponseObject, error) {

	creditCards := []gen.CreditCardList{
		{
			IsDefault:        true,
			MaskedCardNumber: "******123",
		},
		{
			IsDefault:        false,
			MaskedCardNumber: "******567",
		},
		{
			IsDefault:        false,
			MaskedCardNumber: "******890",
		},
	}

	return gen.GetCreditCards200JSONResponse(creditCards), nil
}
