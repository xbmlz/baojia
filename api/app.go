package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/middleware"
	"github.com/xbmlz/baojia/model"
	"golang.org/x/crypto/bcrypt"
)

type AppRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AppLoginRequest = AppRegisterRequest

func AppIndexView(ctx *gin.Context) {
	currentUser := ctx.MustGet(middleware.CurrentUserKey)
	if currentUser == nil {
		ctx.Redirect(http.StatusFound, "/login.html")
		return
	}
	user, _ := model.GetUserByID(currentUser.(int))
	ctx.HTML(http.StatusOK, "app_index.html", gin.H{
		"title": "报价",
		"today": time.Now().Format("2006-01-02"),
		"user":  user,
	})
}

func AppLoginView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "app_login.html", gin.H{
		"title": "登录",
	})
}

func AppRegister(ctx *gin.Context) {
	var req AppRegisterRequest
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

func AppLogin(ctx *gin.Context) {
	var req AppLoginRequest
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
	}

	session := sessions.Default(ctx)
	session.Set(middleware.SessionUserKey, user.ID)
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
	})
}

func AppLogout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete(middleware.SessionUserKey)
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "退出成功",
	})
}

func GetAppProducts(ctx *gin.Context) {
	brand := ctx.Query("brand")
	products := model.GetProductList(brand)
	prices := model.GetCurrentPrice(products.ToIDs())
	for i, product := range products {
		for _, price := range prices {
			if price.ProductID == product.ID {
				products[i].Price = ""
				products[i].Profit = ""
				products[i].RecoveryPrice = strconv.FormatFloat(price.RecoveryPrice, 'f', 2, 64)
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
