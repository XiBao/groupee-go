package model

type Group struct {
	Id     uint64            `json:"id,omitempty"`
	Name   string            `json:"name,omitempty"`
	Tracks map[string]string `json:"tracks,omitempty"`
}

type Contact struct {
	Id   string `json:"id,omitempty"`
	Nick string `json:"nick,omitempty"`
}
