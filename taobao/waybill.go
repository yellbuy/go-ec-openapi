package taobao

import (
	"encoding/json"
	"fmt"

	"github.com/yellbuy/go-ec-openapi/common"
)

func (client *Client) GetWaybill(request *common.WaybillApplyNewRequest) (*common.WaybillApplyNewCols, []byte, error) {
	for index, _ := range request.TradeOrderInfoCols {
		request.TradeOrderInfoCols[index].OrderChannelsType = "TB"
	}
	req := make(map[string]interface{})
	req["waybill_apply_new_request"] = request
	params, err := common.InterfaceToParameter(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	res, err := client.Execute("taobao.wlb.waybill.i.get", params)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	data, err := res.Get("wlb_waybill_i_get_response").Get("waybill_apply_new_cols").Encode()
	if err != nil {
		fmt.Println(err)
		return nil, data, err
	}
	// fmt.Println("waybill_apply_new_cols:", string(data))
	result := new(common.WaybillApplyNewCols)
	err = json.Unmarshal(data, result)
	if err != nil {
		fmt.Println(err)
	}
	return result, data, err
}
