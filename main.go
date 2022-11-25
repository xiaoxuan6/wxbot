package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"time"
)

var (
	err         error
	friends     openwechat.Friends // 有缓存
	currentUser *openwechat.Self
)

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop)

	bot.UUIDCallback = ConsoleQrCode
	bot.SyncCheckCallback = nil

	if err = bot.Login(); err != nil {
		logrus.Fatalf("bot Login err %s", err.Error())
	}

	currentUser, err = bot.GetCurrentUser()
	if err != nil {
		logrus.Fatalf("GetCurrentUser err: %s ", err.Error())
	}

	friends, err = currentUser.Friends(true)
	if err != nil {
		logrus.Fatalf("wx self get friends err: %s ", err.Error())
	}

	//for _, v := range friends {
	//	logrus.Infof("username: %s、nickname：%s、remarkname：%s", v.UserName, v.NickName, v.RemarkName)
	//}

	go Ticker()

	bot.Block()

}

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}

func Ticker() {
	for {
		select {
		case t := <-time.After(1 * time.Minute):
			nowTime := t.Format("15:04") // 2006-01-02 15:04:05
			logrus.Infof("nowTiem %s", nowTime)
			if nowTime == "9:00" {
				go sendMessageWithUser()
			}
		}
	}
}

func sendMessageWithUser() {
	//resFriends := friends.SearchByNickName(1, "君发大头商店")
	resFriends := friends.SearchByRemarkName(1, "小号")
	friend := resFriends.First()
	if friend == nil {
		logrus.Fatalf("search friend fail")
	}

	var message *openwechat.SentMessage
	message, err = friend.SendText("签到")
	if err != nil {
		logrus.Fatalf("friend send text err: %s", err.Error())
	}

	logrus.Infof("friend send text mesif: %s", message.MsgId)
}
