package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
