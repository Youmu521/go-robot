package config

// 所有消息结构

// EventType 事件类型
type EventType struct {
	PostType string `json:"post_type"` // 消息类型
	Time     int64  `json:"time"`      // 事件发生的时间戳
	SelfId   int64  `json:"self_id"`   // 收到事件的机器人 QQ 号
}

// 消息事件 Start

// MsgEvent 消息事件
type MsgEvent struct {
	EventType
	MessageType string `json:"message_type"` // 消息类型 private
	SubType     string `json:"sub_type"`     // 消息子类型 friend group group_self
}

// Sender 消息发送者信息
type Sender struct {
	// 私聊
	UserId   int64  `json:"user_id"`  // 发送者 QQ 号
	Nickname string `json:"nickname"` // 昵称
	Sex      string `json:"sex"`      // 性别 male 或 female 或 unknown
	Age      int32  `json:"age"`      // 年龄

	// 临时会话
	GroupId int64 `json:"group_id"` // 临时群消息来源群号

	// 群聊
	Card  string `json:"card"`  // 群名片／备注
	Area  string `json:"area"`  // 地区
	Level string `json:"level"` // 成员等级
	Role  string `json:"role"`  // 角色, owner 或 admin 或 member
	Title string `json:"title"` // 专属头衔
}

// Msg 消息内容
type Msg struct {
	Type string                 `json:"type"` // 消息类型
	Data map[string]interface{} `json:"data"` // 数据 数据参数名: 数据参数值
}

type MsgType interface {
}

// PrivateMsg 私聊消息数据
type PrivateMsg struct {
	MsgEvent
	MessageId  int32  `json:"message_id"`  // 消息 ID
	UserId     int64  `json:"user_id"`     // 发送者 QQ 号
	Message    []Msg  `json:"message"`     // 消息内容
	RawMessage string `json:"raw_message"` // 原始消息内容
	Font       int32  `json:"font"`        // 字体
	Sender     Sender `json:"sender"`      // 发送人信息
	TempSource int    `json:"temp_source"` // 临时会话来源
}

// GroupMsg 群消息
type GroupMsg struct {
	MsgEvent
	MessageId  int32  `json:"message_id"`  // 消息 ID
	UserId     int64  `json:"user_id"`     // 发送者 QQ 号
	Message    []Msg  `json:"message"`     // 消息内容
	RawMessage string `json:"raw_message"` // 原始消息内容
	Font       int32  `json:"font"`        // 字体
	Sender     Sender `json:"sender"`      // 发送人信息
	TempSource int    `json:"temp_source"` // 临时会话来源
	GroupId    int64  `json:"group_id"`    // 群号
	//Anonymous   Anonymous `json:"anonymous"`    // 匿名信息, 如果不是匿名消息则为 null
}

//// Anonymous 匿名信息
//type Anonymous struct {
//	Id   int64  `json:"id"`   // 匿名用户 ID
//	Name string `json:"name"` // 匿名用户名称
//	Flag string `json:"flag"` // 匿名用户 flag, 在调用禁言 API 时需要传入
//}

// 消息事件 End

// MessageSource 消息来源
type MessageSource struct {
	UserId  int64 `json:"user_id"`  // 发送者 QQ 号
	GroupId int64 `json:"group_id"` // 群号
}
