package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/XiBao/groupee-go/model"
)

const (
	DEFAULT_TIMEOUT time.Duration = 300
)

type Client struct {
	Key     string
	Secret  string
	timeout time.Duration
}

func NewClient(key string, secret string) *Client {
	return &Client{
		Key:     key,
		Secret:  secret,
		timeout: DEFAULT_TIMEOUT,
	}
}

func (this *Client) SetTimeout(t time.Duration) {
	this.timeout = t
}

func (this *Client) Sign(req *model.EventRequest) string {
	payload := html.UnescapeString(req.Payload)
	rawSign := fmt.Sprintf("%s%s%s%d%s", this.Secret, this.Key, payload, req.Timestamp, this.Secret)
	return Md5(rawSign)
}

func (this *Client) NewRequest(events []model.Event) *model.EventRequest {
	buf, _ := json.Marshal(events)
	req := &model.EventRequest{
		Payload:   string(buf),
		Timestamp: time.Now().Unix(),
		Key:       this.Key,
	}
	req.Sign = this.Sign(req)
	return req
}

func (this *Client) SendEvent(gateway string, event *model.Event) ([]model.Event, error) {
	eventReq := this.NewRequest([]model.Event{*event})
	buf, _ := json.Marshal(eventReq)
	request, err := http.NewRequest("POST", gateway, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	request.Header.Add("content-type", "application/json")
	request.Header.Add("charset", "utf-8")

	clt := &http.Client{
		Timeout: time.Millisecond * this.timeout,
	}
	resp, err := clt.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var eventResp model.EventRequest
	err = json.NewDecoder(resp.Body).Decode(&eventResp)
	if err != nil {
		return nil, err
	}
	if eventResp.IsError() {
		return nil, eventResp
	}
	if eventResp.Payload != "" && this.Sign(&eventResp) != eventResp.Sign {
		return nil, errors.New("invalid sign")
	}
	return eventResp.Events()
}
