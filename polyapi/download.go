package polyapi

import (
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

// 产品下载
func (client *Client) DownloadProductList(pageIndex, pageSize int, status string, extData ...string) (res *simplejson.Json, body []byte, err error) {
	reqJson := simplejson.New()
	reqJson.Set("pageindex", pageIndex)
	reqJson.Set("pagesize", pageSize)
	reqJson.Set("status", status)
	reqJson.Set("polyapitoken", client.Params.Session)
	if len(extData) > 0 {
		reqJson.Set("platvalue", extData[0])
	}

	bizcontent, resErr := reqJson.Encode()
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
	// 获取平台SessionKey
	res, body, err = client.Execute("Differ.JH.Other.DelegateQimenDownloadProduct", params)
	if err != nil {
		fmt.Println("Differ.JH.Other.DelegateQimenDownloadProduct:", err)
		return res, body, err
	}
	return res, body, err
}

// 订单下载
func (client *Client) DownloadOrderList(pageIndex, pageSize int, startTime, endTime, timeType, orderStatus string, extData ...string) (res *simplejson.Json, body []byte, err error) {
	reqJson := simplejson.New()
	reqJson.Set("pageindex", pageIndex)
	reqJson.Set("pagesize", pageSize)
	reqJson.Set("orderstatus", orderStatus)
	reqJson.Set("starttime", startTime)
	reqJson.Set("endtime", endTime)
	reqJson.Set("timetype", timeType)
	reqJson.Set("polyapitoken", client.Params.Session)
	reqJson.Set("platvalue", client.Params.PlatId)
	bizcontent, resErr := reqJson.Encode()
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
	// 获取平台SessionKey
	res, body, err = client.Execute(" Differ.JH.Other.DelegateQimenGetOrder", params)
	if err != nil {
		fmt.Println(" Differ.JH.Other.DelegateQimenGetOrder:", err)
		return res, body, err
	}
	return res, body, err
}
