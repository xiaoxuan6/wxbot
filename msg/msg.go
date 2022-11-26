package msg

import (
	"github.com/eatmoreapple/openwechat"
	"os"
	"strings"
)

func MessageHandler(msg *openwechat.Message) {
	// 仅处理 text 类型的消息
	if msg.IsText() {
		if msg.Content == "test" {
			msg.ReplyText("测试")
			return
		}

		if strings.Contains(msg.Content, "打赏") {
			img, _ := os.Open("img.png")
			defer img.Close()

			msg.ReplyImage(img)
			return
		}
	}
}
