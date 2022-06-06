package authenticator

import "fmt"

//把 URL、AppID、密码、时间戳拼接为一个字符串；c
//对字符串通过加密算法加密生成 token；c
//将 token、AppID、时间戳拼接到 URL 中，形成新的 URL；
//解析 URL，得到 token、AppID、时间戳等信息；c
//从存储中取出 AppID 和对应的密码；
//根据时间戳判断 token 是否过期失效；c
//验证两个 token 是否匹配；c

type Authenticator interface {
	Auth(request ApiRequest) error
	AuthUrl(rawUrl string) error
}


type DefaultAuthenticator struct {
	Cs CredentialStorage
}



func NewDefaultAuthenticator (storage CredentialStorage) *DefaultAuthenticator {
	return &DefaultAuthenticator{
		Cs: NewMySQlSto(),
	}
}

func (d *DefaultAuthenticator) Auth(request ApiRequest) error {
	originalToken := request.Token
	appId := request.AppId
	ts := request.TimeStamp
	pwd := d.Cs.GetSecretByAppid(appId)
	authToken := NewAuthToken(originalToken, ts, AuthTokenOption{})
	if authToken.IsExpired() {
		return fmt.Errorf("token is expired")
	}
	newToken := GenerateToken(request.RawUrl, pwd)
	newAuthToken := NewAuthToken(newToken, ts, AuthTokenOption{})
	if !newAuthToken.Match(*authToken) {
		return fmt.Errorf("token is malformd")
	}
	return nil
}

func (d *DefaultAuthenticator) AuthUrl(rawUrl string) error {
	apiRequset, err := NewApiRequest(rawUrl)
	if err != nil {
		return err
	}
	return d.Auth(*apiRequset)
}