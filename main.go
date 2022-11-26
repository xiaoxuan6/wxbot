package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"github.com/xiaoxuan6/wxbot/global"
	msg2 "github.com/xiaoxuan6/wxbot/msg"
	"github.com/xiaoxuan6/wxbot/ticker"
)

var err error

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop)

	bot.UUIDCallback = ConsoleQrCode
	bot.SyncCheckCallback = nil

	// 在用户登录后需要实时接受微信发送过来的消息。
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			_, err = msg.ReplyText("pong")
			if err != nil {
				logrus.Fatalf("ping msg replyText err %s", err.Error())
			}
		}

		msg2.MessageHandler(msg)
	}

	if err = bot.Login(); err != nil {
		logrus.Fatalf("bot Login err %s", err.Error())
	}

	global.CurrentUser, err = bot.GetCurrentUser()
	if err != nil {
		logrus.Fatalf("GetCurrentUser err: %s ", err.Error())
	}

	global.Friends, err = global.CurrentUser.Friends(true)
	if err != nil {
		logrus.Fatalf("wx self get friends err: %s ", err.Error())
	}

	global.Groups, err = global.CurrentUser.Groups(true)
	if err != nil {
		logrus.Fatalf("wx self get groups err: %s", err.Error())
	}

	//for _, v := range friends {
	//	logrus.Infof("username: %s、nickname：%s、remarkname：%s", v.UserName, v.NickName, v.RemarkName)
	//}

	ticker.Ticker()

	bot.Block()

}

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}
