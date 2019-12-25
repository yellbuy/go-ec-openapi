package ecopenapi

import (
	"fmt"
	"testing"

	"github.com/yellbuy/go-ec-openapi/common"
)

func TestExecute(t *testing.T) {
	// http://open.taobao.com/docs/api.htm?apiId=24515
	client, err := NewClient(TAOBAO, &common.ClientParams{"23268761", "b17cc059ffba3f6cf0d9131359d0be2a", "62024115ddfe90c8c9cege1fefda96512336cddb9fe0f852523122485"})
	if err != nil {
		t.Fatal(err)
	}
	request := new(common.WaybillApplyNewRequest)
	request.ShippingAddress = new(common.WaybillAddress)
	// 订单数据，必填
	request.TradeOrderInfoCols = make([]*common.TradeOrderInfo, 0)
	res, err := client.GetWaybill(request)
	fmt.Println(res)
	if err != nil {
		t.Fatal(err)
	}
}

// func TestExecuteErrMsg(t *testing.T) {
// 	_, err := Execute("taobao.tbk.dg.material.optional", ecopenapi.Parameter{
// 		"q":   "女装",
// 		"cat": "16,18",
// 	})

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestExecuteCache(t *testing.T) {
// 	res, err := ExecuteCache("taobao.tbk.item.get", ecopenapi.Parameter{
// 		"fields": "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick",
// 		"q":      "女装",
// 		"cat":    "16,18",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v\n", res)

// }
