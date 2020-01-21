package polyapi

import (
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

// 产品下载
func (client *Client) DownloadProductList(pageIndex, pageSize int, status string, extData ...string) (res []*common.Product, hasNextPage bool, body []byte, err error) {
	res = make([]*common.Product, 0)
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
	if len(extData) > 2 {
		reqJson.Set("shoptype", extData[2])
	} else {
		reqJson.Set("shoptype", "JH_001")
	}

	bizcontent, resErr := reqJson.Encode()
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return res, hasNextPage, body, err
	}

	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return res, hasNextPage, body, err
	}

	// 通过奇门代理平台
	//method := "Differ.JH.Other.DelegateQimenDownloadProduct"
	// 通过polyapi自有平台
	method := "Differ.JH.Business.DownloadProduct"
	resJson := simplejson.New()
	resJson, body, err = client.Execute("Differ.JH.Business.DownloadProduct", params)
	if err != nil {
		fmt.Println(method, err)
		return res, hasNextPage, body, err
	}
	hasNextPageStr, _ := resJson.Get("ishasnextpage").String()
	hasNextPage = hasNextPageStr == "1"
	goodsList, err := resJson.Get("goodslist").Array()
	if err != nil {
		fmt.Println(method, err)
		return res, hasNextPage, body, err
	}
	for index := range goodsList {
		goods := resJson.Get("goodslist").GetIndex(index)
		product := new(common.Product)
		product.ProductId, _ = goods.Get("platproductid").String()
		product.ProductName, _ = goods.Get("name").String()
		product.ProductCode, _ = goods.Get("outerid").String()
		product.Price, _ = goods.Get("price").String()
		product.Num, _ = goods.Get("num").String()
		product.WhseCode, _ = goods.Get("whsecode").String()
		product.Attrbutes, _ = goods.Get("attrbutes").String()
		product.CategoryId, _ = goods.Get("categoryid").String()
		product.Status, _ = goods.Get("status").String()
		product.PropertyAlias, _ = goods.Get("propertyalias").String()

		skuList, _ := goods.Get("skus").Array()
		product.SkuList = make([]*common.Sku, len(skuList))
		for j := range skuList {
			skuJson := goods.Get("skus").GetIndex(j)
			sku := new(common.Sku)
			sku.SkuId, _ = skuJson.Get("skuid").String()
			sku.SkuCode, _ = skuJson.Get("skuouterid").String()
			sku.SkuPrice, _ = skuJson.Get("skuprice").String()
			sku.SkuQuantity, _ = skuJson.Get("skuquantity").String()
			sku.SkuName, _ = skuJson.Get("skuname").String()
			sku.SkuProperty, _ = skuJson.Get("skuproperty").String()
			sku.SkuPictureUrl, _ = skuJson.Get("skupictureurl").String()
			sku.SkuName2, _ = skuJson.Get("skuname2").String()
			product.SkuList[j] = sku
		}
		res = append(res, product)
	}
	return res, hasNextPage, body, err
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
	reqJson.Set("shoptype", "JH_001")
	if len(extData) > 0 {
		reqJson.Set("platvalue", extData[0])
	}
	if len(extData) > 1 {
		reqJson.Set("polyapitoken", extData[1])
	}
	if len(extData) > 2 {
		reqJson.Set("shoptype", extData[2])
	} else {
		reqJson.Set("shoptype", "JH_001")
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
	// 通过奇门代理平台
	//method := "Differ.JH.Other.DelegateQimenGetOrder"
	// 通过polyapi自有平台
	method := "Differ.JH.Business.GetOrder"
	res, body, err = client.Execute(method, params)
	if err != nil {
		fmt.Println(method, err)
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

// Polyapi接口商品下载响应
type GoodsDownloadResponseDto struct {
	Ishasnextpage    string  `json:"ishasnextpage"`
	Totalcount       string  `json:"totalcount"`
	Goodslist        []Goods `json:"goodslist"`
	Requestid        string  `json:"requestid"`
	Code             string  `json:"code"`
	Msg              string  `json:"msg"`
	Subcode          string  `json:"subcode"`
	Submessage       string  `json:"submessage"`
	Polyapitotalms   string  `json:"polyapitotalms"`
	Polyapirequestid string  `json:"polyapirequestid"`
}
type Sku struct {
	Skuid         string `json:"skuid"`
	Skuouterid    string `json:"skuouterid"`
	Skuprice      string `json:"skuprice"`
	Skuquantity   string `json:"skuquantity"`
	Skuname       string `json:"skuname"`
	Skuproperty   string `json:"skuproperty"`
	Skutype       string `json:"skutype"`
	Skupictureurl string `json:"skupictureurl"`
	Skuname2      string `json:"skuname2"`
}
type Goods struct {
	Platproductid      string      `json:"platproductid"`
	Name               string      `json:"name"`
	Outerid            string      `json:"outerid"`
	Price              string      `json:"price"`
	Num                string      `json:"num"`
	Pictureurl         string      `json:"pictureurl"`
	Whsecode           string      `json:"whsecode"`
	Attrbutes          interface{} `json:"attrbutes"`
	Categoryid         string      `json:"categoryid"`
	Status             string      `json:"status"`
	Statusdesc         string      `json:"statusdesc"`
	SkuList            []Sku       `json:"skus"`
	Sendtype           string      `json:"sendtype"`
	Skutype            string      `json:"skutype"`
	Propertyalias      string      `json:"propertyalias"`
	Isplatstorageorder string      `json:"isplatstorageorder"`
	Cooperationno      string      `json:"cooperationno"`
}

// 奇门接口成功响应
type SuccessResDto struct {
	Response *successRes `json:"response"`
}

// 奇门接口错误响应
type ErrorResDto struct {
	Response *errorRes `json:"error_response"`
}

// 奇门接口下载成功响应内容
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
