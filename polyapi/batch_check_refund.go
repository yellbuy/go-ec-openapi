package polyapi

import (
	"encoding/json"
	"errors"
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

//批量检测退货
func (client *Client) BatchCheckRefundStatus(platOrderNoList []string, extData ...string) ([]*common.BatchCheckRefundStatusRes, []byte, error) {
	if len(platOrderNoList) == 0 {
		return nil, nil, errors.New("单号不能为空")
	}
	dto := new(common.BatchCheckRefundStatusReq)
	dto.Orders = make([]*common.BatchCheckRefundStatusOrder, len(platOrderNoList))
	for index, no := range platOrderNoList {
		item := new(common.BatchCheckRefundStatusOrder)
		item.PlatOrderNo = no
		item.ShopType = "JH_001"
		item.CountryCode = "CN"
		dto.Orders[index] = item
	}
	data, err := json.Marshal(dto)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	reqJson, err := simplejson.NewJson(data)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
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

	bizcontent, err := reqJson.Encode()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	//fmt.Println("bizcontent：", string(bizcontent))
	params, err := common.InterfaceToParameter(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	//fmt.Println("bizcontent2：", string(bizcontent))
	// 通过polyapi自有平台
	method := "Differ.JH.Business.BatchCheckRefundStatus"
	//return nil, errors.New("test")
	resJson, body, err := client.Execute(method, params)
	if err != nil {
		fmt.Println(method, err)
		return nil, body, err
	}
	results, err := resJson.Get("results").Array()
	if err != nil {
		fmt.Println(method, err, string(body))
		return nil, body, err
	}
	res := make([]*common.BatchCheckRefundStatusRes, len(results))
	for index, _ := range results {
		val := resJson.Get("results").GetIndex(index)
		model := new(common.BatchCheckRefundStatusRes)
		model.PlatOrderNo, _ = val.Get("platorderno").String()
		model.RefundStatus, _ = val.Get("refundstatus").String()
		model.RefundStatusDescription, _ = val.Get("refundstatusdescription").String()
		model.TradeStatus, _ = val.Get("tradestatus").String()
		children, _ := val.Get("childrenrefundstatus").Array()
		model.ChildrenRefundStatusList = make([]*common.ChildrenRefundStatus, len(children))
		for childIndex, _ := range children {
			childJson := val.Get("ChildrenRefundStatus").GetIndex(childIndex)
			child := new(common.ChildrenRefundStatus)
			child.RefundNo, _ = childJson.Get("refundno").String()
			child.SubOrderNo, _ = childJson.Get("suborderno").String()
			child.ProductName, _ = childJson.Get("productname").String()
			child.PlatProductId, _ = childJson.Get("platproductid").String()
			child.TradeGoodsNo, _ = childJson.Get("tradegoodsno").String()
			child.RefundStatus, _ = childJson.Get("refundstatus").String()
			child.RefundStatusDescription, _ = childJson.Get("refundstatusdescription").String()
			model.ChildrenRefundStatusList[childIndex] = child
		}
		res[index] = model
	}

	//fmt.Println("bizcontent3：", string(bizcontent))
	return res, body, err
}
