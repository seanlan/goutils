/**
	gin 数据加密中间件
支持request请求加密、response响应数据加密
支持加密方式有
 AES 对称加密
也可以自己实现 Encryptor 接口来自定义加密方式

request请求加密方式，将请求的json加密后，以表单的形式提交
enc=请求json加密信息

response响应数据加密，将返回的数据加密后返回

*/

package xlhttp

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

const (
	EncryptFormKey = "enc"
)

type Encryptor interface {
	Decrypt(source string) (string, error) //加密
	Encrypt(dec string) (string, error)    //解密
}

func NewEncryptRequestMiddleware(encryptor Encryptor) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data string
		var err error
		enc := c.Request.PostForm.Get(EncryptFormKey)
		if len(enc) > 0 {
			data, err = encryptor.Decrypt(enc)
			if err != nil {
				c.AbortWithStatus(500)
				return
			}
		} else {
			data = "{}"
		}
		c.Request.Header.Set("Content-Type", gin.MIMEJSON)
		c.Set(gin.BodyBytesKey, []byte(data))
		c.Next()
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.Body.Write(b)
	return len(b), nil
}

func NewEncryptResponseMiddleware(encryptor Encryptor) gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{Body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		c.Next()
		// 处理请求
		var response string
		var err error
		response = w.Body.String()
		if len(response) > 0 {
			response, err = encryptor.Encrypt(response)
			if err != nil {
				c.AbortWithStatus(500)
				return
			}
		}
		_, err = c.Writer.WriteString(response)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
	}
}
