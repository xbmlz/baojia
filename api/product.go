package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/model"
	"gorm.io/gorm"
)

type UpadatePriceRequest struct {
	ProductID int     `json:"product_id"`
	InPrice   float64 `json:"in_price"`
	OutPrice  float64 `json:"out_price"`
	Profit    float64 `json:"profit"`
	CreatedAt string  `json:"created_at"`
}

type AddProductRequest struct {
	Type     int      `json:"type"`
	Brand    string   `json:"brand"`
	Series   string   `json:"series"`
	Model    string   `json:"model"`
	Colors   []string `json:"colors"`
	Versions []string `json:"versions"`
}

func AddProduct(ctx *gin.Context) {
	var req AddProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}

	err := model.GetDB().Transaction(func(tx *gorm.DB) error {
		// 循环添加 1 brand => n colors[i] => n versions[i]
		for _, color := range req.Colors {
			for _, version := range req.Versions {
				if err := tx.Create(&model.Product{
					Type:    req.Type,
					Series:  req.Series,
					Model:   req.Model,
					Brand:   req.Brand,
					Color:   color,
					Version: version,
				}).Error; err != nil {
					// 返回任何错误都会回滚事务
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "服务器内部错误",
		})
		return
	}

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

func GetProductTypes(ctx *gin.Context) {
	types := model.GetProductTypeList()
	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "success",
		"data":  types,
		"count": len(types),
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
		price.InPrice = req.OutPrice - req.InPrice
		price.OutPrice = req.OutPrice
		price.Profit = req.Profit
		price.CreatedAt = req.CreatedAt
		model.UpdatePrice(price)
	} else {
		// 新增价格记录
		price = model.Price{
			ProductID: req.ProductID,
			InPrice:   req.OutPrice - req.InPrice,
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

func GetBrands(ctx *gin.Context) {
	productType := ctx.Query("type")
	pt, _ := strconv.Atoi(productType)

	brands := model.GetBrandList(pt)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "success",
		"data":  brands,
		"count": len(brands),
	})
}
