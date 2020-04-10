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
	method := "taobao.qimen.deliveryorder.confirm"
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}

type DeliveryOrderConfirmReqDto struct {
	XMLName       xml.Name           `xml:"request"`
	Text          string             `xml:",chardata"`
	DeliveryOrder DeliveryOrder      `xml:"deliveryOrder"`
	Packages      DeliveryPackages   `xml:"packages"`
	OrderLines    DeliveryOrderLines `xml:"orderLines"`
}

type DeliveryPackages struct {
	Text        string            `xml:",chardata"`
	PackageList []DeliveryPackage `xml:"package"`
}
type DeliveryPackage struct {
	Text                string `xml:",chardata"`
	LogisticsCode       string `xml:"logisticsCode"`
	LogisticsName       string `xml:"logisticsName"`
	ExpressCode         string `xml:"expressCode"`
	PackageCode         string `xml:"packageCode"`
	Length              string `xml:"length"`
	Width               string `xml:"width"`
	Height              string `xml:"height"`
	TheoreticalWeight   string `xml:"theoreticalWeight"`
	Weight              string `xml:"weight"`
	Volume              string `xml:"volume"`
	InvoiceNo           string `xml:"invoiceNo"`
	PackageMaterialList struct {
		Text            string `xml:",chardata"`
		PackageMaterial struct {
			Text     string `xml:",chardata"`
			Type     string `xml:"type"`
			Quantity string `xml:"quantity"`
		} `xml:"packageMaterial"`
	} `xml:"packageMaterialList"`
	Items DeliveryItems `xml:"items"`
}
type DeliveryItems struct {
	Text     string         `xml:",chardata"`
	ItemList []DeliveryItem `xml:"item"`
}
type DeliveryItem struct {
	Text     string `xml:",chardata"`
	ItemCode string `xml:"itemCode"`
	ItemId   string `xml:"itemId"`
	Quantity int    `xml:"quantity"`
}

type DeliveryOrder struct {
	Text              string `xml:",chardata"`
	DeliveryOrderCode string `xml:"deliveryOrderCode"`
	DeliveryOrderId   string `xml:"deliveryOrderId"`
	WarehouseCode     string `xml:"warehouseCode"`
	OrderType         string `xml:"orderType"`
	Status            string `xml:"status"`
	OutBizCode        string `xml:"outBizCode"`
	ConfirmType       string `xml:"confirmType"`
	OrderConfirmTime  string `xml:"orderConfirmTime"`
	OperatorCode      string `xml:"operatorCode"`
	OperatorName      string `xml:"operatorName"`
	OperateTime       string `xml:"operateTime"`
	StorageFee        string `xml:"storageFee"`
	Invoices          struct {
		Text    string `xml:",chardata"`
		Invoice struct {
			Text    string `xml:",chardata"`
			Header  string `xml:"header"`
			Amount  string `xml:"amount"`
			Content string `xml:"content"`
			Detail  struct {
				Text  string `xml:",chardata"`
				Items struct {
					Text string `xml:",chardata"`
					Item struct {
						Text     string `xml:",chardata"`
						ItemName string `xml:"itemName"`
						Unit     string `xml:"unit"`
						Price    string `xml:"price"`
						Quantity string `xml:"quantity"`
						Amount   string `xml:"amount"`
					} `xml:"item"`
				} `xml:"items"`
			} `xml:"detail"`
			Code   string `xml:"code"`
			Number string `xml:"number"`
		} `xml:"invoice"`
	} `xml:"invoices"`
}
type DeliveryOrderLines struct {
	Text          string              `xml:",chardata"`
	OrderLineList []DeliveryOrderLine `xml:"orderLine"`
}
type DeliveryOrderLine struct {
	Text            string   `xml:",chardata"`
	OrderLineNo     string   `xml:"orderLineNo"`
	OrderSourceCode string   `xml:"orderSourceCode"`
	SubSourceCode   string   `xml:"subSourceCode"`
	OwnerCode       []string `xml:"ownerCode"`
	ItemCode        string   `xml:"itemCode"`
	ItemId          string   `xml:"itemId"`
	SnList          struct {
		Text string `xml:",chardata"`
		Sn   string `xml:"sn"`
	} `xml:"snList"`
	InventoryType  string  `xml:"inventoryType"`
	ItemName       string  `xml:"itemName"`
	ExtCode        string  `xml:"extCode"`
	PlanQty        int     `xml:"planQty"`
	ActualQty      int     `xml:"actualQty"`
	RetailPrice    float32 `xml:"retailPrice"`
	ActualPrice    float32 `xml:"actualPrice"`
	DiscountAmount float32 `xml:"discountAmount"`
	BatchCode      string  `xml:"batchCode"`
	ProductDate    string  `xml:"productDate"`
	ExpireDate     string  `xml:"expireDate"`
	ProduceCode    string  `xml:"produceCode"`
	Batchs         struct {
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
	QrCode string `xml:"qrCode"`
}
