package ecopenapi

import (
	"fmt"
	"testing"

	"github.com/yellbuy/go-ec-openapi/common"
)

func TestExecute(t *testing.T) {
	// http://open.taobao.com/docs/api.htm?apiId=24515
	client, err := NewClient(TB, &common.ClientParams{"23268761", "b17cc059ffba3f6cf0d9131359d0be2a", "62024115ddfe90c8c9cege1fefda96512336cddb9fe0f852523122485", ""})
	if err != nil {
		t.Fatal(err)
	}
	request := new(common.WaybillApplyNewRequest)
	request.ShippingAddress = new(common.WaybillAddress)
	// 订单数据，必填
	request.TradeOrderInfoCols = make([]*common.TradeOrderInfo, 0)
	res, _, err := client.GetWaybill(request)
	fmt.Println(res)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDownloadProductExecute(t *testing.T) {
	// 752代表走奇门委托
	platId := fmt.Sprintf("%v", 2)
	platformType := POLYAPI
	client, err := NewClient(platformType, &common.ClientParams{"8e770a60b9684c558f40e4796a96710f", "c9cb1df531b441a8872c60ffb7f900a6", "f50d2f8b2cdf4ad8a5b6eb25bc58e4df", platId})
	if err == nil {
		res, _, body, err := client.DownloadProductList(0, 10, "JH_01")
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(string(body))
		fmt.Println(res)
	}

}

func TestDownloadOrderExecute(t *testing.T) {
	platId := fmt.Sprintf("%v", 2)
	platformType := POLYAPI
	client, err := NewClient(platformType, &common.ClientParams{"8e770a60b9684c558f40e4796a96710f", "c9cb1df531b441a8872c60ffb7f900a6", "f50d2f8b2cdf4ad8a5b6eb25bc58e4df", platId})
	if err == nil {
		_, _, _, body, err := client.DownloadOrderList(0, 100, "2020-01-20 00:00:00", "2020-01-22 00:00:00", "JH_03", "JH_02", "")
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(string(body))
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
