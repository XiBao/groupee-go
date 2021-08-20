package model

// EventType event type
type EventType = string

const (
	// UNKNOWN_EVENT 未知事件
	UNKNOWN_EVENT EventType = ""
	// MSG_EVENT 消息事件
	MSG_EVENT EventType = "msg"
	// JOIN_GROUP_EVENT 进群事件
	JOIN_GROUP_EVENT EventType = "join_group"
	// LEAVE_GROUP_EVENT 退群事件
	LEAVE_GROUP_EVENT EventType = "leave_group"
	// GROUP_CONTACTS_EVENT 群联系人变更事件
	GROUP_CONTACTS_EVENT EventType = "group_contacts"
	// KICKOUT_GROUP_EVENT 踢出群事件
	KICKOUT_GROUP_EVENT EventType = "kickout_group"
)

// MsgType message type
type MsgType = uint

const (
	// TXT_MSG 文本消息
	TXT_MSG MsgType = 1
	// IMG_MSG 图片消息
	IMG_MSG MsgType = 3
	// AUDIO_MSG 语音消息
	AUDIO_MSG MsgType = 34
	// IDCARD_MSG 名片消息
	IDCARD_MSG MsgType = 42
	// VIDEO_MSG 视频消息
	VIDEO_MSG MsgType = 43
	// ARTICLE_MSG 公号或链接消息
	ARTICLE_MSG MsgType = 5
	// MINIAPP_MSG 小程序消息
	MINIAPP_MSG MsgType = 33
	// COUPON_MSG 优惠券消息
	COUPON_MSG MsgType = 16
	// TRANSFER_MSG 转账消息
	TRANSFER_MSG MsgType = 2000
	// REDPACKET_MSG 红包消息
	REDPACKET_MSG MsgType = 2001
)

// EventStyle 小程序样式
type EventStyle = uint

const (
	// NORMAL_STYLE 普通样式
	NORMAL_STYLE EventStyle = 0
	// SIMPLE_STYLE 精简样式
	SIMPLE_STYLE EventStyle = 1
)
