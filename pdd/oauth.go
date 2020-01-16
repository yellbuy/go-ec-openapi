package pdd

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

type AccessToken struct {
	AccessToken  string   `json:"access_token"`
	ExpiresIn    int      `json:"expires_in"`
	OwnerId      string   `json:"owner_id"`
	OwnerName    string   `json:"owner_name"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
}

func (client *Client) GetAuthorizeUrl(redirectUri, state string) (string, error) {
	oauthUrl := fmt.Sprint("https://mms.pinduoduo.com/open.html?response_type=code&client_id=%v&redirect_uri=%v&state=%v", client.Params.AppKey, redirectUri, state)
	return oauthUrl, nil
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
	if token.AccessToken == "" {
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
	res.AccessToken = token.AccessToken
	res.ExpireIn = token.ExpiresIn
	res.RefreshToken = token.RefreshToken
	res.VenderId = token.OwnerId
	res.Nickname = token.OwnerName
	res.TokenType = "PDD"
	res.Scope = token.Scope
	return
}
