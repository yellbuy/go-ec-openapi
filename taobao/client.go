package taobao

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/cache"
	"github.com/yellbuy/go-ec-openapi/common"
)

var (
	router = "http://gw.api.taobao.com/router/rest"
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

func InitClient(appKey, appSecret, session string) *Client {
	client := new(Client)
	client.Params = &common.ClientParams{appKey, appSecret, session}
	return client
}

// execute 执行API接口
func execute(client *Client, param common.Parameter) (bytes []byte, err error) {
	err = checkConfig(client)
	if err != nil {
		return
	}

	var req *http.Request
	req, err = http.NewRequest("POST", router, strings.NewReader(param.GetRequestData()))
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
func (client *Client) GetWaybill(request *common.WaybillApplyNewRequest) (*common.WaybillApplyNewCols, error) {
	req := make(map[string]interface{})
	req["waybill_apply_new_request"] = request
	params, err := common.InterfaceToParameter(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err := client.Execute("taobao.wlb.waybill.i.get", params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	data, err := res.Get("wlb_waybill_i_get_response").Encode()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := new(common.WaybillApplyNewCols)
	err = json.Unmarshal(data, result)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// Execute 执行API接口
// func (client *Client) Execute(method string, param common.Parameter) (res *common.WaybillApplyNewCols, err error) {
// 	param["method"] = method
// 	param.SetRequestData(client.Params)

// 	var bodyBytes []byte
// 	bodyBytes, err = execute(client, param)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	res = new(common.WaybillApplyNewCols)
// 	err = json.Unmarshal(bodyBytes, res)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return res, err
// }

// Execute 执行API接口
func (client *Client) Execute(method string, param common.Parameter) (res *simplejson.Json, err error) {
	param["method"] = method
	param.SetRequestData(client.Params)

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
		fmt.Println("code:", responseError.Get("code").MustInt())
		fmt.Println("msg:", responseError.Get("msg").MustString())
		if subMsg, subOk := responseError.CheckGet("sub_msg"); subOk {
			fmt.Println("sub_code:", responseError.Get("sub_code").MustString())
			fmt.Println("sub_msg:", responseError.Get("sub_msg").MustString())
			err = errors.New(subMsg.MustString())
		} else {
			err = errors.New(responseError.Get("msg").MustString())
		}
		res = nil
	}
	return
}

// ExecuteCache 执行API接口，缓存
func (client *Client) ExecuteCache(method string, param common.Parameter) (res *simplejson.Json, err error) {
	param["method"] = method
	param.SetRequestData(client.Params)

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
		return errors.New("AppKey 不能为空")
	}
	if client.Params.AppSecret == "" {
		return errors.New("AppSecret 不能为空")
	}
	if router == "" {
		return errors.New("router 不能为空")
	}
	return nil
}
