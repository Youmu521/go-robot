package common

// 事件类型
const (
	EventTypeMessage = "message"    // 消息事件
	EventTypeRequest = "request"    // 请求事件
	EventTypeNotice  = "notice"     // 通知事件
	EventTypeMeta    = "meta_event" // 元事件
)

// 消息事件类型
const (
	MessageTypePrivate = "private" // 私聊消息
	MessageTypeGroup   = "group"   // 群聊消息
)

// 通知事件类型

// 请求事件类型
const (
	RequestTypePrivate = "private" // 好友添加请求
	RequestTypeGroup   = "group"   // 群邀请 / 群申请
)
