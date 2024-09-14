package cron

import (
	"log"
	"strings"

	"github.com/robfig/cron/v3"
	"github.com/xbmlz/baojia/model"
	"github.com/xbmlz/baojia/pkg/wechat"
)

func Run() {
	c := cron.New()
	c.AddFunc("0 10,14,18,22 * * *", func() {
		log.Println("wx notify cron job start")
		sendWxMsg(`https://bj.xbmlz.cc/
最新报价网站 华为参考价格 苹果参考价格比较准 
本消息由报价机器人发送`)
		log.Println("wx notify cron job end")
	})

	c.AddFunc("0 0 11,15,21 * *?", func() {
		log.Println("wx notify cron job start")
		sendWxMsg(`今天明天到货提前联系群主啦
中秋不放假
本消息由报价机器人发送机器人发送`)
		log.Println("wx notify cron job end")
	})
	c.Start()
}

func sendWxMsg(msg string) {

	g, err := wechat.Self.Groups()
	if err != nil {
		log.Println(err)
		return
	}

	configGroups, err := model.GetConfigByKey("cron:wx-notify-groups")
	if err != nil {
		log.Println(err)
		return
	}

	if configGroups.ConfigVal == "" {
		log.Println("config is empty")
		return
	}

	toGroups := strings.Split(configGroups.ConfigVal, ",")

	for _, toGroup := range toGroups {
		for _, group := range g {
			if group.NickName == toGroup {
				log.Println("send message to " + toGroup)
				group.SendText(msg)
			}
		}
	}
}
