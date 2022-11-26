package ticker

import (
	"github.com/sirupsen/logrus"
	"github.com/xiaoxuan6/wxbot/global"
	"time"
)

func sign() {
	for {
		select {
		case t := <-time.After(1 * time.Minute):
			nowTime := t.Format("15:04") // 2006-01-02 15:04:05
			//logrus.Infof("nowTiem %s", nowTime)
			if nowTime == "16:00" {
				go sendMessageWithUser()
			}

			//go sendMessageWithGroup()
		}
	}
}

func sendMessageWithUser() {
	//resFriends := friends.SearchByNickName(1, "君发大头商店")
	resFriends := global.Friends.SearchByRemarkName(1, "A009小鸭")
	friend := resFriends.First()
	if friend == nil {
		logrus.Fatalf("search friend fail")
	}

	message, err := friend.SendText("签到")
	if err != nil {
		logrus.Fatalf("friend send text err: %s", err.Error())
	}

	logrus.Infof("friend send text mesif: %s", message.MsgId)
}

func sendMessageWithGroup() {
	group := global.Groups.SearchByNickName(1, "开发群")
	//group := global.Groups.SearchByNickName(1, "开发群").SendText("")

	first := group.First()
	if first == nil {
		logrus.Fatalf("search group fail")
	}

	first.SendText("测试")
}
