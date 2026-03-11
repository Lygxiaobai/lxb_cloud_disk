package helper

import (
	"cloud_disk/core/internal/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"time"
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

// 发送邮箱验证码
func MailCodeSend(userEmail string, code string) error {
	//定义随机种子
	e := email.NewEmail()
	e.From = "Jordan Wright <18163688304@163.com>"
	e.To = []string{userEmail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("你的验证码为：<h1> " + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "18163688304@163.com", define.MailPassword,
		"smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil

}

// 生成随机验证码
func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLen; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

// 生成UUID
func UUID() string {
	return uuid.NewV4().String()
}
