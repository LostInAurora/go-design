package authenticator

import (
	"net/url"
	"strconv"
)

//把 URL、AppID、密码、时间戳拼接为一个字符串；c
//对字符串通过加密算法加密生成 token；c
//将 token、AppID、时间戳拼接到 URL 中，形成新的 URL；
//解析 URL，得到 token、AppID、时间戳等信息；c
//从存储中取出 AppID 和对应的密码；
//根据时间戳判断 token 是否过期失效；c
//验证两个 token 是否匹配；c

type ApiRequest struct {
	Url *url.URL
	RawUrl string
	Token string
	AppId string
	TimeStamp int64
}


func NewApiRequest(rawUrl string) (*ApiRequest, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	m, err := url.ParseQuery(u.RawQuery)
	ts, err := strconv.Atoi(m.Get("ts"))
	if err != nil {
		return nil ,err
	}
	return &ApiRequest{
		Url: u,
		RawUrl: rawUrl,
		Token: m.Get("token"),
		AppId: m.Get("appId"),
		TimeStamp: int64(ts),
	}, nil
}