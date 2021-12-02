package feishu

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/seanlan/goutils/xljson"
	"io/ioutil"
	"net/http"
)

func New(hook, secret string) *Client {
	return &Client{
		Hook:   hook,
		Secret: secret,
	}
}

type Client struct {
	Hook   string
	Secret string
}

// sendMsg 发送消息
func (f Client) sendMsg(msg interface{}) (result xljson.JsonObject, err error) {
	if f.Hook == "" {
		err = errors.New("no specified webhook")
		return
	}
	data, err := json.Marshal(msg)
	if err != nil {
		err = errors.New("marshal data error")
		return
	}
	resp, err := http.Post(f.Hook, "application/json", bytes.NewReader(data))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// 保存在[]byte中复用
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	result = xljson.JsonObject{Buff: d}
	return
}

func (f Client) SendText(msg string) (result xljson.JsonObject, err error) {
	var message *TextMsg
	if len(f.Secret) > 0 {
		message = NewTextMsgWithSign(f.Secret, msg)
	} else {
		message = NewTextMsg(msg)
	}
	return f.sendMsg(message)
}

func (f Client) SendCard(title string, elements ...string) (result xljson.JsonObject, err error) {
	var message *CardMsg
	if len(f.Secret) > 0 {
		message = NewCardMsgWithSign(f.Secret, title)
	} else {
		message = NewCardMsg(title)
	}
	for _, ele := range elements {
		message.AddElement(ele)
	}
	return f.sendMsg(message)
}

func (f Client) SendPost(title string, tags ...Tag) (result xljson.JsonObject, err error) {
	var message *PostMsg
	if len(f.Secret) > 0 {
		message = NewPostMsgWithSign(f.Secret, title)
	} else {
		message = NewPostMsg(title)
	}
	for _, tag := range tags {
		content := message.Content.Post.ZhCn.Content
		content[0] = append(content[0], tag)
	}
	return f.sendMsg(message)
}
