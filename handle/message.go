package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"robot/common"
	"robot/config"
	"robot/message"
)

// Type 判断消息类型，并处理消息
func Type(data []byte) {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()
	// json 转 结构体
	eventType := config.EventType{}
	err := json.Unmarshal(data, &eventType)
	if err != nil {
		return
	}
	switch eventType.PostType {
	case common.EventTypeMeta:
		// 元事件
		fmt.Println("连接成功")
	case common.EventTypeMessage:
		// 消息
		msgEvent := config.MsgEvent{}
		err := json.Unmarshal(data, &msgEvent)
		if err != nil {
			return
		}
		// 解析消息
		message.Analysis(msgEvent, data)
	case common.EventTypeRequest:
		// 请求
	case common.EventTypeNotice:
		//通知
	default:
		// 类型不正确
		fmt.Println("消息格式不正确")
	}
}
