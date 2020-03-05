package polyapi

import (
	"encoding/json"
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

// 订单发货
func (client *Client) LogisticsSend(dto *common.LogisticsSendReqDto, extData ...string) ([]byte, error) {
	data, err := json.Marshal(dto)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	reqJson, err := simplejson.NewJson(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
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
		return nil, err
	}

	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	//fmt.Println("bizcontent：", string(bizcontent))
	params, err := common.InterfaceToParameter(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//fmt.Println("bizcontent2：", string(bizcontent))
	// 通过polyapi自有平台
	method := "Differ.JH.Business.Send"
	//return nil, errors.New("test")
	_, body, err := client.Execute(method, params)
	if err != nil {
		fmt.Println(method, err)
	}
	//fmt.Println("bizcontent3：", string(bizcontent))
	return body, err
}
