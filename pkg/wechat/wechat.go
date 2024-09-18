package wechat

import (
	"fmt"
	"log"
	"strings"

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

	Bot.MessageHandler = messageHandler

	err := Bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
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

func messageHandler(msg *openwechat.Message) {
	// 如果消息是以 `bot` 开头，则进行相应的处理
	if msg.IsText() {
		// 处理消息
		fmt.Println(msg.Content)

		if strings.HasPrefix(msg.Content, "qun") {
			sendMsg := strings.TrimPrefix(msg.Content, "qun")
			user, err := msg.Bot().GetCurrentUser()
			if err != nil {
				log.Printf("GetCurrentUser error: %v", err)
				return
			}

			groups, err := user.Groups()

			if err != nil {
				log.Printf("Groups error: %v", err)
				return
			}

			groupNams := make([]string, 0)
			for _, group := range groups {
				groupNams = append(groupNams, group.NickName)
			}
			groups.SendText(sendMsg)

			msg.ReplyText(fmt.Sprintf("已发送至群组：%s", strings.Join(groupNams, ", ")))
		}

	}
}
