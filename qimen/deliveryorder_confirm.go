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

// 奇门订单流水回传信息
func (client *Client) QimenOrderProcess(dto *common.QimenOrderProcessReportRequest) (body []byte, err error) {
	var bytes []byte
	bytes, err = xml.Marshal(dto)
	if err != nil {
		return nil, err
	}
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
	method := "taobao.qimen.orderprocess.report"
	// fmt.Println("奇门订单流水信息推送报文", string(bytes))
	// nowbody, _ := json.Marshal(client)
	// // fmt.Println("请求参数", string(nowbody))
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}

// 奇门出库单回传
func (client *Client) QimenStockoutConfirm(dto *WmsQimenStockoutConfirm) (body []byte, err error) {
	var bytes []byte
	bytes, err = xml.Marshal(dto)
	if err != nil {
		return nil, err
	}
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
	method := "taobao.qimen.stockout.confirm"
	fmt.Println("奇门出库单推送报文", string(bytes))
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}

// 奇门淘宝订单解密
func (client *Client) QimenTBDecrypt(dto *WmsQimenTBDecrypt) (*WmsQimenTBDecryptReturn, error) {
	var bytes []byte
	bytes, err := xml.Marshal(dto)

	req := make(map[string]interface{})
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr

		return nil, err
	}

	// 通过奇门代理平台
	method := "taobao.qimen.receiverinfo.query"
	body, err := client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	outData := new(WmsQimenTBDecryptReturn)
	err = xml.Unmarshal(body, outData)
	if err != nil {
		return nil, err
	}
	return outData, nil
}

type WmsQimenTBDecrypt struct {
	XMLName           xml.Name `xml:"request"`
	Oaid              string   `json:"oaid" xml:"oaid"`                           //false	订单收件人 ID, string (50)	订单收件人 ID, string (50)
	DeliveryOrderCode string   `json:"deliveryOrderCode" xml:"deliveryOrderCode"` //true	出库单号, string (50) , 必填	出库单号, string (50) , 必填
	OwnerCode         string   `json:"ownerCode" xml:"ownerCode"`                 //false	货主ID, string (50)	货主ID
	WarehouseCode     string   `json:"warehouseCode" xml:"warehouseCode"`         //false	WS1231，string (64)	仓库编码
	Scene             string   `json:"scene" xml:"scene"`                         //false	使用场景, string(10)，必填	使用场景。1001，顺丰电子面单发货；1002，4通一达电子面单发货；1003，EMS电子面单发货；1004，其他电子面单发货；2001，客户售后服务
}
type WmsQimenTBDecryptReturn struct {
	Flag              string                               `json:"flag" xml:"flag"`                           //success|failure，必填	success|failure，必填
	Code              string                               `json:"code" xml:"code"`                           //响应码	响应码
	Message           string                               `json:"message" xml:"message"`                     //响应信息	响应信息
	Oaid              string                               `json:"oaid" xml:"oaid"`                           //订单收件人 ID, string (50)	订单收件人 ID, string (50)
	DeliveryOrderCode string                               `json:"deliveryOrderCode" xml:"deliveryOrderCode"` //出库单号, string (50) , 必填	出库单号, string (50) , 必填
	ReceiverInfo      *WmsQimenTBDecryptReturnReceiverInfo `json:"receiverInfo" xml:"receiverInfo"`           //收货人信息
}
type WmsQimenTBDecryptReturnReceiverInfo struct {
	Oaid          string `json:"oaid" xml:"oaid"`                   //订单收件人 ID, string (50)	订单收件人 ID, string (50)
	Name          string `json:"name" xml:"name"`                   //姓名, string (50) , 必填	姓名, string (50) , 必填
	Tel           string `json:"tel" xml:"tel"`                     //固定电话, string (50)	固定电话, string (50)
	Mobile        string `json:"mobile" xml:"mobile"`               //移动电话, string (50) , 必填	移动电话, string (50) , 必填
	CountryCode   string `json:"countryCode" xml:"countryCode"`     //国家二字码，string(50)	国家二字码，string(50)
	Province      string `json:"province" xml:"province"`           //省份, string (50) , 必填	省份, string (50) , 必填
	DetailAddress string `json:"detailAddress" xml:"detailAddress"` //详细地址, string (200) , 必填	详细地址, string (200) , 必填
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
	ConfirmType       int    `xml:"confirmType"`
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
	Text            string `xml:",chardata"`
	OrderLineNo     string `xml:"orderLineNo"`
	OrderSourceCode string `xml:"orderSourceCode"`
	SubSourceCode   string `xml:"subSourceCode"`
	OwnerCode       string `xml:"ownerCode"`
	ItemCode        string `xml:"itemCode"`
	ItemId          string `xml:"itemId"`
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
	// Batchs         struct {
	// 	Text  string `xml:",chardata"`
	// 	Batch struct {
	// 		Text          string `xml:",chardata"`
	// 		BatchCode     string `xml:"batchCode"`
	// 		ProductDate   string `xml:"productDate"`
	// 		ExpireDate    string `xml:"expireDate"`
	// 		ProduceCode   string `xml:"produceCode"`
	// 		InventoryType string `xml:"inventoryType"`
	// 		ActualQty     string `xml:"actualQty"`
	// 	} `xml:"batch"`
	// } `xml:"batchs"`
	QrCode string `xml:"qrCode"`
}
