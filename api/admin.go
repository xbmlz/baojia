package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/model"
)

func AdminView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.html", gin.H{
		"title": "管理后台",
		"today": time.Now().Format("2006-01-02"),
	})
}

func GetProducts(ctx *gin.Context) {
	brand := ctx.Query("brand")

	products := model.GetProductList(brand)
	prices := model.GetPriceList(products.ToIDs())
	for i, p := range products {
		products[i].Price = prices.GetTodayPrice(p.ID)
		products[i].LastPrice = prices.GetLastDayPrice(p.ID)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "success",
		"data":  products,
		"count": len(products),
	})
}

func SavePrice(ctx *gin.Context) {
	var price model.Price
	if err := ctx.ShouldBindJSON(&price); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
	}
	// 更新或新增价格
	model.SavePrice(price)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
