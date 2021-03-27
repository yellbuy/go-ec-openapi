package taobao

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/yellbuy/go-ec-openapi/common"
)

func (client *Client) GetWaybill(request *common.WaybillApplyNewRequest, extData ...string) (*common.WaybillApplyNewCols, []byte, error) {
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

func (client *Client) GetWaybillTemplates(request *common.WaybillTemplateRequest, extData ...string) (res *common.WaybillTemplateDto, body []byte, err error) {
	err = errors.New("未实现")
	return
}
func (client *Client) CancelWaybill(request []common.WaybillCancel, extData ...string) (*common.WaybillCancelReturn, error) {
	err := errors.New("未实现")
	return nil, err
}
func (client *Client) DownloadOrderListV2(request common.DownLoadOrderListPostBizcontent, extData ...string) (common.DownloadOrderListReturn, error) {
	err := errors.New("未实现")
	var OutData common.DownloadOrderListReturn
	return OutData, err
}
