package authenticator

import (
	"crypto/sha256"
	"fmt"
	"time"
)

//把 URL、AppID、密码、时间戳拼接为一个字符串；c
//对字符串通过加密算法加密生成 token；c
//将 token、AppID、时间戳拼接到 URL 中，形成新的 URL；
//解析 URL，得到 token、AppID、时间戳等信息；
//从存储中取出 AppID 和对应的密码；
//根据时间戳判断 token 是否过期失效；c
//验证两个 token 是否匹配；c
const DefaultExpiredTimeInterval int64 = 3600

type AuthToken struct {
	Token string
	CreatedTime int64
	ExpiredTimeInterval int64
}

type AuthTokenOption struct {
	ExpiredTimeInterval *int64
}

// NewAuthToken 构造函数
func NewAuthToken(token string, createdTime int64, option AuthTokenOption) *AuthToken {
	expiredTimeInterval := DefaultExpiredTimeInterval
	if option.ExpiredTimeInterval != nil {
		expiredTimeInterval = *option.ExpiredTimeInterval
	}

	return &AuthToken{
		Token:               token,
		CreatedTime:         createdTime,
		ExpiredTimeInterval: expiredTimeInterval,
	}
}


func GenerateToken(url, secret string) string {
	//Generate Token by sha256
	hash := sha256.Sum256([]byte(url+secret))
	token := fmt.Sprintf("%x", hash[:])
	return token
}

func (a *AuthToken) IsExpired() bool {
	return time.Now().Unix() > a.CreatedTime + a.ExpiredTimeInterval
}

func (a *AuthToken) Match(authToken AuthToken) bool {
	return a.Token == authToken.Token
}