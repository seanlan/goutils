package xlhttp

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	SigningKey []byte
	Expire     time.Duration
}

func NewJWT(secretKey string, d time.Duration) *JWT {
	return &JWT{
		SigningKey: []byte(secretKey),
		Expire:     d,
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(data string) (string, error) {
	now := time.Now()
	claims := &jwt.StandardClaims{
		Audience:  "",
		IssuedAt:  now.Unix(),
		Issuer:    "TimeToken",
		NotBefore: 0,
		Subject:   data,
	}
	if j.Expire > 0 {
		claims.ExpiresAt = now.Add(j.Expire).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(token string) (string, error) {
	var err error
	var claims jwt.StandardClaims
	_, err = jwt.ParseWithClaims(
		token,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return j.SigningKey, nil
		})
	if err != nil {
		return "", err
	}
	err = claims.Valid()
	if err != nil {
		return "", err
	} else {
		return claims.Subject, err
	}
}
