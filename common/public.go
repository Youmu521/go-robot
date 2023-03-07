package common

import (
	"robot/message/keywords"
	"robot/request"
)

var KeyMap = map[string]func(msg string){
	"点歌": func(msg string) {
		keywords.ChooseMusic(msg)
	},
	"唱歌": func(msg string) {
		keywords.SingMusic(msg)
	},
	"#": func(msg string) {
		request.M().At(1312563525).Send("点歌成功")
	},
}

// LastMessages 记录每个群最近四条消息
var LastMessages = make(map[int64][]string)
