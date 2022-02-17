package jwt

import (
	"errors"
	"go-scaffold/pkg/configs"
	"time"

	"go.uber.org/zap"

	"github.com/dgrijalva/jwt-go"
)

var (
	Secret = []byte(configs.AllConfig.Auth.JWTSecret) // 密钥
)

// 自定义信息载体
type ChronosClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"` // 用户ID
}

// GenerateToken 生成Token
func GenerateToken(userID int64) (tokenStr string, err error) {
	// JWT Token 过期时间
	TokenExpire := time.Duration(configs.AllConfig.Auth.JWTExpire) * time.Second
	// Payload 载荷
	claims := ChronosClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "chronos",
			ExpiresAt: time.Now().Add(TokenExpire).Unix(),
		},
	}

	// 使用指定的加密算法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用指定的密钥签名并获取完整的Token
	return token.SignedString(Secret)
}

// ParseToken 解析Token
func ParseToken(tokenStr string) (chronosClaims *ChronosClaims, err error) {
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenStr, &ChronosClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		zap.L().Error("unable parse token", zap.Error(err))
		return nil, err
	}

	if claims, ok := token.Claims.(*ChronosClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
