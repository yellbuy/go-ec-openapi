/*
文档地址：https://qimen.taobao.com/qimen/index.htm#/api-doc?fromPage=api-list&&apiServiceId=&officialApiId=26002&_k=xignfc
*/

package qimen

import (
	"encoding/xml"
)

// 菠萝派销售单下载同步解析
func DeliveryOrderCreatePolyApiParse(body []byte) (res *DeliveryOrderPolyApiCreateReqDto, err error) {
	res = new(DeliveryOrderPolyApiCreateReqDto)
	err = xml.Unmarshal(body, res)
	return
}

type DeliveryOrderPolyApiCreateReqDto struct {
	XMLName     xml.Name                    `xml:"request"`
	Orders      *DeliveryOrderPolyApiOrders `xml:"orders"`
	ExtendProps *ExtendProps                `xml:"extendProps"`
}
type DeliveryOrderPolyApiOrders struct {
	Text      string                      `xml:",chardata"`
	OrderList []DeliveryOrderPolyApiOrder `xml:"order"`
}
type DeliveryOrderPolyApiOrder struct {
	Text          string                  `xml:",chardata"`
	DeliveryOrder *DeliveryOrderCreateDto `xml:"deliveryOrder"`
	OrderLines    *OrderLines             `xml:"orderLines"`
}
