package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/k6_load_test/api/jester/internal/gen"
)

func (c *Controllers) GetProducts(ctx *gin.Context, request gen.GetProductsRequestObject) (gen.GetProductsResponseObject, error) {

	metadata := gen.ProductNextCursor{NextCursor: "MjAwMDI5OTA="}
	products := []gen.GetProducts{
		{
			Id:              20001001,
			Name:            "プレミアムコーヒー",
			CategoryId:      10,
			CategoryName:    "飲料",
			Description:     "香り高いアラビカ種のコーヒーです。",
			Price:           500.0,
			DiscountFlag:    true,
			DiscountName:    "2024年クリスマスキャンペーン",
			DiscountRate:    20,
			DiscountedPrice: 400.0,
			StockQuantity:   50,
			VipOnly:         false,
			ImageUrl:        "https://example.com/images/20001001/product.jpg",
		},
		{
			Id:              20001002,
			Name:            "エスプレッソマシン",
			CategoryId:      20,
			CategoryName:    "キッチン用品",
			Description:     "自宅で本格的なエスプレッソを楽しめるマシン。",
			Price:           20000.0,
			DiscountFlag:    false,
			DiscountName:    "",
			DiscountRate:    0,
			DiscountedPrice: 20000.0,
			StockQuantity:   10,
			VipOnly:         true,
			ImageUrl:        "https://example.com/images/20001002/product.jpg",
		},
		{
			Id:              20001003,
			Name:            "ハンドメイドマグカップ",
			CategoryId:      30,
			CategoryName:    "食器",
			Description:     "温もりのあるデザインのハンドメイドマグカップ。",
			Price:           1500.0,
			DiscountFlag:    true,
			DiscountName:    "新生活応援セール",
			DiscountRate:    10,
			DiscountedPrice: 1350.0,
			StockQuantity:   100,
			VipOnly:         false,
			ImageUrl:        "https://example.com/images/20001002/product.jpg",
		},
	}

	return gen.GetProducts200JSONResponse{
		Metadata: metadata,
		Products: products,
	}, nil
}
