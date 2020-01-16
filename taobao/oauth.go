package taobao

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`

	W1ExpiresIn int    `json:"w1_expires_in"`
	W2ExpiresIn int    `json:"w2_expires_in"`
	ReExpiresIn int    `json:"re_expires_in"`
	R1ExpiresIn int    `json:"r1_expires_in"`
	R2ExpiresIn int    `json:"r2_expires_in"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`

	TaobaoUserId   string   `json:"taobao_user_id"`
	TaobaoUserNick string   `json:"taobao_user_nick"`
	RefreshToken   string   `json:"refresh_token"`
	Scope          []string `json:"scope"`
}

func (client *Client) GetAuthorizeUrl(redirectUri, state string) (string, error) {
	query := url.Values{}
	query.Add("response_type", "code")
	query.Add("client_id", client.Params.AppKey)
	query.Add("redirect_uri", redirectUri)
	query.Add("state", fmt.Sprintf("%v", state))
	queryStr := query.Encode()
	oauthUrl := fmt.Sprintf("https://oauth.taobao.com/authorize?%s", queryStr)
	return oauthUrl, nil
}

func (client *Client) GetAccessToken(code, redirectUri, state string) (res *common.AccessToken, body []byte, err error) {
	postData := url.Values{"grant_type": {"authorization_code"}, "client_id": {client.Params.AppKey}, "client_secret": {client.Params.AppSecret},
		"code": {code}, "redirect_uri": {redirectUri}}
	postStr := postData.Encode()
	var resp *http.Response
	resp, err = http.Post("https://oauth.taobao.com/token", "application/x-www-form-urlencoded;charset=utf-8", strings.NewReader(postStr))
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
	var resJson *simplejson.Json
	resJson, err = simplejson.NewJson(body)
	if err != nil {
		return
	}
	if responseError, ok := resJson.CheckGet("error"); ok {
		fmt.Println("error:", responseError.MustString())
		errmsg := resJson.Get("error_description").MustString()
		fmt.Println("error_description:", errmsg)
		err = errors.New(errmsg)
		return
	}
	token := new(AccessToken)
	if err = json.Unmarshal(body, token); err != nil {
		fmt.Println("body:", err)
		return
	}
	//fmt.Println("access_token body:", string(body))
	if token.AccessToken == "" {
		err = errors.New("access_token不存在")
	}
	postData = url.Values{"grant_type": {"refresh_token"}, "refresh_token": {token.RefreshToken},
		"client_id": {client.Params.AppKey}, "client_secret": {client.Params.AppSecret}}
	resp, err = http.Post("https://oauth.taobao.com/token", "application/x-www-form-urlencoded;charset=utf-8", strings.NewReader(postData.Encode()))
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
	//fmt.Println(string(body))
	resJson, err = simplejson.NewJson(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if responseError, ok := resJson.CheckGet("error"); ok {
		fmt.Println("error:", responseError.MustString())
		errmsg := resJson.Get("error_description").MustString()
		fmt.Println("error_description:", errmsg)
		return nil, body, errors.New(errmsg)
	}

	if err := json.Unmarshal(body, token); err != nil {
		fmt.Println("body:", err)
		return nil, body, err
	}
	res = new(common.AccessToken)
	res.AccessToken = token.AccessToken
	res.ExpireIn = token.ExpiresIn
	res.RefreshToken = token.RefreshToken
	res.VenderId = token.TaobaoUserId
	res.Nickname = token.TaobaoUserNick
	res.TokenType = token.TokenType
	res.Scope = token.Scope
	return
}
