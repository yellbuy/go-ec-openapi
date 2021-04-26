package polyapi

import (
	"encoding/json"
	"errors"
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

// Token
type AccessToken struct {
	// 授权SessionKey
	SessionKey string `json:"sessionkey"`
	// 刷新SessionKey(当过期时)所需RefreshToken
	RefreshToken string `json:"refreshtoken"`
	//SessionKey过期时间
	SessionKeyExpireTime string `json:"sessionkeyexpiretime"`
	//RefreshToken过期时间
	RefreshTokenExpireTime string `json:"refreshtokenexpiretime"`
	// SessionKey过期时间间隔(单位：秒)
	ExpireTimeInterval int `json:"expiretimeinterval"`
	// 店铺ID
	VenderId string `json:"venderid"`
	//用户昵称
	Nickname string `json:"nickname"`
	//店铺类别（B代表天猫店铺，C代表 集市卖家
	Subplat string   `json:"subplat"`
	Scope   []string `json:"scope"`
}

type authorizeReqDto struct {
	AppKey      string `json:"appkey"`
	AppSecret   string `json:"appsecret"`
	CallbackUrl string `json:"callbackurl"`
	State       string `json:"state"`
	ItemCode    string `json:"itemcode"`
}

type tokenReqDto struct {
	AppKey      string `json:"appkey"`
	AppSecret   string `json:"appsecret"`
	CallbackUrl string `json:"callbackurl"`
	// 是否校验服务订购状态
	IsCheckAppOrder bool `json:"ischeckapporder"`
}

func (client *Client) GetAuthorizeUrl(redirectUri, state string, extData ...string) (string, error) {
	reqDto := new(authorizeReqDto)
	if len(extData) > 0 {
		reqDto.AppKey = extData[0]
	}
	if len(extData) > 1 {
		reqDto.AppSecret = extData[1]
	}
	reqDto.CallbackUrl = redirectUri
	reqDto.State = state
	bizcontent, err := json.Marshal(reqDto)
	if err != nil {
		fmt.Println(err)
	}

	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, _, err := client.Execute("Differ.JH.BuildeAuthorizeUrl", params)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	authorizeUrl := res.Get("authorizeurl").MustString()
	if authorizeUrl == "" {
		return "", errors.New("authorizeUrl为空")
	}
	return authorizeUrl, nil
}

func (client *Client) GetAccessToken(code, redirectUri, state string, extData ...string) (res *common.AccessToken, body []byte, err error) {
	reqDto := new(tokenReqDto)
	platAppKey, platAppSecret := "", ""
	if len(extData) > 0 {
		platAppKey = extData[0]
	}
	if len(extData) > 1 {
		platAppSecret = extData[1]
	}
	reqDto.AppKey = platAppKey
	reqDto.AppSecret = platAppSecret
	reqDto.CallbackUrl = redirectUri
	reqDto.IsCheckAppOrder = false
	bizcontent, resErr := json.Marshal(reqDto)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return res, body, err
	}

	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return res, body, err
	}
	var resJson *simplejson.Json
	// 获取平台SessionKey
	resJson, body, err = client.Execute("Differ.JH.GetAuthorizeSessionKey", params)
	if err != nil {
		fmt.Println("Differ.JH.GetAuthorizeSessionKey:", err)
		return res, body, err
	}

	res = new(common.AccessToken)
	res.AccessToken = resJson.Get("sessionkey").MustString()
	res.ExpireIn = resJson.Get("expiretimeinterval").MustInt()
	res.RefreshToken = resJson.Get("refreshtoken").MustString()
	res.VenderId = resJson.Get("venderid").MustString()
	res.Nickname = resJson.Get("nickname").MustString()
	res.ExpireTime = resJson.Get("sessionkeyexpiretime").MustString()
	res.RefreshTokenExpireTime = resJson.Get("refreshtokenexpiretime").MustString()
	res.Subplat = resJson.Get("subplat").MustString()
	res.TokenType = "POLYAPI"

	// 平台账号同步
	resJson = simplejson.New()

	resJson.Set("AppKey", platAppKey)
	resJson.Set("AppSecret", platAppSecret)
	resJson.Set("SessionKey", res.AccessToken)
	resJson.Set("SessionKeyExpireTime", res.ExpireTime)
	resJson.Set("SessionKeyTimeout", res.ExpireIn)
	resJson.Set("RefreshTokenKey", res.RefreshToken)
	// params, err = resJson.Map()

	req = make(map[string]interface{})
	bizcontent, err = resJson.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return res, body, err
	}
	//fmt.Println("bizcontent:", string(bizcontent))
	bizcontent, err = aesEncrypt(client.Params.AppSecret, bizcontent)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := byteArrToHexString(bizcontent)
	//fmt.Println("bizcontent:", content)
	req["bizcontent"] = content
	resJson, body, err = client.Execute("Differ.JH.SyncAccount", req)
	if err != nil {
		fmt.Println("Differ.JH.SyncAccount", err)
		return res, body, err
	}
	// 获取最终token
	res.ThirdPlatToken = resJson.Get("token").MustString()
	return res, body, err
}
