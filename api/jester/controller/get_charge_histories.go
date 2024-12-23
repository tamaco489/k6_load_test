package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/k6_load_test/api/jester/internal/gen"
)

func (c *Controllers) GetChargeHistories(ctx *gin.Context, request gen.GetChargeHistoriesRequestObject) (gen.GetChargeHistoriesResponseObject, error) {

	return gen.GetChargeHistories200JSONResponse{}, nil
}
