package model

// Group group struct
type Group struct {
	// Id group id
	Id Uint64 `json:"id,omitempty"`
	// Name group name
	Name string `json:"name,omitempty"`
	// IsWork 是否企业微信
	IsWork uint `json:"is_work,omitempty"`
	// Tracks tracking params
	Tracks map[string]string `json:"tracks,omitempty"`
}

// Contact contact struct
type Contact struct {
	// Id contact id
	Id string `json:"id,omitempty"`
	// Nick contack nick
	Nick string `json:"nick,omitempty"`
	// Avatar contact avatar
	Avatar string `json:"avatar,omitempty"`
}
