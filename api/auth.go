package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/middleware"
	"github.com/xbmlz/baojia/model"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest = RegisterRequest

func Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// check username is exist
	_, exist, err := model.GetUserByUsername(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	if exist {
		ctx.JSON(http.StatusConflict, gin.H{
			"code": 5001,
			"msg":  "用户名已存在",
		})
		return
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := model.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"msg":  "注册成功",
	})
}

func Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, exist, err := model.GetUserByUsername(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 4011,
			"msg":  "用户不存在",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 4012,
			"msg":  "密码错误",
		})
		return
	}

	session := sessions.Default(ctx)
	session.Set(middleware.SessionUserKey, user.ID)
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
	})
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete(middleware.SessionUserKey)
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "退出成功",
	})
}

func GetUserInfo(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get(middleware.SessionUserKey)
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 4013,
			"msg":  "用户未登录",
		})
		return
	}
	user, err := model.GetUserByID(userID.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": user,
	})
}

// func GetProductPicker(ctx *gin.Context) {
// 	products := model.GetProductList("")
// 	// 按照品牌分组
// 	// [{text: "品牌1", children: [{text: "产品1"}, {text: "产品2"}]}, {text: "品牌2", children: [{text: "产品3"}, { text: "产品4"}]}]
// 	type Picker struct {
// 		Text     string `json:"text"`
// 		Children []struct {
// 			Text string `json:"text"`
// 		} `json:"children"`
// 	}
// 	brandMap := make(map[string][]model.Product)
// 	for _, product := range products {
// 		brandMap[product.Brand] = append(brandMap[product.Brand], product)
// 	}

// 	pickers := make([]Picker, 0)
// 	for brand, products := range brandMap {
// 		picker := Picker{
// 			Text: brand,
// 			Children: make([]struct {
// 				Text string `json:"text"`
// 			}, 0),
// 		}
// 		for _, product := range products {
// 			picker.Children = append(picker.Children, struct {
// 				Text string `json:"text"`
// 			}{
// 				Text: fmt.Sprintf("%s-%s-%s", product.Model, product.Color, product.Version),
// 			})
// 		}
// 		pickers = append(pickers, picker)
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"msg":  "success",
// 		"data": pickers,
// 	})
// }

// func GetAppProducts(ctx *gin.Context) {
// 	brand := ctx.Query("brand")
// 	products := model.GetProductList(brand)
// 	prices := model.GetCurrentPrice(products.ToIDs())
// 	for i, product := range products {
// 		for _, price := range prices {
// 			if price.ProductID == product.ID {
// 				products[i].Price = ""
// 				products[i].Profit = ""
// 				products[i].RecoveryPrice = strconv.FormatFloat(price.RecoveryPrice, 'f', 2, 64)
// 			}
// 		}
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code":  0,
// 		"msg":   "success",
// 		"data":  products,
// 		"count": len(products),
// 	})
// }
