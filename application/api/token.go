package api

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

// 生成token，在客户端与服务器之间通讯时使用
func GenerateToken(Username string, Hostname string, uid uint32, ip string, secretKey []byte) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"Hostname": Hostname,
		"Username": Username,
		"uid":      strconv.Itoa(int(uid)),
		"ip":       ip,
		"exp":      expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token == nil {
		return "", errors.New("null ptr!")
	} else {
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}

}
