package ecopenapi

import (
	"fmt"
	"testing"

	"github.com/yellbuy/go-ec-openapi/common"
)

func TestExecute(t *testing.T) {
	// http://open.taobao.com/docs/api.htm?apiId=24515
	client, err := NewClient(TAOBAO, &common.ClientParams{"23268761", "b17cc059ffba3f6cf0d9131359d0be2a", "62026037725650ee7e9056985050e24865e52523122485"})
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.Execute("taobao.wlb.waybill.i.get", common.Parameter{
		"cp_code": "ZTO",
		"shipping_address":       "女装",
		"cat":     "16,18",
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("商品数量:", res.Get("tbk_item_get_response").Get("total_results").MustInt())
	var imtes []interface{}
	imtes, _ = res.Get("tbk_item_get_response").Get("results").Get("n_tbk_item").Array()
	for _, v := range imtes {
		fmt.Println("======")
		item := v.(map[string]interface{})
		fmt.Println("商品名称:", item["title"])
		fmt.Println("商品价格:", item["reserve_price"])
		fmt.Println("商品链接:", item["item_url"])
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
