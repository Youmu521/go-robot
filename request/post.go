package request

import (
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"strconv"
)

// sendMsg post 请求
func sendMsg(msg string) {
	// 测试模式
	//client := req.C().DevMode()
	client := req.C()
	// 发送请求
	resp, err := client.R().
		SetHeader("accept", "application/json").
		SetFormData(map[string]string{
			"group_id": "587498574",
			"message":  msg,
		}).
		Post("http://0.0.0.0:8283/send_group_msg")

	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccessState() {
		fmt.Println("bad response status:", resp.Status)
		return
	}
	return
}

// 消息
type message struct {
	Message []string
}

func M() *message {
	m := &message{}
	return m
}

// At  添加 @at 操作
func (m *message) At(qq int) *message {
	var msg = "[CQ:at,qq=" + strconv.Itoa(qq) + "]"
	m.Message = append(m.Message, msg)
	return m
}

// Send 发送消息
func (m *message) Send(s string) {
	var msg = s
	m.Message = append(m.Message, msg)
	var message string
	for _, v := range m.Message {
		message += v
	}
	sendMsg(message)
}
