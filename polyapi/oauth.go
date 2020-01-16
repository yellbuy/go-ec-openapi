package polyapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

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

func (client *Client) GetAuthorizeUrl(redirectUri, state string) (string, error) {
	return "", nil
}

func (client *Client) GetAccessToken(code, redirectUri, state string) (res *common.AccessToken, body []byte, err error) {
	postData := url.Values{"client_id": {client.Params.AppKey}, "client_secret": {client.Params.AppSecret}, "code": {code},
		"grant_type": {"authorization_code"}, "redirect_uri": {redirectUri}, "state": {state}}
	var resp *http.Response
	resp, err = http.Post("http://open-api.pinduoduo.com/oauth/token", "application/x-www-form-urlencoded;charset=utf-8", strings.NewReader(postData.Encode()))
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	token := new(AccessToken)
	if err = json.Unmarshal(body, token); err != nil {
		fmt.Println("body:", err)
		return
	}
	if token.SessionKey == "" {
		err = errors.New("access_token不存在")
		return
	}
	postData = url.Values{"client_id": {client.Params.AppKey}, "client_secret": {client.Params.AppSecret}, "code": {code},
		"grant_type": {"refresh_token"}, "refresh_token": {token.RefreshToken}, "redirect_uri": {redirectUri}, "state": {state}}
	resp, err = http.Post("http://open-api.pinduoduo.com/oauth/token", "application/x-www-form-urlencoded;charset=utf-8", strings.NewReader(postData.Encode()))
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = json.Unmarshal(body, token); err != nil {
		fmt.Println("body:", err)
		return
	}
	res = new(common.AccessToken)
	res.AccessToken = token.SessionKey
	res.ExpireIn = token.ExpireTimeInterval
	res.RefreshToken = token.RefreshToken
	res.VenderId = token.VenderId
	res.Nickname = token.Nickname
	res.TokenType = "POLYAPI"
	res.Scope = token.Scope
	return
}
