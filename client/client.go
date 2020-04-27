package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/XiBao/groupee-go/model"
)

const (
	DEFAULT_CONNECT_TIMEOUT  time.Duration = 100
	DEFAULT_RESPONSE_TIMEOUT time.Duration = 100
)

type Client struct {
	Key             string
	Secret          string
	connectTimeout  time.Duration
	responseTimeout time.Duration
}

func NewClient(key string, secret string) *Client {
	return &Client{
		Key:             key,
		Secret:          secret,
		connectTimeout:  DEFAULT_CONNECT_TIMEOUT,
		responseTimeout: DEFAULT_RESPONSE_TIMEOUT,
	}
}

func (this *Client) SetConnectTimeout(t time.Duration) {
	this.connectTimeout = t
}

func (this *Client) SetResponseTimeout(t time.Duration) {
	this.responseTimeout = t
}

func (this *Client) Sign(req *model.EventRequest) string {
	rawSign := fmt.Sprintf("%s%d%s%d%s", this.Secret, this.Key, req.Payload, req.Timestamp, this.Secret)
	return Md5(rawSign)
}

func (this *Client) NewRequest(event *model.Event) *model.EventRequest {
	buf, _ := json.Marshal(this)
	req := &model.EventRequest{
		Payload:   string(buf),
		Timestamp: time.Now().Unix(),
		Key:       this.Key,
	}
	req.Sign = this.Sign(req)
	return req
}

func (this *Client) SendEvent(gateway string, event *model.Event) (*model.Event, error) {
	eventReq := this.NewRequest(event)
	buf, _ := json.Marshal(eventReq)
	fmt.Println(string(buf))
	request, err := http.NewRequest("POST", gateway, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	request.Header.Add("content-type", "application/json")
	request.Header.Add("charset", "utf-8")

	clt := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Millisecond*this.connectTimeout)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Millisecond * this.connectTimeout))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Millisecond * this.responseTimeout,
		},
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
	return eventResp.Event()
}
