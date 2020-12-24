/*
文档地址：https://qimen.taobao.com/qimen/index.htm#/api-doc?fromPage=api-list&&apiServiceId=&officialApiId=26002&_k=xignfc
*/

package qimen

import (
	"encoding/xml"
	"time"

)

// 销售单下载同步解析
func DeliveryOrderCreateParse(body []byte) (res *DeliveryOrderCreateReqDto, err error) {
	res = new(DeliveryOrderCreateReqDto)
	err = xml.Unmarshal(body, res)
	return
}

type DeliveryOrderCreateReqDto struct {
	XMLName       xml.Name                `xml:"request"`
	DeliveryOrder *DeliveryOrderCreateDto `xml:"deliveryOrder"`
	OrderLines    *OrderLines             `xml:"orderLines"`
	ExtendProps   *ExtendProps            `xml:"extendProps"`
}
type DeliveryOrderCreateDto struct {
	WarehouseCode     string  `xml:"warehouseCode"`
	SourceOrderCode   string  `xml:"sourceOrderCode"`
	PayMethod         string  `xml:"payMethod"`
	ShopName          string  `xml:"shopName"`
	ReceiveOrderTime  string  `xml:"receiveOrderTime"`
	SellerId          string  `xml:"sellerId"`
	DeliveryNote      string  `xml:"deliveryNote"`
	ActualAmount      float32 `xml:"actualAmount"`
	BuyerName         string  `xml:"buyerName"`
	BuyerPhone        string  `xml:"buyerPhone"`
	PresaleOrderType  string  `xml:"presaleOrderType"`
	ItemCode          string  `xml:"itemCode"`
	ItemName          string  `xml:"itemName"`
	Quantity          string  `xml:"quantity"`
	Price             string  `xml:"price"`
	OrderNote         string  `xml:"orderNote"`
	LineNumber        int     `xml:"lineNumber"`
	Status            string  `xml:"status"`
	ConfirmType       int     `xml:"confirmType"`
	OrderConfirmTime  string  `xml:"orderConfirmTime"`
	OrderStatus       string  `xml:"orderStatus"`
	ShopCode          string  `xml:"shopCode"`
	DeliveryOrderCode string  `xml:"deliveryOrderCode"`
	// 原始订单号
	PreDeliveryOrderCode string `xml:"preDeliveryOrderCode"`
	OrderType            string `xml:"orderType"`
	OrderFlag            string `xml:"orderFlag"`
	//订单来源要特别注意（京东，淘宝，拼多多，三类在接收订单后，一定要传到当前实体对应的平台店铺上去。）比如京东来源订单一定要传到京东类型店铺上去。以上三个平台以外的都默认录淘宝平台就可以了。
	SourcePlatformCode string `xml:"sourcePlatformCode"`
	SourcePlatformName string `xml:"sourcePlatformName"`
	CreateTime         string `xml:"createTime"`
	PlaceOrderTime     string `xml:"placeOrderTime"`
	PayTime            string `xml:"payTime"`
	PayNo              string `xml:"payNo"`
	ShopNick           string `xml:"shopNick"`
	SellerNick         string `xml:"sellerNick"`
	BuyerNick          string `xml:"buyerNick"`
	//订单总金额(订单总金额=应收金额+已收金额=商品总金额-订单折扣金额+快递费用 ;单位 元)
	TotalAmount float32 `xml:"totalAmount"`
	//商品总金额(元)
	ItemAmount float32 `xml:"itemAmount"`
	//订单折扣金额(元)
	DiscountAmount float32 `xml:"discountAmount"`
	// 快递费用(元)
	Freight float32 `xml:"freight"`
	// 应收金额
	ArAmount float32 `xml:"arAmount"`
	// 实收金额
	GotAmount     float32 `xml:"gotAmount"`
	ServiceFee    float32 `xml:"serviceFee"`
	LogisticsCode string  `xml:"logisticsCode"`
	LogisticsName string  `xml:"logisticsName"`
	ExpressCode   string  `xml:"expressCode"`
	Remark        string  `xml:"remark"`
	BuyerMessage  string  `xml:"buyerMessage"`
	SellerMessage string  `xml:"sellerMessage"`
	// 发件人信息
	SenderInfo *DeliveryOrderAddress `xml:"senderInfo"`
	// 收件人信息
	ReceiverInfo *DeliveryOrderAddress `xml:"receiverInfo"`
}
type ExtendProps struct {
	Key1 string `xml:"key1"`
	Key2 string `xml:"key2"`
}
type DeliveryOrderAddress struct {
	Name          string `xml:"name"`
	ZipCode       string `xml:"zipCode"`
	Tel           string `xml:"tel"`
	Mobile        string `xml:"mobile"`
	Fax           string `xml:"fax"`
	Nick          string `xml:"nick"`
	Email         string `xml:"email"`
	CountryCode   string `xml:"countryCode"`
	Province      string `xml:"province"`
	City          string `xml:"city"`
	Area          string `xml:"area"`
	Town          string `xml:"town"`
	DetailAddress string `xml:"detailAddress"`
}

type DeliveryOrderCreateResponse struct {
	XMLName xml.Name `xml:"response"`
	Response
	DeliveryOrderId string       `xml:"DeliveryOrderId,omitempty"`
	CreateTime      time.Time    `xml:"createTime,omitempty"`
	WarehouseCode   string       `xml:"warehouseCode,omitempty"`
	LogisticsCode   string       `xml:"logisticsCode,omitempty"`
	OrderLines      []*OrderLine `xml:"orderLines"`
}

func NewDeliveryOrderCreateSuccessResponse(deliveryOrderId, warehouseCode, logisticsCode string) *DeliveryOrderCreateResponse {
	dto := new(DeliveryOrderCreateResponse)
	dto.Flag = "success"
	dto.Code = "0"
	dto.DeliveryOrderId = deliveryOrderId
	dto.WarehouseCode = warehouseCode
	dto.LogisticsCode = logisticsCode
	dto.CreateTime = time.Now()
	dto.OrderLines = make([]*OrderLine, 0)
	return dto
}
