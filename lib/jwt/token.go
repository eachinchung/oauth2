// @Title        token
// @Description  有关 JWT 的处理
// @Author       Eachin
// @Date         2021/4/3 4:36 上午

package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"oauth/constant/exception"
	"oauth/logger"
	"time"
)

type Claims interface {
	Valid() error
	InitClaims(expires time.Duration)
}

// BaseClaims 自定义声明结构体并内嵌jwt.StandardClaims
type BaseClaims struct {
	UserID int64 `json:"uid"`
	jwt.StandardClaims
}

func (m *BaseClaims) InitClaims(expires time.Duration) {
	m.ExpiresAt = time.Now().Add(expires).Unix()
	m.Issuer = "Kaimon.cn"
}

func NewClaims(userID int64) *BaseClaims {
	return &BaseClaims{
		UserID: userID,
	}
}

// GenToken 生成JWT
func GenToken(ctx *gin.Context, claims Claims, secret string, expires time.Duration) (string, error) {
	claims.InitClaims(expires)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString([]byte(secret)); err != nil {
		log := logger.NewSentry(ctx)
		log.Error("生成 Token 失败", err)
		return "", exception.ErrorServerBusy
	} else {
		return tokenString, nil
	}
}

// ParseToken 解析JWT
func ParseToken(tokenString string, claims Claims, secret string) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		zap.L().Debug("token 验证出错", zap.Error(err))
		return exception.ErrorInvalidToken
	}
	if !token.Valid {
		zap.L().Debug("token 验证不通过", zap.Any("token", token))
		return exception.ErrorInvalidToken
	}
	return nil
}
