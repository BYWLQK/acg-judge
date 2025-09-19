package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY_USERID = "user_id"
var JWT_KEY_EXP = "exp"
var JWT_SECRET = "bywlqk_acg-judge"

func GenerateToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[JWT_KEY_USERID] = userId
	claims[JWT_KEY_EXP] = time.Now().Add(time.Hour * 24).Unix() // 有效期24小时

	return token.SignedString([]byte(JWT_SECRET))
}

func ParseUserIDFromToken(tokenString string) (string, error) {
	// 解析 Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil // 使用密钥验证
	})
	if err != nil || !token.Valid {
		return "", errors.New("无效 Token")
	}

	// 提取 Claims 中的 user_id
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("无法解析 Claims")
	}
	userID, exists := claims[JWT_KEY_USERID].(string)
	if !exists {
		return "", errors.New("user_id 字段不存在")
	}
	return userID, nil
}
