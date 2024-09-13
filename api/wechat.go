package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/pkg/wechat"
)

type SendMessageRequest struct {
	ToUsers []string `json:"to_users"`
	Content string   `json:"content"`
}

func SendMessage(c *gin.Context) {
	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "invalid request body",
		})
		return
	}

	f, err := wechat.Self.Friends()

	for _, friend := range f {
		log.Println(friend.NickName)
	}

	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "send message failed, err: " + err.Error(),
		})
		return
	}

	for _, toUser := range req.ToUsers {
		for _, friend := range f {
			if friend.NickName == toUser {
				log.Println("send message to " + toUser)
				friend.SendText(req.Content)
			}
		}
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
