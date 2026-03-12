package helper

import (
	"cloud_disk/core/internal/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"path"
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

// 文件上传到阿里云
func FileUpload(r *http.Request) (string, error) {

	// 获取要上传的文件
	file, header, err := r.FormFile("file")
	var objectName = "cloud-disk/" + UUID() + path.Ext(header.Filename)

	// 检查bucket名称是否为空
	if len(define.BucketName) == 0 {
		flag.PrintDefaults()
		log.Fatalf("invalid parameters, bucket name required")
	}

	// 检查region是否为空
	if len(define.Region) == 0 {
		flag.PrintDefaults()
		log.Fatalf("invalid parameters, region required")
	}

	// 检查object名称是否为空
	if len(objectName) == 0 {
		flag.PrintDefaults()
		log.Fatalf("invalid parameters, object name required")
	}

	//读取文件 获取文件字节流
	//imageBytes, err := os.ReadFile("../images/1.jpg")

	body := file

	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion(define.Region)

	// 创建OSS客户端
	client := oss.NewClient(cfg)

	// 创建上传对象的请求
	request := &oss.PutObjectRequest{
		Bucket: oss.Ptr(define.BucketName), // 存储空间名称
		Key:    oss.Ptr(objectName),        // 对象名称
		Body:   body,                       // 要上传的字符串内容
	}

	// 发送上传对象的请求
	result, err := client.PutObject(context.TODO(), request)
	if err != nil {
		log.Fatalf("failed to put object %v", err)
	}

	// 打印上传对象的结果
	log.Printf("put object result:%#v\n", result)
	return "https://" + define.BucketName + ".oss-" + define.Region + ".aliyuncs.com" + "/" + objectName, nil
}

// 解析token信息
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return nil, errors.New("invalid token")
	}

	return uc, err
}
