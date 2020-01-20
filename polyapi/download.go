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
	if len(extData) > 0 {
		reqJson.Set("platvalue", extData[0])
	}
	if len(extData) > 1 {
		reqJson.Set("polyapitoken", extData[1])
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
	if len(extData) > 0 {
		reqJson.Set("platvalue", extData[0])
	}
	if len(extData) > 1 {
		reqJson.Set("polyapitoken", extData[1])
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
	res, body, err = client.Execute(" Differ.JH.Other.DelegateQimenGetOrder", params)
	if err != nil {
		fmt.Println(" Differ.JH.Other.DelegateQimenGetOrder:", err)
		return res, body, err
	}
	return res, body, err
}

func NewSuccessResDto(isSuccess bool, code int, message, itemId string) *SuccessResDto {
	dto := new(SuccessResDto)
	dto.Response = new(successRes)
	if isSuccess {
		dto.Response.Flag = "success"
	} else {
		dto.Response.Flag = "failure"
	}
	dto.Response.Code = code
	dto.Response.Message = message
	dto.Response.ItemId = itemId
	return dto
}

func NewErrorResDto(code int, message string, subCode int, subMsg string) *ErrorResDto {
	dto := new(ErrorResDto)
	dto.Response = new(errorRes)
	dto.Response.Code = code
	dto.Response.Message = message
	dto.Response.SubCode = subCode
	dto.Response.SubMsg = subMsg
	return dto
}

type SuccessResDto struct {
	Response *successRes `json:"response"`
}

type ErrorResDto struct {
	Response *errorRes `json:"error_response"`
}

// 奇门下载成功响应内容
type successRes struct {
	//区名称（三级地址）
	Flag    string `json:"flag"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	ItemId  string `json:"itemId"`
}
type errorRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	SubCode int    `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}
