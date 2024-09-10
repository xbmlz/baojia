package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/model"
)

type ConfirmSaleRequest struct {
	ID    int     `json:"id"`
	Price float64 `json:"price"`
	// 结算方式 1: 支付宝 2 微信 3 银行卡
	Payment int `json:"payment"`
}

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

	user, err := getLoginUser(c)
	if err != nil {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "用户未登录",
			"data": err.Error(),
		})
		return
	}
	req.UserID = user.ID

	err = model.CreateSale(req)
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

func GetSale(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	sale, err := model.GetSale(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取成功",
		"data": sale,
	})
}

func GetSales(c *gin.Context) {

	// session := sessions.Default(ctx)
	// userID := session.Get(middleware.SessionUserKey)
	status := c.Query("status")

	statusInt, _ := strconv.Atoi(status)

	user, err := getLoginUser(c)
	if err != nil {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "用户未登录",
			"data": err.Error(),
		})
		return
	}
	var userId int
	// if is admin, get all sales, else get sales of the user
	if user.IsAdmin {
		userId = 0
	} else {
		userId = user.ID
	}

	sales, err := model.GetSales(userId, statusInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取成功",
		"data": sales,
	})
}

func ConfirmSale(c *gin.Context) {
	var req ConfirmSaleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
			"data": err.Error(),
		})
		return
	}

	user, err := getLoginUser(c)
	if err != nil {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "用户未登录",
			"data": err.Error(),
		})
		return
	}

	err = model.ConfirmSale(req.ID, req.Price, req.Payment, user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "确认失败",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "确认成功",
	})
}
