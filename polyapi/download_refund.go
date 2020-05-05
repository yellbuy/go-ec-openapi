package polyapi

import (
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

// 退货订单下载
func (client *Client) DownloadRefundList(pageIndex, pageSize int, startTime, endTime, timeType, status, orderToken string, extData ...string) (res []*common.OrderInfo, hasNextPage bool, nextToken string, body []byte, err error) {
	res = make([]*common.OrderInfo, 0)
	hasNextPage = false
	reqJson := simplejson.New()
	reqJson.Set("pageindex", pageIndex)
	reqJson.Set("pagesize", pageSize)
	reqJson.Set("starttime", startTime)
	reqJson.Set("endtime", endTime)
	reqJson.Set("timetype", timeType)
	reqJson.Set("status", status)
	reqJson.Set("refundtype", "JH_04")
	reqJson.Set("nexttoken", orderToken)
	reqJson.Set("randomnumber", orderToken)

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
		return res, hasNextPage, nextToken, body, err
	}

	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return res, hasNextPage, nextToken, body, err
	}
	// 通过奇门代理平台
	//method := "Differ.JH.Other.DelegateQimenGetOrder"
	// 通过polyapi自有平台
	method := "Differ.JH.Business.GetRefund"
	resJson := simplejson.New()
	resJson, body, err = client.Execute(method, params)
	if err != nil {
		fmt.Println(method, err)
		return res, hasNextPage, nextToken, body, err
	}
	nextToken, _ = resJson.Get("nexttoken").String()
	hasNextPageStr, _ := resJson.Get("ishasnextpage").String()
	hasNextPage = hasNextPageStr == "1"
	orderList, err := resJson.Get("refunds").Array()
	if err != nil {
		fmt.Println(method, err)
		return res, hasNextPage, nextToken, body, err
	}
	for index := range orderList {
		order := resJson.Get("refunds").GetIndex(index)
		orderInfo := new(common.OrderInfo)
		orderInfo.PlatOrderNo, _ = order.Get("platorderno").String()
		orderInfo.SubPlatOrderNo, _ = order.Get("subplatorderno").String()
		orderInfo.TradeStatus, _ = order.Get("refundstatus").String()
		orderInfo.Mobile, _ = order.Get("mobile").String()
		orderInfo.Phone, _ = order.Get("telephone").String()
		orderInfo.ReceiverName, _ = order.Get("receivername").String()
		orderInfo.LogisticName, _ = order.Get("logisticname").String()
		orderInfo.LogisticNo, _ = order.Get("logisticno").String()
		// orderInfo.Country, _ = order.Get("country").String()
		// orderInfo.Province, _ = order.Get("province").String()
		// orderInfo.City, _ = order.Get("city").String()
		// orderInfo.Area, _ = order.Get("area").String()
		// orderInfo.Town, _ = order.Get("town").String()
		orderInfo.Address, _ = order.Get("address").String()
		orderInfo.PayTime, _ = order.Get("createtime").String()
		orderInfo.TradeTime, _ = order.Get("createtime").String()

		goodsList, _ := order.Get("refundgoods").Array()
		orderInfo.GoodsInfoList = make([]*common.GoodsInfo, len(goodsList))
		for j := range goodsList {
			goodsJson := order.Get("refundgoods").GetIndex(j)
			goods := new(common.GoodsInfo)
			//goods.SubOrderNo, _ = goodsJson.Get("suborderno").String()
			goods.PlatGoodsId, _ = goodsJson.Get("platproductid").String()
			goods.PlatSkuId, _ = goodsJson.Get("sku").String()
			goods.OutItemId, _ = goodsJson.Get("outerid").String()
			goods.RefundStatus, _ = goodsJson.Get("refundstatus").String()
			goods.GoodsCount, _ = goodsJson.Get("productnum").String()
			goods.TradeGoodsName, _ = goodsJson.Get("productname").String()
			goods.TradeGoodsSpec, _ = order.Get("sku").String()
			//goods.Price, _ = goodsJson.Get("price").String()
			//goods.DiscountMoney, _ = goodsJson.Get("discountmoney").String()
			orderInfo.GoodsInfoList[j] = goods
		}
		res = append(res, orderInfo)
	}
	return res, hasNextPage, nextToken, body, err
}
