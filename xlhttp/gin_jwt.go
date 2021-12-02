package xlhttp

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	JWTIdentityKey   = "jwt_uid"
	RequestTokenHEAD = "X-JWT-TOKEN"
)

// JWTBodyMiddleware jwt token 位于请求内
func JWTBodyMiddleware(secretKey string, d time.Duration) gin.HandlerFunc {
	jwt := NewJWT(secretKey, d)
	return func(c *gin.Context) {
		r := Build(c)
		var req struct {
			Token string `form:"token" json:"token" binding:"required"`
		}
		err := r.RequestParser(&req)
		if err != nil {
			r.ctx.Abort()
			return
		}
		jwtUid, err := jwt.ParseToken(req.Token)
		if err != nil {
			r.JsonReturn(ErrToken)
			r.ctx.Abort()
			return
		}
		if jwtUid == "" {
			r.JsonReturn(ErrToken)
			r.ctx.Abort()
			return
		}
		r.ctx.Set(JWTIdentityKey, jwtUid)
		r.ctx.Next()
	}
}

// JWTHeadMiddleware jwt token 位于请求头
func JWTHeadMiddleware(secretKey string, d time.Duration) gin.HandlerFunc {
	jwt := NewJWT(secretKey, d)
	return func(c *gin.Context) {
		r := Build(c)
		jwtUid, err := jwt.ParseToken(c.GetHeader(RequestTokenHEAD))
		if err != nil {
			r.JsonReturn(ErrToken)
			r.ctx.Abort()
			return
		}
		if jwtUid == "" {
			r.JsonReturn(ErrToken)
			r.ctx.Abort()
			return
		}
		r.ctx.Set(JWTIdentityKey, jwtUid)
		r.ctx.Next()
	}
}
