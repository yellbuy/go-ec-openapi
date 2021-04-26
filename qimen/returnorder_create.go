/*
文档地址：https://qimen.taobao.com/qimen/index.htm#/api-doc?fromPage=api-list&&apiServiceId=&officialApiId=25991&_k=2c1orl
*/

package qimen

import (
	"encoding/xml"
)

// 退货单下载同步解析
func ReturnOrderCreateParse(body []byte) (res *ReturnOrderCreateReqDto, err error) {
	res = new(ReturnOrderCreateReqDto)
	err = xml.Unmarshal(body, res)
	return
}

type ReturnOrderCreateReqDto struct {
	XMLName     xml.Name        `xml:"request"`
	ReturnOrder *ReturnOrderDto `xml:"returnOrder"`
	OrderLines  *OrderLines     `xml:"orderLines"`
}
type ReturnOrderDto struct {
	Text            string `xml:",chardata"`
	ReturnOrderCode string `xml:"returnOrderCode"`
	ReturnOrderId   string `xml:"returnOrderId"`
	OutBizCode      string `xml:"outBizCode"`
	OwnerCode       string `xml:"ownerCode"`
	WarehouseCode   string `xml:"warehouseCode"`
	OrderType       string `xml:"orderType"`
	ReturnReason    string `xml:"returnReason"`
	LogisticsCode   string `xml:"logisticsCode"`
	LogisticsName   string `xml:"logisticsName"`
	ExpressCode     string `xml:"expressCode"`
	SenderInfo      struct {
		Text          string `xml:",chardata"`
		Company       string `xml:"company"`
		Name          string `xml:"name"`
		ZipCode       string `xml:"zipCode"`
		Tel           string `xml:"tel"`
		Mobile        string `xml:"mobile"`
		Email         string `xml:"email"`
		CountryCode   string `xml:"countryCode"`
		Province      string `xml:"province"`
		City          string `xml:"city"`
		Area          string `xml:"area"`
		Town          string `xml:"town"`
		DetailAddress string `xml:"detailAddress"`
	} `xml:"senderInfo,omitempty"`
	Remark string `xml:"remark"`
}

type ReturnOrderLines struct {
	OrderLineList []ReturnOrderLine `xml:"orderLine"`
}
type ReturnOrderLine struct {
	OrderLineNo        string `xml:"orderLineNo"`
	SourceOrderCode    string `xml:"sourceOrderCode"`
	SubSourceOrderCode string `xml:"subSourceOrderCode"`
	OwnerCode          string `xml:"ownerCode"`
	ItemCode           string `xml:"itemCode"`
	ItemId             string `xml:"itemId"`
	SnList             struct {
		Text string `xml:",chardata"`
		Sn   string `xml:"sn"`
	} `xml:"snList"`
	InventoryType string       `xml:"inventoryType"`
	PlanQty       int          `xml:"planQty"`
	ActualQty     int          `xml:"actualQty"`
	BatchCode     string       `xml:"batchCode"`
	ProductDate   string       `xml:"productDate"`
	ExpireDate    string       `xml:"expireDate"`
	ProduceCode   string       `xml:"produceCode"`
	Batchs        ReturnBatchs `xml:"batchs,omitempty"`
	QrCode        string       `xml:"qrCode"`
}
type ReturnBatchs struct {
	BatchList []ReturnBatch `xml:"batch"`
}
type ReturnBatch struct {
	Text          string `xml:",chardata"`
	BatchCode     string `xml:"batchCode"`
	ProductDate   string `xml:"productDate"`
	ExpireDate    string `xml:"expireDate"`
	ProduceCode   string `xml:"produceCode"`
	InventoryType string `xml:"inventoryType"`
	ActualQty     int    `xml:"actualQty"`
}
type ReturnOrderCreateResponse struct {
	XMLName xml.Name `xml:"response"`
	Response
	ReturnOrderId string `xml:"returnOrderId,omitempty"`
}

func NewReturnOrderCreateSuccessResponse(returnOrderId string) *ReturnOrderCreateResponse {
	dto := new(ReturnOrderCreateResponse)
	dto.Flag = "success"
	dto.Code = "0"
	dto.ReturnOrderId = returnOrderId
	return dto
}
