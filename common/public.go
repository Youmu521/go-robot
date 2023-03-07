package common

import "robot/request"

var KeyMap = map[string]func(){
	"点歌": func() {
		request.M().Send("点歌成功")
	},
	"#": func() {
		request.M().At(1312563525).Send("点歌成功")
	},
}

// LastMessages 记录每个群最近四条消息
var LastMessages = make(map[int64][]string)
