package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

const dingTalkGateHost = "https://oapi.dingtalk.com/robot/send"

func New(accessToken, secret string) *Client {
	return &Client{accessToken, secret}
}

type Client struct {
	accessToken string
	secret      string
}

type DingTalkResponse struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (c *Client) SendMessage(msg string) (resp DingTalkResponse, err error) {
	jsonObject := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	timestamp := time.Now().UnixNano() / 1e6
	h := hmac.New(sha256.New, []byte(c.secret))
	sourceStr := fmt.Sprintf("%d\n%s", timestamp, c.secret)
	h.Write([]byte(sourceStr))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
	apiUrl := fmt.Sprintf("%s?access_token=%s&timestamp=%d&sign=%s",
		dingTalkGateHost, c.accessToken, timestamp, sign)
	request := gorequest.New()
	_, body, errs := request.Post(apiUrl).
		Set("Content-Type", "application/json").
		Send(jsonObject).End()
	if len(errs) > 0 {
		err = errs[0]
	} else {
		_ = json.Unmarshal([]byte(body), &resp)
		err = nil
	}
	return resp, err
}
