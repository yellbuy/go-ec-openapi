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
	dto.CustomerCode = reqData.CustomerCode
	dto.CustomerName = reqData.CustomerName

	dto.NumPackage = "1"
	dto.OrderType = "JH_Normal"
	dto.BusinessType = "JH_Normal"
	//快递支付方式(立即付款=0，货到付款=1，发件人月结付款=2，收件人月结付款=3，预付款=4，银行转账=5，欠款=6，现金付款=7，第三方付费=8，寄方付=9，收方付=10)
	dto.PayMode = "9"
	dto.IsInsurance = "0"
	dto.OrderSource = "OTHERS"

	// 发件人
	dto.Sender = new(LogisticsAddress)
	dto.Sender.Name = reqData.SendName
	dto.Sender.Phone = reqData.SendPhone
	//dto.Sender.Mobile = reqData.SendName
	dto.Sender.Country = request.ShippingAddress.Country
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
	dto.Receiver.Country = reqData.ConsigneeAddress.Country
	dto.Receiver.Province = reqData.ConsigneeAddress.Province
	dto.Receiver.City = reqData.ConsigneeAddress.City
	dto.Receiver.Area = reqData.ConsigneeAddress.Area
	dto.Receiver.Town = reqData.ConsigneeAddress.Town
	dto.Receiver.Address = reqData.ConsigneeAddress.AddressDetail
	dto.LogisticType = reqData.ProductType
	// 京东接口需要传递ProviderCode
	dto.ProviderCode = reqData.ProductType
	dto.SiteCode = reqData.SiteCode
	dto.OrderSource = "0030001"
	//fmt.Println("dto.LogisticType:", dto.LogisticType)
	if client.Params.PlatId == "566" {
		// 拼多多需要先获取电子面单模板
		templateReq := new(common.WaybillTemplateRequest)
		//订单信息(所有模版=ALL，客户拥有的模版=OWNER)
		templateReq.TemplatesType = "ALL"
		// 快递公司类别
		templateReq.LogisticType = reqData.ProductType
		templateRes, body, err := client.GetWaybillTemplates(templateReq, extData...)
		if err != nil {
			return nil, body, err
		}
		if len(templateRes.Results) == 0 {
			return nil, body, fmt.Errorf("电子面单模板信息为空，平台id：%s", client.Params.PlatId)
		}
		//beego.Error("%+v", templateRes)
		dto.TemplateUrl = templateRes.Results[0].Url
		//fmt.Println("TemplateUrl:", dto.TemplateUrl)

		dto.OrderSource = "PDD"
	}
	goods := new(LogisticsGoods)
	goods.CnName = "商品"
	goods.Weight = "1"
	goods.Count = "1"
	dto.Goods = make([]*LogisticsGoods, 1)
	dto.Goods[0] = goods

	reqDto.Orders[0] = dto

	data, err := json.Marshal(reqDto)
	if err != nil {
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
	//fmt.Printf("params：", params)
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
	packages, hasPackages := waybill.CheckGet("packages")
	if hasPackages {
		pkg := packages.GetIndex(0)
		if pkg != nil {
			waybillInfo.PrintInfo, _ = pkg.Get("printinfo").String()
		}

	}
	//fmt.Println("waybillInfo.PrintInfo:", waybillInfo.PrintInfo)
	res.WaybillApplyNewInfo[0] = waybillInfo

	return res, body, err
}

// 查询电子面单信息
func (client *Client) GetWaybillTemplates(request *common.WaybillTemplateRequest, extData ...string) (*common.WaybillTemplateDto, []byte, error) {

	data, err := json.Marshal(request)
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
	//fmt.Printf("params：", params)
	// 通过polyapi自有平台
	method := "Differ.JH.Logistics.GetTemplates"
	res := new(common.WaybillTemplateDto)
	resJson := simplejson.New()
	resJson, body, err := client.Execute(method, params)
	if err != nil {
		fmt.Println(method, err)
		return res, body, err
	}
	root := resJson.Get("results").GetIndex(0)
	results, err := root.Get("results").Array()
	if err != nil {
		fmt.Println(method, err)
		return res, body, err
	}
	//fmt.Println("body:", string(body))
	res.Results = make([]*common.WaybillTemplateInfo, 0)
	for index := range results {
		result := root.Get("results").GetIndex(index)
		waybillInfo := new(common.WaybillTemplateInfo)
		waybillInfo.Id, _ = result.Get("id").String()
		waybillInfo.Name, _ = result.Get("name").String()
		waybillInfo.Url, _ = result.Get("url").String()
		waybillInfo.TemplateType, _ = result.Get("templatetype").String()
		res.Results = append(res.Results, waybillInfo)
	}
	return res, body, err
}

