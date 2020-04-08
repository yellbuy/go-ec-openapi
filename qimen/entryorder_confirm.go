// https://qimen.taobao.com/qimen/index.htm#/api-doc?fromPage=api-list&&apiServiceId=&officialApiId=25995&_k=crdysi
// taobao.qimen.entryorder.confirm 入库单确认接口

package qimen

import (
	"encoding/xml"
	"fmt"

	"github.com/yellbuy/go-ec-openapi/common"
)

// 采购单回传
func (client *Client) EntryOrderConfirm(dto *EntryOrderConfirmReqDto) (body []byte, err error) {
	var bytes []byte
	bytes, err = xml.Marshal(dto)

	req := make(map[string]interface{})
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		body = nil
		return
	}

	// 通过奇门代理平台
	method := "taobao.qimen.entryorder.confirm"
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}

type EntryOrderConfirmReqDto struct {
	XMLName    xml.Name        `xml:"request"`
	Text       string          `xml:",chardata"`
	EntryOrder EntryOrder      `xml:"entryOrder"`
	OrderLines EntryOrderLines `xml:"orderLines"`
}
type EntryOrder struct {
	Text            string `xml:",chardata"`
	TotalOrderLines string `xml:"totalOrderLines"`
	EntryOrderCode  string `xml:"entryOrderCode"`
	OwnerCode       string `xml:"ownerCode"`
	WarehouseCode   string `xml:"warehouseCode"`
	EntryOrderId    string `xml:"entryOrderId"`
	EntryOrderType  string `xml:"entryOrderType"`
	OutBizCode      string `xml:"outBizCode"`
	ConfirmType     int    `xml:"confirmType"`
	Status          string `xml:"status"`
	Freight         string `xml:"freight"`
	OperateTime     string `xml:"operateTime"`
	Remark          string `xml:"remark"`
}
type EntryOrderLines struct {
	Text          string           `xml:",chardata"`
	OrderLineList []EntryOrderLine `xml:"orderLine"`
}

type EntryOrderLine struct {
	Text        string `xml:",chardata"`
	OutBizCode  string `xml:"outBizCode"`
	OrderLineNo string `xml:"orderLineNo"`
	OwnerCode   string `xml:"ownerCode"`
	ItemCode    string `xml:"itemCode"`
	ItemId      string `xml:"itemId"`
	SnList      struct {
		Text string `xml:",chardata"`
		Sn   string `xml:"sn"`
	} `xml:"snList"`
	ItemName      string `xml:"itemName"`
	InventoryType string `xml:"inventoryType"`
	PlanQty       int    `xml:"planQty"`
	ActualQty     int    `xml:"actualQty"`
	BatchCode     string `xml:"batchCode"`
	ProductDate   string `xml:"productDate"`
	ExpireDate    string `xml:"expireDate"`
	ProduceCode   string `xml:"produceCode"`
	Batchs        struct {
		Text  string `xml:",chardata"`
		Batch struct {
			Text          string `xml:",chardata"`
			BatchCode     string `xml:"batchCode"`
			ProductDate   string `xml:"productDate"`
			ExpireDate    string `xml:"expireDate"`
			ProduceCode   string `xml:"produceCode"`
			InventoryType string `xml:"inventoryType"`
			ActualQty     string `xml:"actualQty"`
		} `xml:"batch"`
	} `xml:"batchs"`
	Remark string `xml:"remark"`
}
