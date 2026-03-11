package helper

import (
	"cloud_disk/core/internal/define"
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

// MD5转换
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// Token生成
func GenerateToken(id int, identity string, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
