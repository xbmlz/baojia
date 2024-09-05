package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func IndexView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Baojia",
		"today": time.Now().Format("2006-01-02"),
	})
}
