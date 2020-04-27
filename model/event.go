package model

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Type           EventType  `json:"type,omitempty"`
	MsgId          uint64     `json:"msg_id,omitempty"`
	MsgIdStr       string     `json:"msg_id_str,omitempty"`
	MsgType        MsgType    `json:"msg_type,omitempty"`
	FromUser       string     `json:"from_user,omitempty"`
	FromUserNick   string     `json:"from_user_nick,omitempty"`
	FromUserAvatar string     `json:"from_user_head,omitempty"`
	Sender         string     `json:"sender,omitempty"`
	ToUser         string     `json:"to_user,omitempty"`
	Content        string     `json:"content,omitempty"`
	Url            string     `json:"url,omitempty"`
	Image          string     `json:"image,omitempty"`
	Title          string     `json:"title,omitempty"`
	AppOriId       string     `json:"app_ori_id,omitempty"`
	AppId          string     `json:"appid,omitempty"`
	Path           string     `json:"path,omitempty"`
	Style          EventStyle `json:"style,omitempty"`
	Contacts       []Contact  `json:"contacts,omitempty"`
	ContactCount   uint       `json:"contact_count,omitempty"`
	CreateTime     int64      `json:"create_time,omitempty"`
	Group          *Group     `json:"group,omitempty"`
	AtList         string     `json:"at_list,omitempty"`
}

type EventRequest struct {
	Payload   string `json:"payload,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Sign      string `json:"sign,omitempty"`
	Key       string `json:"key,omitempty"`
	Code      int    `json:"code,omitempty"`
	ErrMsg    string `json:"msg,omitempty"`
}

func (this *EventRequest) IsError() bool {
	return this.Code != 0
}

func (this EventRequest) Error() string {
	return fmt.Sprintf("CODE: %d, MSG: %s", this.Code, this.ErrMsg)
}

func (this *EventRequest) Event() (*Event, error) {
	if this.Payload == "" {
		return nil, nil
	}
	var event Event
	err := json.Unmarshal([]byte(this.Payload), &event)
	return &event, err
}
