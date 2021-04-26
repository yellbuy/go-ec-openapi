package common

type LogisticsSendReqDto struct {
	PlatOrderNo  string `json:"platorderno"`
	LogisticNo   string `json:"logisticno"`
	LogisticType string `json:"logistictype"`
	LogisticName string `json:"logisticname"`
	SendType     string `json:"sendtype"`
	IsSplit      string `json:"issplit"`

	ShopType      string           `json:"shoptype"`
	SenderName    string           `json:"sendername"`
	SenderTel     string           `json:"sendertel"`
	SenderAddress string           `json:"senderaddress"`
	IsHwgFlag     string           `json:"ishwgflag"`
	DeliveryType  string           `json:"deliverytype"`
	PackageNumber string           `json:"packagenumber"`
	Goods         []*LogisticGoods `json:"goods"`
	// LogisticFee   struct {
	// } `json:"logisticfee"`
	IsModSendInfo       string `json:"ismodsendinfo"`
	WhseCode            string `json:"whsecode"`
	CountryCode         string `json:"countrycode"`
	SellerNick          string `json:"sellernick"`
	PackageOrderId      string `json:"packageorderid"`
	Website             string `json:"website"`
	OlderLogisticNo     string `json:"olderlogisticno"`
	OlderLogisticType   string `json:"olderlogistictype"`
	InvoiceNo           string `json:"invoiceno"`
	InvoiceDate         string `json:"invoicedate"`
	TradeTime           string `json:"tradetime"`
	TransactionId       string `json:"transactionid"`
	FxtId               string `json:"fxtid"`
	IsDomesticDelivery  string `json:"isdomesticdelivery"`
	ZyLogisticCompany   string `json:"zylogisticcompany"`
	ZyLogisticPhone     string `json:"zylogisticphone"`
	InstallLogisticName string `json:"installlogisticname"`
	InstallLogisticCode string `json:"installlogisticcode"`
	Feature             string `json:"feature"`
	ResId               string `json:"resid"`
	ResCode             string `json:"rescode"`
	OrderStatus         string `json:"orderstatus"`
	IsRequestWaybill    string `json:"isrequestwaybill"`
	// WaybillRequest      struct {
	// 	LogisticNo     string        `json:"logisticno"`
	// 	BusinessPlat   string        `json:"businessplat"`
	// 	CustomerCode   string        `json:"customercode"`
	// 	OrderNo        string        `json:"orderno"`
	// 	PlatTradeNo    string        `json:"plattradeno"`
	// 	NumPackage     string        `json:"numpackage"`
	// 	Weight         string        `json:"weight"`
	// 	IsInsurance    string        `json:"isinsurance"`
	// 	SupporPayMoney string        `json:"supporpaymoney"`
	// 	IsCod          string        `json:"iscod"`
	// 	CodPayMoney    string        `json:"codpaymoney"`
	// 	BusinessType   string        `json:"businesstype"`
	// 	WareCode       string        `json:"warecode"`
	// 	Sender         *ShippingInfo `json:"sender"`
	// 	Receiver       *ShippingInfo `json:"receiver"`
	// } `json:"waybillrequest"`
	ReceiverAddress string `json:"receiveraddress"`
	OrderType       string `json:"ordertype"`
	ShopId          string `json:"shopid"`
	DeliveryName    string `json:"deliveryname"`
	DeliveryPhone   string `json:"deliveryphone"`
	CustomerRemark  string `json:"customerremark"`
	GiftCount       string `json:"giftcount"`
	OrderNumber     string `json:"ordernumber"`
	VerifyCode      string `json:"verifycode"`
	EventCreateTime string `json:"eventcreatetime"`
	PlanTime        string `json:"plantime"`
	Desc            string `json:"desc"`
	SequenceNo      string `json:"sequenceno"`
	Event           string `json:"event"`
}

type LogisticGoods struct {
	SubOrderNo string `json:"suborderno"`
	Count      string `json:"count"`

	PlatProductId   string `json:"platproductid"`
	TradeGoodsNo    string `json:"tradegoodsno"`
	SubLogisticType string `json:"sublogistictype"`
	SubLogisticName string `json:"sublogisticname"`
	SubLogisticNo   string `json:"sublogisticno"`
	Barcode         string `json:"barcode"`
	Isgift          string `json:"isgift"`
	ProductItemId   string `json:"productitemid"`
	SkuId           string `json:"skuid"`
	DeliveryType    string `json:"deliverytype"`
	PayOrderId      string `json:"payorderid"`
	PackageOrderId  string `json:"packageorderid"`
	Price           string `json:"price"`
	RealCount       string `json:"realcount"`
	NumPackage      string `json:"numpackage"`
}

type ShippingInfo struct {
	Name            string `json:"name"`
	Company         string `json:"company"`
	Phone           string `json:"phone"`
	Mobile          string `json:"mobile"`
	Country         string `json:"country"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Area            string `json:"area"`
	Town            string `json:"town"`
	Address         string `json:"address"`
	Zip             string `json:"zip"`
	Email           string `json:"email"`
	UserId          string `json:"userid"`
	CertificateType string `json:"certificatetype"`
	Certificate     string `json:"certificate"`
}

func GetStrCn(str string) (cnStr string) {
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			cnStr = cnStr + string(r[i])
		}
	}
	return
}
