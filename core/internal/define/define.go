package define

import "github.com/golang-jwt/jwt/v4"

// 用于生成Token
type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"
var MailPassword = "ZPjMZku8J3884Q4H"

// 验证码长度
var CodeLen = 6

// 过期时间 s
var CodeExpireTime = 60

var Region = "cn-hangzhou"
var BucketName = "lvxiaobai"

// 分页参数
var Page = 1
var PageSize = 10

// 分享文件过期时间 s
var FileExpireTime = 60 * 60 * 24

// 点击次数
var DefaultClickNum = 0
