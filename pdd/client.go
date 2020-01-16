package pdd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/cache"
	"github.com/yellbuy/go-ec-openapi/common"
)

var (
	router = "https://gw-api.pinduoduo.com/api/router"
	// router = "http://gw.api.tbsandbox.com/router/rest"
	// router = "https://baidu.com"
	// Timeout ...
	Timeout time.Duration
	// CacheExpiration 缓存过期时间
	CacheExpiration = time.Hour
	// GetCache 获取缓存
	GetCache cache.GetCacheFunc
	// SetCache 设置缓存
	SetCache cache.SetCacheFunc
)

type Client struct {
	Params *common.ClientParams
}

func InitClient(clientId, clientSecret, access_token string) *Client {
	client := new(Client)
	client.Params = &common.ClientParams{clientId, clientSecret, access_token}
	return client
}
func setRequestData(p common.Parameter, params *common.ClientParams) common.Parameter {
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	p["timestamp"] = strconv.FormatInt(loc.Unix(), 10)
	p["data_type"] = "json"
	p["version"] = "v1"
	if params.Session != "" {
		p["access_token"] = params.Session
	}
	// 设置签名
	p["sign"] = common.GetSign(params.AppSecret, p)
	return p
}

// execute 执行API接口
func execute(client *Client, param common.Parameter) (bytes []byte, err error) {
	err = checkConfig(client)
	if err != nil {
		return
	}

	var req *http.Request
	data := param.GetRequestData()
	req, err = http.NewRequest("POST", router, strings.NewReader(data))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	httpClient := &http.Client{}
	httpClient.Timeout = Timeout
	var response *http.Response
	response, err = httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}
	defer response.Body.Close()
	bytes, err = ioutil.ReadAll(response.Body)
	return
}

// Execute 执行API接口
func (client *Client) Execute(method string, param common.Parameter) (res *simplejson.Json, err error) {
	param["type"] = method
	param = setRequestData(param, client.Params)

	var bodyBytes []byte
	bodyBytes, err = execute(client, param)
	if err != nil {
		fmt.Println(err)
		return
	}

	return bytesToResult(bodyBytes)
}

func bytesToResult(bytes []byte) (res *simplejson.Json, err error) {
	res, err = simplejson.NewJson(bytes)
	if err != nil {
		return
	}

	if responseError, ok := res.CheckGet("error_response"); ok {
		fmt.Println("error_code:", responseError.Get("error_code").MustInt())
		fmt.Println("error_msg:", responseError.Get("error_msg").MustString())
		if subMsg, subOk := responseError.CheckGet("sub_msg"); subOk {
			fmt.Println("sub_code:", responseError.Get("sub_code").MustString())
			fmt.Println("sub_msg:", responseError.Get("sub_msg").MustString())
			err = errors.New(subMsg.MustString())
		} else {
			err = errors.New(responseError.Get("error_msg").MustString())
		}
		res = nil
	}
	return
}

// ExecuteCache 执行API接口，缓存
func (client *Client) ExecuteCache(method string, param common.Parameter) (res *simplejson.Json, err error) {
	param["method"] = method
	param = setRequestData(param, client.Params)

	cacheKey := common.NewCacheKey(param)
	// 获取缓存
	if GetCache != nil {
		cacheBytes := GetCache(cacheKey)
		if len(cacheBytes) > 0 {
			res, err = simplejson.NewJson(cacheBytes)
			if err == nil && res != nil {
				return
			}
		}
	}

	var bodyBytes []byte
	bodyBytes, err = execute(client, param)
	if err != nil {
		return
	}
	res, err = bytesToResult(bodyBytes)
	if err != nil {
		return
	}
	ejsonBody, _ := res.MarshalJSON()
	// 设置缓存
	if SetCache != nil {
		go SetCache(cacheKey, ejsonBody, CacheExpiration)
	}
	return
}

// 检查配置
func checkConfig(client *Client) error {
	if client.Params.AppKey == "" {
		return errors.New("clientId 不能为空")
	}
	if client.Params.AppSecret == "" {
		return errors.New("clientSecret 不能为空")
	}
	if router == "" {
		return errors.New("router 不能为空")
	}
	return nil
}
