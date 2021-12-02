package feishu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)

// 签名

type Sign struct {
	Timestamp string `json:"timestamp,omitempty"`
	Sign      string `json:"sign,omitempty"`
}

func NewSign(secret string) *Sign {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	sign := GenSign(secret, timestamp)
	return &Sign{
		Timestamp: timestamp,
		Sign:      sign,
	}
}

func GenSign(secret, timestamp string) string {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := timestamp + "\n" + secret
	h := hmac.New(sha256.New, []byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}
