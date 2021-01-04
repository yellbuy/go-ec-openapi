package polyapi

import (
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"

)

// 订单下载
func (client *Client) DownloadOrderListByQimen(pageIndex, pageSize int, startTime, endTime, timeType, orderStatus string, extData ...string) (hasNextPage bool, body []byte, err error) {
	defer func() {
		e := recover()
		if e == nil {
			return
		} else {
			err = fmt.Errorf("%s", e)
		}
	}()
	hasNextPage = false
	reqJson := simplejson.New()
	reqJson.Set("isuseinterface", true)
	reqJson.Set("pageindex", pageIndex)
	reqJson.Set("pagesize", pageSize)
	reqJson.Set("starttime", startTime)
	reqJson.Set("endtime", endTime)
	reqJson.Set("timetype", timeType)
	//reqJson.Set("token", "ed7d59ddb5a74df0a63d7307cea0435f")
	reqJson.Set("orderstatus", orderStatus)

	if len(extData) > 0 && extData[0] != "" {
		reqJson.Set("platvalue", extData[0])
	}
	if len(extData) > 1 && extData[1] != "" {
		reqJson.Set("polyapitoken", extData[1])
	}
	if len(extData) > 2 && extData[2] != "" {
		reqJson.Set("shoptype", extData[2])
	} else {
		reqJson.Set("shoptype", "SOP")
	}
	bizcontent, resErr := reqJson.Encode()
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return hasNextPage, body, err
	}

	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return hasNextPage, body, err
	}
	// 通过奇门代理平台
	method := "Differ.JH.Other.DelegateQimenGetOrder"
	resJson := simplejson.New()
	resJson, body, err = client.Execute(method, params)
	if err != nil {
		fmt.Println(method, err)
		return hasNextPage, body, err
	}
	total, _ := resJson.Get("numtotalorder").Int()
	if pageIndex*pageSize < total {
		hasNextPage = true
	}
	return hasNextPage, body, err
}
