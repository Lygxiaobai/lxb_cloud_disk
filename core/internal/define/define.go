package define

import "github.com/golang-jwt/jwt/v4"

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
