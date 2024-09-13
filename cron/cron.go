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
	c.AddFunc("0 10,13,16,19 * * *", func() {
		log.Println("wx notify cron job start")

		configGroups, err := model.GetConfigByKey("cron:wx-notify-groups")
		if err != nil {
			log.Println(err)
			return
		}

		configContent, err := model.GetConfigByKey("cron:wx-notify-content")
		if err != nil {
			log.Println(err)
			return
		}

		g, err := wechat.Self.Groups()
		if err != nil {
			log.Println(err)
			return
		}

		if configGroups.ConfigVal == "" || configContent.ConfigVal == "" {
			log.Println("config is empty")
			return
		}

		toGroups := strings.Split(configGroups.ConfigVal, ",")

		for _, toGroup := range toGroups {
			for _, group := range g {
				if group.NickName == toGroup {
					log.Println("send message to " + toGroup)
					group.SendText(configContent.ConfigVal)
				}
			}
		}

		log.Println("wx notify cron job end")

	})
	c.Start()
}
