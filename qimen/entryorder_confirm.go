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
	method := "qimen.entryorder.confirm"
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}

type EntryOrderConfirmReqDto struct {
	XMLName    xml.Name    `xml:"request"`
	Text       string      `xml:",chardata"`
	EntryOrder *EntryOrder `xml:"entryOrder"`
	OrderLines *OrderLines `xml:"orderLines"`
	Items      *Items      `xml:"items"`
}
type EntryOrder struct {
	SenderInfo struct {
		Text   string `xml:",chardata"`
		Remark string `xml:"remark"`
	} `xml:"senderInfo"`
	ReceiverInfo struct {
		Text   string `xml:",chardata"`
		Remark string `xml:"remark"`
	} `xml:"receiverInfo"`
	RelatedOrders struct {
		Text   string `xml:",chardata"`
		Remark string `xml:"remark"`
	} `xml:"relatedOrders"`
	OrderCode             string `xml:"orderCode"`
	OrderId               string `xml:"orderId"`
	OrderType             string `xml:"orderType"`
	WarehouseName         string `xml:"warehouseName"`
	LogisticsCode         string `xml:"logisticsCode"`
	LogisticsName         string `xml:"logisticsName"`
	TotalOrderLines       string `xml:"totalOrderLines"`
	EntryOrderCode        string `xml:"entryOrderCode"`
	OwnerCode             string `xml:"ownerCode"`
	PurchaseOrderCode     string `xml:"purchaseOrderCode"`
	WarehouseCode         string `xml:"warehouseCode"`
	EntryOrderId          string `xml:"entryOrderId"`
	EntryOrderType        string `xml:"entryOrderType"`
	OutBizCode            string `xml:"outBizCode"`
	ConfirmType           int    `xml:"confirmType"`
	Status                string `xml:"status"`
	OperateTime           string `xml:"operateTime"`
	Remark                string `xml:"remark"`
	SubOrderType          string `xml:"subOrderType"`
	ResponsibleDepartment string `xml:"responsibleDepartment"`
	ShopNick              string `xml:"shopNick"`
	ShopCode              string `xml:"shopCode"`
}
