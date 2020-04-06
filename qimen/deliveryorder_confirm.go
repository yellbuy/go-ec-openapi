// https://qimen.taobao.com/qimen/index.htm#/api-doc?fromPage=api-list&&apiServiceId=&officialApiId=25995&_k=crdysi
// taobao.qimen.entryorder.confirm 入库单确认接口

package qimen

import (
	"encoding/xml"
	"fmt"

	"github.com/yellbuy/go-ec-openapi/common"
)

// 采购单回传
func (client *Client) DeliveryOrderConfirm(dto *DeliveryOrderConfirmReqDto) (body []byte, err error) {
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
	method := "qimen.deliveryorder.confirm"
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}

type DeliveryOrderConfirmReqDto struct {
	XMLName       xml.Name       `xml:"request"`
	Text          string         `xml:",chardata"`
	DeliveryOrder *DeliveryOrder `xml:"deliveryOrder"`
	OrderLines    *OrderLines    `xml:"orderLines"`
}
type DeliveryOrder struct {
	Text                 string `xml:",chardata"`
	DeliveryRequirements struct {
		Text   string `xml:",chardata"`
		Remark string `xml:"remark"`
	} `xml:"deliveryRequirements"`
	SenderInfo struct {
		Text   string `xml:",chardata"`
		Remark string `xml:"remark"`
	} `xml:"senderInfo"`
	ReceiverInfo struct {
		Text   string `xml:",chardata"`
		Remark string `xml:"remark"`
	} `xml:"receiverInfo"`
	PickerInfo struct {
		Text   string `xml:",chardata"`
		Remark string `xml:"remark"`
	} `xml:"pickerInfo"`
	OrderLines struct {
		Text   string `xml:",chardata"`
		Batchs struct {
			Text   string `xml:",chardata"`
			Remark string `xml:"remark"`
			SnCode string `xml:"snCode"`
		} `xml:"batchs"`
		SnCode string `xml:"snCode"`
		Remark string `xml:"remark"`
		SnList struct {
			Text string `xml:",chardata"`
			Sn   string `xml:"sn"`
		} `xml:"snList"`
	} `xml:"orderLines"`
	Items struct {
		Text   string `xml:",chardata"`
		Batchs struct {
			Text   string `xml:",chardata"`
			Remark string `xml:"remark"`
		} `xml:"batchs"`
		PriceAdjustment struct {
			Text   string `xml:",chardata"`
			Remark string `xml:"remark"`
		} `xml:"priceAdjustment"`
		Remark string `xml:"remark"`
	} `xml:"items"`
	Packages struct {
		Text                string `xml:",chardata"`
		PackageMaterialList struct {
			Text   string `xml:",chardata"`
			Remark string `xml:"remark"`
		} `xml:"packageMaterialList"`
		Items struct {
			Text   string `xml:",chardata"`
			Remark string `xml:"remark"`
		} `xml:"items"`
		Remark string `xml:"remark"`
	} `xml:"packages"`
	RelatedOrders struct {
		Text   string `xml:",chardata"`
		Remark string `xml:"remark"`
	} `xml:"relatedOrders"`
	DeliveryOrderCode string `xml:"deliveryOrderCode"`
	DeliveryOrderId   string `xml:"deliveryOrderId"`
	WarehouseCode     string `xml:"warehouseCode"`
	OrderType         string `xml:"orderType"`
	Status            string `xml:"status"`
	LogisticsCode     string `xml:"logisticsCode"`
	LogisticsName     string `xml:"logisticsName"`
	ExpressCode       string `xml:"expressCode"`
	OutBizCode        string `xml:"outBizCode"`
	ConfirmType       string `xml:"confirmType"`
	OrderConfirmTime  string `xml:"orderConfirmTime"`
	OperatorCode      string `xml:"operatorCode"`
	OperatorName      string `xml:"operatorName"`
	OperateTime       string `xml:"operateTime"`
	StorageFee        string `xml:"storageFee"`
	Invoices          struct {
		Text    string `xml:",chardata"`
		Remark  string `xml:"remark"`
		Header  string `xml:"header"`
		Amount  string `xml:"amount"`
		Content string `xml:"content"`
		Detail  struct {
			Text  string `xml:",chardata"`
			Items struct {
				Text   string `xml:",chardata"`
				Batchs struct {
					Text   string `xml:",chardata"`
					Remark string `xml:"remark"`
				} `xml:"batchs"`
				PriceAdjustment struct {
					Text   string `xml:",chardata"`
					Remark string `xml:"remark"`
				} `xml:"priceAdjustment"`
				Remark   string `xml:"remark"`
				ItemName string `xml:"itemName"`
				Unit     string `xml:"unit"`
				Price    string `xml:"price"`
				Quantity string `xml:"quantity"`
				Amount   string `xml:"amount"`
				ItemCode string `xml:"itemCode"`
				ItemId   string `xml:"itemId"`
			} `xml:"items"`
		} `xml:"detail"`
		Code   string `xml:"code"`
		Number string `xml:"number"`
	} `xml:"invoices"`
}
