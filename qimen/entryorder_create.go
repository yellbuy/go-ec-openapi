/*
文档地址：https://qimen.taobao.com/qimen/index.htm#/api-doc?fromPage=api-list&&apiServiceId=&officialApiId=25991&_k=2c1orl
*/

package qimen

import (
	"encoding/xml"
)

// 入库单下载同步解析
func EntryOrderCreateParse(body []byte) (res *EntryOrderCreateReqDto, err error) {
	res = new(EntryOrderCreateReqDto)
	err = xml.Unmarshal(body, res)
	return
}

type EntryOrderCreateReqDto struct {
	XMLName    xml.Name             `xml:"request"`
	EntryOrder *EntryOrderCreateDto `xml:"entryOrder"`
	OrderLines *OrderLines          `xml:"orderLines"`
	Items      *Items               `xml:"items"`
}
type EntryOrderCreateDto struct {
	Text              string `xml:",chardata"`
	EntryOrderCode    string `xml:"entryOrderCode"`
	OwnerCode         string `xml:"ownerCode"`
	PurchaseOrderCode string `xml:"purchaseOrderCode"`
	WarehouseCode     string `xml:"warehouseCode"`
	OrderCreateTime   string `xml:"orderCreateTime"`
	OrderType         string `xml:"orderType"`
	RelatedOrders     struct {
		Text      string `xml:",chardata"`
		Remark    string `xml:"remark"`
		OrderType string `xml:"orderType"`
		OrderCode string `xml:"orderCode"`
	} `xml:"relatedOrders"`
	ExpectStartTime string `xml:"expectStartTime"`
	ExpectEndTime   string `xml:"expectEndTime"`
	LogisticsCode   string `xml:"logisticsCode"`
	LogisticsName   string `xml:"logisticsName"`
	ExpressCode     string `xml:"expressCode"`
	SupplierCode    string `xml:"supplierCode"`
	SupplierName    string `xml:"supplierName"`
	OperatorCode    string `xml:"operatorCode"`
	OperatorName    string `xml:"operatorName"`
	OperateTime     string `xml:"operateTime"`
	SenderInfo      struct {
		Text          string `xml:",chardata"`
		Remark        string `xml:"remark"`
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
		ID            string `xml:"id"`
		CarNo         string `xml:"carNo"`
	} `xml:"senderInfo"`
	ReceiverInfo struct {
		Text          string `xml:",chardata"`
		Remark        string `xml:"remark"`
		Company       string `xml:"company"`
		Name          string `xml:"name"`
		ZipCode       string `xml:"zipCode"`
		Tel           string `xml:"tel"`
		Mobile        string `xml:"mobile"`
		IdType        string `xml:"idType"`
		IdNumber      string `xml:"idNumber"`
		Email         string `xml:"email"`
		CountryCode   string `xml:"countryCode"`
		Province      string `xml:"province"`
		City          string `xml:"city"`
		Area          string `xml:"area"`
		Town          string `xml:"town"`
		DetailAddress string `xml:"detailAddress"`
		ID            string `xml:"id"`
		CarNo         string `xml:"carNo"`
	} `xml:"receiverInfo"`
	Remark                string `xml:"remark"`
	ContractCode          string `xml:"contractCode"`
	PlanArrivalTime       string `xml:"planArrivalTime"`
	Status                string `xml:"status"`
	SizeDetail            string `xml:"sizeDetail"`
	IsCheck               string `xml:"isCheck"`
	IsNudePackage         string `xml:"isNudePackage"`
	OrderSource           string `xml:"orderSource"`
	ExtOrderCode          string `xml:"extOrderCode"`
	ConsignId             string `xml:"consignId"`
	BusinessId            string `xml:"businessId"`
	LogisticsContactName  string `xml:"logisticsContactName"`
	LogisticsContactNo    string `xml:"logisticsContactNo"`
	LogisticsContactPhone string `xml:"logisticsContactPhone"`
	SupplierZipCode       string `xml:"supplierZipCode"`
	SupplierTel           string `xml:"supplierTel"`
	SupplierPhone         string `xml:"supplierPhone"`
	SupplierProvince      string `xml:"supplierProvince"`
	SupplierCity          string `xml:"supplierCity"`
	SupplierArea          string `xml:"supplierArea"`
	SupplierTown          string `xml:"supplierTown"`
	SupplierAddress       string `xml:"supplierAddress"`
	SupplierEmail         string `xml:"supplierEmail"`
	OutBizCode            string `xml:"outBizCode"`
	ConfirmType           string `xml:"confirmType"`
	TotalOrderLines       string `xml:"totalOrderLines"`
	WarehouseName         string `xml:"warehouseName"`
	SourceWarehouseCode   string `xml:"sourceWarehouseCode"`
	SourceWarehouseName   string `xml:"sourceWarehouseName"`
}
type OrderLines struct {
	OrderLineList []*OrderLine `xml:"orderLine"`
}
type OrderLine struct {
	OrderLineNo string `xml:"orderLineNo"`
	ItemId      string `xml:"itemId,omitempty"`
	// 货品编码
	ItemCode string `xml:"itemCode,omitempty"`
	//商品数量
	Quantity      int     `xml:"quantity,omitempty"`
	Amount        float64 `xml:"amount"`
	PurchasePrice float64 `xml:"purchasePrice"`
	ActualPrice   float64 `xml:"actualPrice"`
	//应收商品数量
	PlanQty   int    `xml:"planQty,omitempty"`
	OwnerCode string `xml:"ownerCode,omitempty"`
}
type Items struct {
	ItemList []*Item `xml:"item"`
}
type Item struct {
	Quantity int    `xml:"quantity"`
	ExCode   string `xml:"exCode,omitempty"`
	// 货品编码
	ItemId string `xml:"itemId,omitempty"`
	// 货品编码
	ItemCode      string `xml:"itemCode,omitempty"`
	InventoryType string `xml:"inventoryType"`
	PlanQty       int    `xml:"planQty"`
	//应收商品数量
	ActualQty int `xml:"actualQty,omitempty"`
}
type EntryOrderCreateResponse struct {
	XMLName xml.Name `xml:"response"`
	Response
	EntryOrderId string `xml:"entryOrderId,omitempty"`
}

func NewEntryOrderCreateSuccessResponse(entryOrderId string) *EntryOrderCreateResponse {
	dto := new(EntryOrderCreateResponse)
	dto.Flag = "success"
	dto.Code = "0"
	dto.EntryOrderId = entryOrderId
	return dto
}
