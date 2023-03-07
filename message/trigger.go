package message

import (
	"encoding/json"
	"robot/common"
	"robot/config"
	"robot/message/reread"
	"strings"
)

// Analysis 解析字符串
func Analysis(event config.MsgEvent, data []byte) {
	switch event.MessageType {
	case common.MessageTypePrivate:
		privateMsg := config.PrivateMsg{}
		err := json.Unmarshal(data, &privateMsg)
		if err != nil {
			return
		}
		AnalysisMessage(privateMsg.RawMessage)
	case common.MessageTypeGroup:
		groupMsg := config.GroupMsg{}
		err := json.Unmarshal(data, &groupMsg)
		if err != nil {
			return
		}
		// 处理群消息，复读
		reread.Reread(groupMsg.GroupId, groupMsg.RawMessage)
		// 触发
		AnalysisMessage(groupMsg.RawMessage)
	}

}

func AnalysisMessage(msg string) {
	for k, v := range common.KeyMap {
		if strings.HasPrefix(msg, k) {
			// 删除字符串前面的 k 部分
			msg = msg[len(k):]
			v(msg)
		}
	}
}