type LogisticsOrderReqDto struct {
	Orders []*LogisticsOrder `json:"orders"`
}
type LogisticsOrder struct {
	//订单号 必填
	OrderNo string `json:"orderno,omitempty"`
	//必填 平台原始单号(多个原始单号以英文“,”分隔
	PlatTradeNo string `json:"plattradeno,omitempty"`
	// 可选，是否为子母件(子母件=1，非子母件=0；默认0)
	IsMultiplePieces string `json:"ismultiplepieces,omitempty"`
	// 必填 包裹数量(默认填写值为1，有子母件时“IsMultiplePieces=1”包裹数量要和运单号数量一致)
	NumPackage string `json:"numpackage,omitempty"`
	//运单号(仅限于先预约单号的平台，如果是子母单，以半角逗号分隔，主单号在第一个位置，如755123456789,001123456789,002123456789)
	LogisticNo   string `json:"logisticno,omitempty"`
	BusinessType string `json:"businesstype,omitempty"`
	BusinessPlat string `json:"businessplat,omitempty"`
	// 必填：快递支付方式(立即付款=0，货到付款=1，发件人月结付款=2，收件人月结付款=3，预付款=4，银行转账=5，欠款=6，现金付款=7，第三方付费=8，寄方付=9，收方付=10)
	PayMode string `json:"paymode,omitempty"`
	// 必填：订单类型(普通订单=JH_Normal，退货单=JH_Refund，保价单=JH_Support，货到付款单=JH_COD，海外订单=JH_OverSea，便携式订单=JH_Portable，快递制单=JH_Express，仓储订单=JH_Storage)
	OrderType string `json:"ordertype,omitempty"`
	// 货到付款金额(OrderType=JH_COD时必传)
	CodPayMoney string `json:"codpaymoney,omitempty"`
	// 订单包裹物品金额
	PackageMoney   string `json:"packagemoney,omitempty"`
	Weight         string `json:"weight,omitempty"`
	SupporPayMoney string `json:"supporpaymoney,omitempty"`
	Length         string `json:"length,omitempty"`
	Width          string `json:"width,omitempty"`
	Height         string `json:"height,omitempty"`
	Volume         string `json:"volume,omitempty"`
	IsPickUp       string `json:"ispickup,omitempty"`
	ProductType    string `json:"producttype,omitempty"`
	LogisticType   string `json:"logistictype,omitempty"`
	// 承运公司编码
	CpCode      string            `json:"cpcode,omitempty"`
	DmsSorting  string            `json:"dmssorting,omitempty"`
	NeedEncrypt string            `json:"needencrypt,omitempty"`
	Sender      *LogisticsAddress `json:"sender,omitempty"`
	Receiver    *LogisticsAddress `json:"receiver,omitempty"`

	Goods        []*LogisticsGoods `json:"goods"`
	ShipperNo    string            `json:"shipperno,omitempty"`
	WareCode     string            `json:"warecode,omitempty"`
	BatchNo      string            `json:"batchno,omitempty"`
	IsLiquid     string            `json:"isliquid,omitempty"`
	IsHasBattery string            `json:"ishasbattery,omitempty"`
	// 是否保险，默认0
	IsInsurance           string `json:"isinsurance,omitempty"`
	DeliveryType          string `json:"deliverytype,omitempty"`
	BackSignBill          string `json:"backsignbill,omitempty"`
	IsLessTruck           string `json:"islesstruck,omitempty"`
	CustomerCode          string `json:"customercode,omitempty"`
	CustomerName          string `json:"customername,omitempty"`
	ImageStyle            string `json:"imagestyle,omitempty"`
	InputSite             string `json:"inputsite,omitempty"`
	SiteCode              string `json:"sitecode,omitempty"`
	IsCheckRange          string `json:"ischeckrange,omitempty"`
	TempRangeType         string `json:"temprangetype,omitempty"`
	MainSubPayMode        string `json:"mainsubpaymode,omitempty"`
	TransType             string `json:"transtype,omitempty"`
	IsBbc                 string `json:"isbbc,omitempty"`
	TransTypeCode         string `json:"transtypecode,omitempty"`
	ProdCode              string `json:"prodcode,omitempty"`
	IsNanJi               string `json:"isnanji,omitempty"`
	CurrencyType          string `json:"currencytype,omitempty"`
	PlatWebSite           string `json:"platwebsite,omitempty"`
	CrossCodeId           string `json:"crosscodeid,omitempty"`
	DefinedOrderInfo      string `json:"definedorderinfo,omitempty"`
	DefinedGoodsInfo      string `json:"definedgoodsinfo,omitempty"`
	PayOrderNo            string `json:"payorderno,omitempty"`
	PayAmount             string `json:"payamount,omitempty"`
	SenderAccount         string `json:"senderaccount,omitempty"`
	Is5KgPacking          string `json:"is5kgpacking,omitempty"`
	Is3Pl                 string `json:"is3pl,omitempty"`
	IsuseStock            string `json:"isusestock,omitempty"`
	IsEconomic            string `json:"iseconomic,omitempty"`
	SpecialHandling       string `json:"specialhandling,omitempty"`
	CodType               string `json:"codtype,omitempty"`
	PackageService        string `json:"packageservice,omitempty"`
	IsOut                 string `json:"isout,omitempty"`
	ReceiverAccount       string `json:"receiveraccount,omitempty"`
	ReceiverAccountname   string `json:"receiveraccountname,omitempty"`
	IsFresh               string `json:"isfresh,omitempty"`
	Remark                string `json:"remark,omitempty"`
	CustomerRemark        string `json:"customerremark,omitempty"`
	OrderSource           string `json:"ordersource,omitempty"`
	ProviderId            string `json:"providerid,omitempty"`
	ProviderCode          string `json:"providercode,omitempty"`
	ExpressPayMethod      string `json:"expresspaymethod,omitempty"`
	UndeliverableDecision string `json:"undeliverabledecision,omitempty"`
	ServiceName           string `json:"servicename,omitempty"`
	CumstomsCode          string `json:"cumstomscode,omitempty"`
	TotalLogisticsNo      string `json:"totallogisticsno,omitempty"`
	StockFlag             string `json:"stockflag,omitempty"`
	EbpCode               string `json:"ebpcode,omitempty"`
	EcpName               string `json:"ecpname,omitempty"`
	EcpCodeG              string `json:"ecpcodeg,omitempty"`
	EcpNameG              string `json:"ecpnameg,omitempty"`
	AgentCode             string `json:"agentcode,omitempty"`
	AgentName             string `json:"agentname,omitempty"`
	TotalShippingFee      string `json:"totalshippingfee,omitempty"`
	FeeUnit               string `json:"feeunit,omitempty"`
	ClearCode             string `json:"clearcode,omitempty"`
	SellerId              string `json:"sellerid,omitempty"`
	UserId                string `json:"user_id,omitempty"`
	LogisticsServices     string `json:"logistics_services,omitempty"`
	ObjectId              string `json:"object_id,omitempty"`
	TemplateUrl           string `json:"template_url,omitempty"`
	OrderChannelsType     string `json:"order_channels_type,omitempty"`
	TradeOrderList        string `json:"trade_order_list,omitempty"`
	LogisticsProductName  string `json:"logisticsproductname,omitempty"`
	DeptNo                string `json:"deptno,omitempty"`
	SenderTc              string `json:"sendertc,omitempty"`
	PickUpDate            string `json:"pickupdate,omitempty"`
	InstallFlag           string `json:"installflag,omitempty"`
	ThirdCategoryNo       string `json:"thirdcategoryno,omitempty"`
	BrandNo               string `json:"brandno,omitempty"`
	ProductSku            string `json:"productsku,omitempty"`
	PlatCode              string `json:"platcode,omitempty"`
	SequenceNo            string `json:"sequenceno,omitempty"`
	ChinaShipName         string `json:"chinashipname,omitempty"`
	TaxPayType            string `json:"taxpaytype,omitempty"`
	TaxSetAccounts        string `json:"taxsetaccounts,omitempty"`
	PracelType            string `json:"praceltype,omitempty"`
	AddressId             string `json:"addressid,omitempty"`
	ConsignPreferenceId   string `json:"consignpreferenceid,omitempty"`
	IsKuaiYun             string `json:"iskuaiyun,omitempty"`
}
type LogisticsAddress struct {
	Name            string `json:"name,omitempty"`
	Company         string `json:"company,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Mobile          string `json:"mobile,omitempty"`
	Country         string `json:"country,omitempty"`
	Province        string `json:"province,omitempty"`
	City            string `json:"city,omitempty"`
	Area            string `json:"area,omitempty"`
	Town            string `json:"town,omitempty"`
	Address         string `json:"address,omitempty"`
	Zip             string `json:"zip,omitempty"`
	Email           string `json:"email,omitempty"`
	UserId          string `json:"userid,omitempty"`
	CertificateType string `json:"certificatetype,omitempty"`
	Certificate     string `json:"certificate,omitempty"`
	CertificateName string `json:"certificatename,omitempty"`
	AddressCode     string `json:"addresscode,omitempty"`
	Linker          string `json:"linker,omitempty"`
	TaxPayerIdent   string `json:"taxpayerident,omitempty"`
}

type LogisticsGoods struct {
	CnName          string `json:"cnname,omitempty"`
	EnName          string `json:"enname,omitempty"`
	Count           string `json:"count,omitempty"`
	CurrencyType    string `json:"currencytype,omitempty"`
	Price           string `json:"price,omitempty"`
	Weight          string `json:"weight,omitempty"`
	Unit            string `json:"unit,omitempty"`
	TaxationId      string `json:"taxationid,omitempty"`
	ProductId       string `json:"productid,omitempty"`
	InnerCount      string `json:"innercount,omitempty"`
	Length          string `json:"length,omitempty"`
	Width           string `json:"width,omitempty"`
	Height          string `json:"height,omitempty"`
	DutyMoney       string `json:"dutymoney,omitempty"`
	IsBlinsure      string `json:"isblinsure,omitempty"`
	Remark          string `json:"remark,omitempty"`
	IsAnerOidMarkUp string `json:"isaneroidmarkup,omitempty"`
	IsOnlyBattery   string `json:"isonlybattery,omitempty"`
	ProductBrand    string `json:"productbrand,omitempty"`
	ProductAttrs    string `json:"productattrs,omitempty"`
	ProductMaterial string `json:"productmaterial,omitempty"`
	HsCode          string `json:"hscode,omitempty"`
	GoodUrl         string `json:"goodurl,omitempty"`
	CategoryId      string `json:"categoryid,omitempty"`
	CategoryId2     string `json:"categoryid2,omitempty"`
	PlatTradeNo     string `json:"plattradeno,omitempty"`
	OriginCountry   string `json:"origincountry,omitempty"`
	OuterId         string `json:"outerid,omitempty"`
	Position        string `json:"position,omitempty"`
	SupportBattery  string `json:"supportbattery,omitempty"`
	Description     string `json:"description,omitempty"`
	ElecQuaId       string `json:"elecquaid,omitempty"`
}
type Address struct {
	Country string `json:"country,omitempty"`
	// 必填,省
	Province string `json:"province,omitempty"`
	City     string `json:"city,omitempty"`
	District string `json:"district,omitempty"`
	Town     string `json:"town,omitempty"`
	// 必填.详细地址
	Detail string `json:"detail,omitempty"`
}
