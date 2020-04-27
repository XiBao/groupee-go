package model

type EventType = string

const (
	UNKNOWN_EVENT        EventType = ""
	MSG_EVENT            EventType = "msg"
	JOIN_GROUP_EVENT     EventType = "join_group"
	LEAVE_GROUP_EVENT    EventType = "leave_group"
	GROUP_CONTACTS_EVENT EventType = "group_contacts"
	KICKOUT_GROUP_EVENT  EventType = "kickout_group"
)

type MsgType = uint

const (
	TXT_MSG       MsgType = 1
	IMG_MSG       MsgType = 3
	AUDIO_MSG     MsgType = 34
	IDCARD_MSG    MsgType = 42
	VIDEO_MSG     MsgType = 43
	ARTICLE_MSG   MsgType = 5
	MINIAPP_MSG   MsgType = 33
	COUPON_MSG    MsgType = 16
	TRANSER_MSG   MsgType = 2000
	REDPACKET_MSG MsgType = 2001
)

type EventStyle = uint

const (
	NORMAL_STYLE EventStyle = 0
	SIMPLE_STYLE EventStyle = 1
)
