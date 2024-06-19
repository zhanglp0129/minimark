package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"minimark/config"
	"time"
)

// CreateJwtToken 创建jwt令牌，传入jwt令牌中存储的数据。只存储关键数据即可，不要存储太多的数据
func CreateJwtToken(data map[string]any) (string, error) {
	cfg := config.GetConfig()
	claims := jwt.MapClaims{}
	for key, value := range data {
		claims[key] = value
	}
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.Security.JwtExpire) * time.Second))
	claims["iss"] = "minimark"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.Security.JwtKey))
}

// ParseJwtToken 解析jwt令牌，并校验
func ParseJwtToken(tokenString string) (*jwt.Token, error) {
	cfg := config.GetConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Security.JwtKey), nil
	})

	return token, err
}
