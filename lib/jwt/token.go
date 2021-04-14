// @Title        token
// @Description  有关 JWT 的处理
// @Author       Eachin
// @Date         2021/4/3 4:36 上午

package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"oauth/constant/exception"
	"oauth/logger"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
type MyClaims struct {
	UserID int64 `json:"uid"`
	jwt.StandardClaims
}

func NewClaims(userID int64) *MyClaims {
	return &MyClaims{
		UserID: userID,
	}
}

// GenToken 生成JWT
func GenToken(ctx *gin.Context, claims *MyClaims, secret string, expiresAt int) (string, error) {
	fmt.Println(claims)
	claims.ExpiresAt = time.Now().Add(time.Duration(expiresAt) * time.Hour).Unix()
	claims.Issuer = "Kaimon.cn"
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
func ParseToken(ctx *gin.Context, tokenString string, secret string) (*MyClaims, error) {
	// 解析token
	var claims = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		zap.L().Warn("ParseWithClaimsErr", zap.Error(err))
		return nil, exception.ErrorServerBusy
	}
	if !token.Valid {
		return nil, exception.ErrorServerBusy
	}
	return claims, nil
}
