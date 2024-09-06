package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/model"
)

type SavePriceRequest struct {
	ProductID int     `json:"product_id"`
	Price     float64 `json:"price"`
	Profit    float64 `json:"profit"`
}

func AdminView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin_index.html", gin.H{
		"title": "管理后台",
		"today": time.Now().Format("2006-01-02"),
	})
}

func formatPrice(value float64) string {
	if value == 0 {
		return ""
	}
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func GetProducts(ctx *gin.Context) {
	brand := ctx.Query("brand")
	products := model.GetProductList(brand)
	prices := model.GetCurrentPrice(products.ToIDs())
	for i, product := range products {
		for _, price := range prices {
			if price.ProductID == product.ID {
				products[i].Price = formatPrice(price.Price)
				products[i].Profit = formatPrice(price.Profit)
				products[i].RecoveryPrice = formatPrice(price.RecoveryPrice)
			}
		}

	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "success",
		"data":  products,
		"count": len(products),
	})
}

func SavePrice(ctx *gin.Context) {
	var req SavePriceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
	}
	// 更新或新增价格
	model.SavePrice(req.ProductID, req.Price, req.Profit)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
