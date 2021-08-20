package model

import (
	"encoding/json"
	"fmt"
	"html"
)

// Event event struct
type Event struct {
	// Type EventType
	Type EventType `json:"type,omitempty"`
	// MsgId message id
	MsgId uint64 `json:"msg_id,omitempty"`
	// MsgIdStr message id string
	MsgIdStr string `json:"msg_id_str,omitempty"`
	// MsgType message type
	MsgType MsgType `json:"msg_type,omitempty"`
	// FromUser group name if message from a group; or sender
	FromUser string `json:"from_user,omitempty"`
	// FromUserNick group nikc if message from a group; or sender nick
	FromUserNick string `json:"from_user_nick,omitempty"`
	// FromUserAvatar group avarar if message from a group; or sender avatar
	FromUserAvatar string `json:"from_user_head,omitempty"`
	// Sender sender nick
	Sender string `json:"sender,omitempty"`
	// ToUser receiver
	ToUser string `json:"to_user,omitempty"`
	// Content text message
	Content string `json:"content,omitempty"`
	// Url link
	Url string `json:"url,omitempty"`
	// Image image link
	Image string `json:"image,omitempty"`
	// Title title for miniapp or article
	Title string `json:"title,omitempty"`
	// AppName miniapp name
	AppName string `json:"app_name,omitempty"`
	// AppOriId miniapp ori_id
	AppOriId string `json:"app_ori_id,omitempty"`
	// AppId miniapp app id
	AppId string `json:"appid,omitempty"`
	// Path miniapp path
	Path string `json:"path,omitempty"`
	// Style miniapp style
	Style EventStyle `json:"style,omitempty"`
	// Contacts contact list in a group
	Contacts []Contact `json:"contacts,omitempty"`
	// ConcactCount number of contacts in a group
	ContactCount uint `json:"contact_count,omitempty"`
	// CreateTime message create time
	CreateTime int64 `json:"create_time,omitempty"`
	// Group group info
	Group *Group `json:"group,omitempty"`
	// AtList @ list
	AtList string `json:"at_list,omitempty"`
	// Tags message tags
	Tags []string `json:"tags,omitempty"`
}

// EventRequest event send or receive request
type EventRequest struct {
	// Payload encrypted event
	Payload string `form:"payload" json:"payload,omitempty" binding:"required"`
	// Timestamp ts
	Timestamp int64 `form:"timestamp" json:"timestamp,omitempty" binding:"required"`
	// Sign signature
	Sign string `form:"sign" json:"sign,omitempty" binding:"required"`
	// Key apikey
	Key string `form:"key" json:"key,omitempty" binding:"required"`
	// Code response code
	Code int `form:"code" json:"code,omitempty"`
	// ErrMsg response error msg
	ErrMsg string `form:"msg" json:"msg,omitempty"`
}

// IsError check the response is error
func (this *EventRequest) IsError() bool {
	return this.Code != 0
}

// Error implement error interface
func (this EventRequest) Error() string {
	return fmt.Sprintf("CODE: %d, MSG: %s", this.Code, this.ErrMsg)
}

// Events decrypted events from payload
func (this *EventRequest) Events() ([]Event, error) {
	if this.Payload == "" {
		return nil, nil
	}
	var events []Event
	payload := html.UnescapeString(this.Payload)
	err := json.Unmarshal([]byte(payload), &events)
	return events, err
}
