/*
文档地址：https://qimen.taobao.com/qimen/index.htm#/api-doc?fromPage=api-list&&apiServiceId=&officialApiId=26000&_k=mgwspb
*/

package qimen

import (
	"encoding/xml"
)

// 订单取消解析
func OrderCancelParse(body []byte) (res *OrderCancelReqDto, err error) {
	res = new(OrderCancelReqDto)
	err = xml.Unmarshal(body, res)
	return
}

type OrderCancelReqDto struct {
	XMLName       xml.Name          `xml:"request"`
	WarehouseCode string            `xml:"warehouseCode"`
	OwnerCode     string            `xml:"ownerCode"`
	OrderCode     string            `xml:"orderCode"`
	OrderId       string            `xml:"orderId"`
	OrderType     string            `xml:"orderType"`
	CancelReason  string            `xml:"cancelReason"`
	ExtendProps   map[string]string `xml:"extendProps"`
	remark        string            `xml:"remark"`
}
