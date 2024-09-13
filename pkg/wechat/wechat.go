package wechat

import (
	"fmt"
	"log"

	"github.com/eatmoreapple/openwechat"
)

var (
	Bot  *openwechat.Bot
	Self *openwechat.Self
)

func InitWeChatBot() {
	Bot := openwechat.DefaultBot(openwechat.Desktop)
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()

	Bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}

	err := Bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	if err != nil {
		fmt.Println(err)
	}
	// 获取当前登录用户
	Self, err = Bot.GetCurrentUser()
	if err != nil {
		log.Fatal(err)
	}

	err = Bot.Block()
	if err != nil {
		log.Fatal(err)
	}
}
