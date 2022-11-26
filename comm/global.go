package comm

import "github.com/eatmoreapple/openwechat"

var (
	CurrentUser *openwechat.Self
	Friends     openwechat.Friends // 有缓存
	Groups      openwechat.Groups
)
