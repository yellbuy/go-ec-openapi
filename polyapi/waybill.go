package polyapi

import (
	"encoding/json"
	"errors"
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

func (client *Client) GetWaybill(request *common.WaybillApplyNewRequest, extData ...string) (*common.WaybillApplyNewCols, []byte, error) {
	if len(request.TradeOrderInfoCols) == 0 {
		return nil, nil, errors.New("订单信息不能为空")
	}
	reqData := request.TradeOrderInfoCols[0]
	if len(reqData.TradeOrderList) == 0 {
		return nil, nil, errors.New("订单数据不能为空")
	}
	reqDto := new(LogisticsOrderReqDto)
	reqDto.Orders = make([]*LogisticsOrder, 1)
	dto := new(LogisticsOrder)
	dto.OrderNo = reqData.OrderNo
	dto.PlatTradeNo = reqData.PlatTradeNo
	dto.NumPackage = "1"
	dto.OrderType = "JH_Normal"
	//快递支付方式(立即付款=0，货到付款=1，发件人月结付款=2，收件人月结付款=3，预付款=4，银行转账=5，欠款=6，现金付款=7，第三方付费=8，寄方付=9，收方付=10)
	//dto.PayMode="0"
	dto.IsInsurance = "0"

	// 发件人
	dto.Sender = new(LogisticsAddress)
	dto.Sender.Name = reqData.SendName
	dto.Sender.Phone = reqData.SendPhone
	//dto.Sender.Mobile = reqData.SendName
	dto.Sender.Province = request.ShippingAddress.Province
	dto.Sender.City = request.ShippingAddress.City
	dto.Sender.Area = request.ShippingAddress.Area
	dto.Sender.Town = request.ShippingAddress.Town
	dto.Sender.Address = request.ShippingAddress.AddressDetail
	// 收件人
	dto.Receiver = new(LogisticsAddress)
	dto.Receiver = new(LogisticsAddress)
	dto.Receiver.Name = reqData.ConsigneeName
	dto.Receiver.Phone = reqData.ConsigneePhone
	//dto.Sender.Mobile = reqData.SendName
	dto.Receiver.Province = reqData.ConsigneeAddress.Province
	dto.Receiver.City = reqData.ConsigneeAddress.City
	dto.Receiver.Area = reqData.ConsigneeAddress.Area
	dto.Receiver.Town = reqData.ConsigneeAddress.Town
	dto.Receiver.Address = reqData.ConsigneeAddress.AddressDetail
	reqDto.Orders[0] = dto
	data, err := json.Marshal(reqDto)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	reqJson, err := simplejson.NewJson(data)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	if len(extData) > 0 {
		reqJson.Set("platvalue", extData[0])
	}
	if len(extData) > 1 {
		reqJson.Set("polyapitoken", extData[1])
	}
	if len(extData) > 2 {
		reqJson.Set("shoptype", extData[2])
	} else {
		reqJson.Set("shoptype", "JH_001")
	}

	bizcontent, err := reqJson.Encode()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)

	//fmt.Println("bizcontent：", string(bizcontent))
	params, err := common.InterfaceToParameter(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	//fmt.Println("bizcontent2：", string(bizcontent))
	// 通过polyapi自有平台
	method := "Differ.JH.Logistics.PostOrder"
	res := new(common.WaybillApplyNewCols)
	resJson := simplejson.New()
	resJson, body, err := client.Execute(method, params)
	if err != nil {
		fmt.Println(method, err)
		return res, body, err
	}
	res.WaybillApplyNewInfo = make([]*common.WaybillApplyNewInfo, 1)
	waybill := resJson.Get("results").GetIndex(0)
	waybillInfo := new(common.WaybillApplyNewInfo)
	waybillInfo.WaybillCode, _ = waybill.Get("logisticno").String()
	waybillInfo.PackageCenterName, _ = waybill.Get("destcode").String()
	waybillInfo.ShortAddress, _ = waybill.Get("markers").String()
	res.WaybillApplyNewInfo[0] = waybillInfo

	return res, body, err
}

type LogisticsOrderReqDto struct {
	Orders []*LogisticsOrder `json:"orders"`
}
type LogisticsOrder struct {
	//订单号 必填
	OrderNo string `json:"orderno"`
	//必填 平台原始单号(多个原始单号以英文“,”分隔
	PlatTradeNo string `json:"plattradeno"`
	// 可选，是否为子母件(子母件=1，非子母件=0；默认0)
	IsMultiplePieces string `json:"ismultiplepieces"`
	// 必填 包裹数量(默认填写值为1，有子母件时“IsMultiplePieces=1”包裹数量要和运单号数量一致)
	NumPackage string `json:"numpackage"`
	//运单号(仅限于先预约单号的平台，如果是子母单，以半角逗号分隔，主单号在第一个位置，如755123456789,001123456789,002123456789)
	LogisticNo   string `json:"logisticno"`
	BusinessType string `json:"businesstype"`
	BusinessPlat string `json:"businessplat"`
	// 必填：快递支付方式(立即付款=0，货到付款=1，发件人月结付款=2，收件人月结付款=3，预付款=4，银行转账=5，欠款=6，现金付款=7，第三方付费=8，寄方付=9，收方付=10)
	PayMode string `json:"paymode"`
	// 必填：订单类型(普通订单=JH_Normal，退货单=JH_Refund，保价单=JH_Support，货到付款单=JH_COD，海外订单=JH_OverSea，便携式订单=JH_Portable，快递制单=JH_Express，仓储订单=JH_Storage)
	OrderType string `json:"ordertype"`
	// 货到付款金额(OrderType=JH_COD时必传)
	CodPayMoney string `json:"codpaymoney"`
	// 订单包裹物品金额
	PackageMoney   string `json:"packagemoney"`
	Weight         string `json:"weight"`
	SupporPayMoney string `json:"supporpaymoney"`
	Length         string `json:"length"`
	Width          string `json:"width"`
	Height         string `json:"height"`
	Volume         string `json:"volume"`
	IsPickUp       string `json:"ispickup"`
	ProductType    string `json:"producttype"`
	LogisticType   string `json:"logistictype"`
	// 承运公司编码
	CpCode      string            `json:"cpcode"`
	DmsSorting  string            `json:"dmssorting"`
	NeedEncrypt string            `json:"needencrypt"`
	Sender      *LogisticsAddress `json:"sender"`
	Receiver    *LogisticsAddress `json:"receiver"`

	Goods        []*LogisticsGoods `json:"goods"`
	ShipperNo    string            `json:"shipperno"`
	WareCode     string            `json:"warecode"`
	BatchNo      string            `json:"batchno"`
	IsLiquid     string            `json:"isliquid"`
	IsHasBattery string            `json:"ishasbattery"`
	// 是否保险，默认0
	IsInsurance           string `json:"isinsurance"`
	DeliveryType          string `json:"deliverytype"`
	BackSignBill          string `json:"backsignbill"`
	IsLessTruck           string `json:"islesstruck"`
	CustomerCode          string `json:"customercode"`
	CustomerName          string `json:"customername"`
	ImageStyle            string `json:"imagestyle"`
	InputSite             string `json:"inputsite"`
	SiteCode              string `json:"sitecode"`
	IsCheckRange          string `json:"ischeckrange"`
	TempRangeType         string `json:"temprangetype"`
	MainSubPayMode        string `json:"mainsubpaymode"`
	TransType             string `json:"transtype"`
	IsBbc                 string `json:"isbbc"`
	TransTypeCode         string `json:"transtypecode"`
	ProdCode              string `json:"prodcode"`
	IsNanJi               string `json:"isnanji"`
	CurrencyType          string `json:"currencytype"`
	PlatWebSite           string `json:"platwebsite"`
	CrossCodeId           string `json:"crosscodeid"`
	DefinedOrderInfo      string `json:"definedorderinfo"`
	DefinedGoodsInfo      string `json:"definedgoodsinfo"`
	PayOrderNo            string `json:"payorderno"`
	PayAmount             string `json:"payamount"`
	SenderAccount         string `json:"senderaccount"`
	Is5KgPacking          string `json:"is5kgpacking"`
	Is3Pl                 string `json:"is3pl"`
	IsuseStock            string `json:"isusestock"`
	IsEconomic            string `json:"iseconomic"`
	SpecialHandling       string `json:"specialhandling"`
	CodType               string `json:"codtype"`
	PackageService        string `json:"packageservice"`
	IsOut                 string `json:"isout"`
	ReceiverAccount       string `json:"receiveraccount"`
	ReceiverAccountname   string `json:"receiveraccountname"`
	IsFresh               string `json:"isfresh"`
	Remark                string `json:"remark"`
	CustomerRemark        string `json:"customerremark"`
	OrderSource           string `json:"ordersource"`
	ProviderId            string `json:"providerid"`
	ProviderCode          string `json:"providercode"`
	ExpressPayMethod      string `json:"expresspaymethod"`
	UndeliverableDecision string `json:"undeliverabledecision"`
	ServiceName           string `json:"servicename"`
	CumstomsCode          string `json:"cumstomscode"`
	TotalLogisticsNo      string `json:"totallogisticsno"`
	StockFlag             string `json:"stockflag"`
	EbpCode               string `json:"ebpcode"`
	EcpName               string `json:"ecpname"`
	EcpCodeG              string `json:"ecpcodeg"`
	EcpNameG              string `json:"ecpnameg"`
	AgentCode             string `json:"agentcode"`
	AgentName             string `json:"agentname"`
	TotalShippingFee      string `json:"totalshippingfee"`
	FeeUnit               string `json:"feeunit"`
	ClearCode             string `json:"clearcode"`
	SellerId              string `json:"sellerid"`
	UserId                string `json:"user_id"`
	LogisticsServices     string `json:"logistics_services"`
	ObjectId              string `json:"object_id"`
	TemplateUrl           string `json:"template_url"`
	OrderChannelsType     string `json:"order_channels_type"`
	TradeOrderList        string `json:"trade_order_list"`
	LogisticsProductName  string `json:"logisticsproductname"`
	DeptNo                string `json:"deptno"`
	SenderTc              string `json:"sendertc"`
	PickUpDate            string `json:"pickupdate"`
	InstallFlag           string `json:"installflag"`
	ThirdCategoryNo       string `json:"thirdcategoryno"`
	BrandNo               string `json:"brandno"`
	ProductSku            string `json:"productsku"`
	PlatCode              string `json:"platcode"`
	SequenceNo            string `json:"sequenceno"`
	ChinaShipName         string `json:"chinashipname"`
	TaxPayType            string `json:"taxpaytype"`
	TaxSetAccounts        string `json:"taxsetaccounts"`
	PracelType            string `json:"praceltype"`
	AddressId             string `json:"addressid"`
	ConsignPreferenceId   string `json:"consignpreferenceid"`
	IsKuaiYun             string `json:"iskuaiyun"`
}
type LogisticsAddress struct {
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
	CertificateName string `json:"certificatename"`
	AddressCode     string `json:"addresscode"`
	Linker          string `json:"linker"`
	TaxPayerIdent   string `json:"taxpayerident"`
}

type LogisticsGoods struct {
	CnName          string `json:"cnname"`
	EnName          string `json:"enname"`
	Count           string `json:"count"`
	CurrencyType    string `json:"currencytype"`
	Price           string `json:"price"`
	Weight          string `json:"weight"`
	Unit            string `json:"unit"`
	TaxationId      string `json:"taxationid"`
	ProductId       string `json:"productid"`
	InnerCount      string `json:"innercount"`
	Length          string `json:"length"`
	Width           string `json:"width"`
	Height          string `json:"height"`
	DutyMoney       string `json:"dutymoney"`
	IsBlinsure      string `json:"isblinsure"`
	Remark          string `json:"remark"`
	IsAnerOidMarkUp string `json:"isaneroidmarkup"`
	IsOnlyBattery   string `json:"isonlybattery"`
	ProductBrand    string `json:"productbrand"`
	ProductAttrs    string `json:"productattrs"`
	ProductMaterial string `json:"productmaterial"`
	HsCode          string `json:"hscode"`
	GoodUrl         string `json:"goodurl"`
	CategoryId      string `json:"categoryid"`
	CategoryId2     string `json:"categoryid2"`
	PlatTradeNo     string `json:"plattradeno"`
	OriginCountry   string `json:"origincountry"`
	OuterId         string `json:"outerid"`
	Position        string `json:"position"`
	SupportBattery  string `json:"supportbattery"`
	Description     string `json:"description"`
	ElecQuaId       string `json:"elecquaid"`
}
type Address struct {
	Country string `json:"country"`
	// 必填,省
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Town     string `json:"town"`
	// 必填.详细地址
	Detail string `json:"detail"`
}
