package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/middleware"
	"github.com/xbmlz/baojia/model"
)

func CreateSale(c *gin.Context) {
	var req model.Sale

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
			"data": err.Error(),
		})
		return
	}

	err := model.CreateSale(req)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "创建失败",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建成功",
	})
}

func GetSales(ctx *gin.Context) {
	// session := sessions.Default(ctx)
	// userID := session.Get(middleware.SessionUserKey)
	userID := ctx.GetInt(middleware.CurrentUserKey)
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 4013,
			"msg":  "用户未登录",
		})
		return
	}
	user, err := model.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	// if is admin, get all sales, else get sales of the user
	if user.IsAdmin {
		sales, err := model.GetSales()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "获取成功",
			"data": sales,
		})
		return
	}
	sales, err := model.GetSalesByUser(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取成功",
		"data": sales,
	})
}
