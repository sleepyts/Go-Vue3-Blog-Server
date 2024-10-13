package jwtutil

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

var (
	tokenExpireDuration time.Duration = time.Hour * 24 * 1 // token过期时间
)

// 生成token
func GenToken(userid int64, userName string) (string, error) {
	claims := JWTClaims{
		UserId:   userid,
		Username: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "server",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)), // 设置过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("Blog-Server-Secret"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Blog-Server-Secret"), nil
	})
	if err != nil {
		return "", err
	}
	if _, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return token.Claims.(*JWTClaims).Username, nil
	}
	return "", err
}
