package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/model"
)

type UpadatePriceRequest struct {
	ProductID int     `json:"product_id"`
	InPrice   float64 `json:"in_price"`
	OutPrice  float64 `json:"out_price"`
	Profit    float64 `json:"profit"`
	CreatedAt string  `json:"created_at"`
}

func AddProduct(ctx *gin.Context) {
	var req model.Product
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}

	model.AddProduct(req)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func GetProducts(ctx *gin.Context) {
	product_type := ctx.Query("type")
	brand := ctx.Query("brand")

	pt, _ := strconv.Atoi(product_type)

	products := model.GetProductList(pt, brand)
	// for i, product := range products {
	// 	for _, price := range prices {
	// 		if price.ProductID == product.ID {
	// 			products[i].Price = formatPrice(price.Price)
	// 			products[i].Profit = formatPrice(price.Profit)
	// 			products[i].RecoveryPrice = formatPrice(price.RecoveryPrice)
	// 		}
	// 	}
	// 	price := model.GetLastPrice(product.ID)
	// 	if price.ID > 0 {
	// 		products[i].Price = formatPrice(price.Price)
	// 		products[i].Profit = formatPrice(price.Profit)
	// 		products[i].RecoveryPrice = formatPrice(price.RecoveryPrice)
	// 	}
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "success",
		"data":  products,
		"count": len(products),
	})
}

func UpdatePrice(ctx *gin.Context) {
	var req UpadatePriceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}

	// 查询是否有对应日期的价格记录
	price := model.GetPriceByDate(req.ProductID, req.CreatedAt)
	if price.ID > 0 {
		// 更新价格记录
		price.InPrice = req.InPrice
		price.OutPrice = req.OutPrice
		price.Profit = req.Profit
		price.CreatedAt = req.CreatedAt
		model.UpdatePrice(price)
	} else {
		// 新增价格记录
		price = model.Price{
			ProductID: req.ProductID,
			InPrice:   req.InPrice,
			OutPrice:  req.OutPrice,
			Profit:    req.Profit,
			CreatedAt: req.CreatedAt,
		}
		model.AddPrice(price)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
