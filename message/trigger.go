package message

import (
	"encoding/json"
	"robot/common"
	"robot/config"
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
		AnalysisMessage(groupMsg.RawMessage)
	}

}

func AnalysisMessage(msg string) {
	for k, v := range common.KeyMap {
		if strings.HasPrefix(msg, k) {
			v()
		}
	}
}
