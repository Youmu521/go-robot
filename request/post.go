package request

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"robot/config"
	"robot/redis"
	"strconv"
)

// sendMsg post 请求
func sendMsg(msg string) {
	// 测试模式
	//client := req.C().DevMode()
	client := req.C()
	// 获取 data 数据 并拼装成 结构体
	data := redis.Get("data")
	messageSource := config.MessageSource{}
	err := json.Unmarshal(data, &messageSource)
	if err != nil {
		return
	}
	formData := make(map[string]string)
	// 若 群号不为 0，则来源为群
	if messageSource.GroupId != 0 {
		formData = map[string]string{
			"message_type": "group",
			"group_id":     strconv.FormatInt(messageSource.GroupId, 10),
			"message":      msg,
		}
	} else {
		formData = map[string]string{
			"message_type": "private",
			"user_id":      strconv.FormatInt(messageSource.UserId, 10),
			"message":      msg,
		}
	}
	// 发送请求
	resp, err := client.R().
		SetHeader("accept", "application/json").
		SetFormData(formData).
		Post("http://0.0.0.0:8283/send_msg")

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

// SendMusic 发送音乐
func (m *message) SendMusic(id int64) {
	str := "[CQ:music,type=163,id=" + strconv.FormatInt(id, 10) + "]"
	sendMsg(str)
}

// SendRecord 发送语音
func (m *message) SendRecord(url string) {
	str := "[CQ:record,cache=0,proxy=0,file=" + url + "]"
	sendMsg(str)
}
