package model

type Group struct {
	Id     Uint64            `json:"id,omitempty"`
	Name   string            `json:"name,omitempty"`
	Tracks map[string]string `json:"tracks,omitempty"`
}

type Contact struct {
	Id     string `json:"id,omitempty"`
	Nick   string `json:"nick,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}
