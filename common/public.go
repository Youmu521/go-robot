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
