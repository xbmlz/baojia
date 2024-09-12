package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/model"
)

func GetArticles(c *gin.Context) {
	articles, err := model.GetArticles()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "获取文章失败",
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取文章成功",
		"data": articles,
	})
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	article, err := model.GetArticle(idInt)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "获取文章失败",
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取文章成功",
		"data": article,
	})
}

func AddArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "参数错误",
			"data": nil,
		})
		return
	}

	user, _ := getLoginUser(c)

	article.Author = user.Username
	err := model.CreateArticle(article)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "添加文章失败",
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "添加文章成功",
		"data": nil,
	})
}

func UpdateArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "参数错误",
			"data": nil,
		})
		return
	}
	err := model.UpdateArticle(article)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "更新文章失败",
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "更新文章成功",
		"data": nil,
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	err := model.DeleteArticle(idInt)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "删除文章失败",
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "删除文章成功",
		"data": nil,
	})
}
