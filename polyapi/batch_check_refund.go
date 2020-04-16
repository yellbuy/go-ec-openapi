package polyapi

import (
	"encoding/json"
	"errors"
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

//批量检测退货
func (client *Client) BatchCheckRefundStatus(platOrderNoList []string, extData ...string) (res []*common.BatchCheckRefundStatusRes, body []byte, err error) {
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
	_, body, err = client.Execute(method, params)
	if err != nil {
		fmt.Println(method, err)
	}
	//fmt.Println("bizcontent3：", string(bizcontent))
	return nil, body, err
}
