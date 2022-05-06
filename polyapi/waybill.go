package polyapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

func (client *Client) OrderSendV3(request []*common.WmsOrderBatchSend, extData ...string) (common.WmsOrderBatchSendReturn, error) {
	method := "Differ.JH.Business.BatchSend" //菠萝派批量同步接口
	bizcontent, err := json.Marshal(request)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	_, body, err := client.Execute(method, params)
	//logs.Debug(string(body))
	var OutData common.WmsOrderBatchSendReturn
	if err != nil {
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	return OutData, err
}
func (client *Client) LogisticsPrintOrderList(request []*common.WmsLogisticsPrintOrderBizcontent, extData ...string) (common.WmsLogisticsPrintOrderReturn, error) {
	method := "Differ.JH.Logistics.BatchPrintOrder" //菠萝派批量物流打印接口
	orders := new(common.WmsLogisticsPrintOrderBizcontentOrders)
	orders.Orders = request
	bizcontent, err := json.Marshal(orders)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	_, body, err := client.Execute(method, params)
	//logs.Debug(string(body))
	var OutData common.WmsLogisticsPrintOrderReturn
	if err != nil {
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	return OutData, err
}
func (client *Client) LogisticsPostOrder(request common.WmsLogisticsPostOrder, extData ...string) (common.WmsLogisticsReturn, error) {
	method := "Differ.JH.Logistics.PostOrder" //定义菠萝派退款检测批量接口
	bizcontent, err := json.Marshal(request)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	_, body, err := client.Execute(method, params)
	//logs.Debug(string(body))
	var OutData common.WmsLogisticsReturn
	if err != nil {
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	return OutData, err
}
func (client *Client) CheckRefundV2(request common.BatchCheckRefundStatusBizcontent, extData ...string) (common.CheckRefundReturn, error) {
	method := "Differ.JH.Business.BatchCheckRefundStatus" //定义菠萝派退款检测批量接口
	bizcontent, err := json.Marshal(request)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	_, body, err := client.Execute(method, params)
	//logs.Debug(string(body))
	var OutData common.CheckRefundReturn
	if err != nil {
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	if OutData.Code == "10000" {
		err = nil
	}
	return OutData, err
}

func (client *Client) DownloadOrderListV2(request common.DownLoadOrderListPostBizcontent, extData ...string) (common.DownloadOrderListReturn, error) {
	method := "Differ.JH.Business.GetOrder" //定义菠萝派订单下载接口
	bizcontent, err := json.Marshal(request)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	_, body, err := client.Execute(method, params)
	//logs.Debug(string(body))
	var OutData common.DownloadOrderListReturn
	if err != nil {
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	if OutData.Code == "10000" {
		err = nil
	}
	//logs.Debug("下载订单完成，菠萝派ID[", OutData.Polyapirequestid, "]")
	//logs.Error("下载JSON内容[", string(body), "]")
	return OutData, err
}
func (client *Client) CancelWaybill(request []common.WaybillCancel, extData ...string) (*common.WaybillCancelReturn, error) {
	//开始提交数据
	method := "Differ.JH.Logistics.Cancel"
	var reqA common.WaybillCancelSend
	reqA.Orders = request
	bizcontent, err := json.Marshal(reqA)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	//此处可能还要加工Json
	jsonstr, body, err := client.Execute(method, params)
	if err != nil {
		return nil, err
	}
	logs.Debug("返回内容:", jsonstr)
	var OutData common.WaybillCancelReturn
	err = json.Unmarshal(body, &OutData)
	return &OutData, err
}
func (client *Client) TBDecrypt(request []*BusinessBatchTBDecryptOrders, extData ...string) (*BusinessBatchTBDecryptReturn, error) {
	//开始提交数据
	method := "Differ.JH.Business.BatchTBDecrypt"
	var reqA BusinessBatchTBDecryptBizcontent
	reqA.Orders = request
	bizcontent, err := json.Marshal(reqA)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	//此处可能还要加工Json
	_, body, err := client.Execute(method, params)
	if err != nil {
		return nil, err
	}
	var OutData BusinessBatchTBDecryptReturn
	err = json.Unmarshal(body, &OutData)
	return &OutData, err
}

//同步发货接口V2
func (client *Client) OrderSendV2(request *common.WmsBusinessSendBizcontent) (*common.WmsBusinessSendReturn, error) {
	//开始提交数据
	method := "Differ.JH.Business.Send"
	bizcontent, err := json.Marshal(request)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	//此处可能还要加工Json
	_, body, err := client.Execute(method, params)
	if err != nil {
		return nil, err
	}
	var OutData common.WmsBusinessSendReturn
	err = json.Unmarshal(body, &OutData)
	return &OutData, err
}

//换取抖音打印参数
func (client *Client) GetDyPlatApiReuestInfo(postData string) (*common.WmsPlatApiReturnInfo, error) {
	method := "Differ.JH.Business.GetPlatApiRequestInfo"
	var reqA common.WmsPlatApiReuestInfo
	reqA.PlatMethod = postData
	bizcontent, _ := json.Marshal(reqA)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	//此处可能还要加工Json
	_, body, err := client.Execute(method, params)
	if err != nil {
		var OutData common.WmsPlatApiReturnInfo
		json.Unmarshal(body, &OutData)
		return &OutData, err
	}
	var OutData common.WmsPlatApiReturnInfo
	err = json.Unmarshal(body, &OutData)
	return &OutData, err
}

//预约单号接口V2
func (client *Client) GetWaybillV2(request []*common.WmsLogisticsPostOrder) (*common.WmsLogisticsReturn, error) {
	//开始提交数据
	method := "Differ.JH.Logistics.PostOrder"
	var reqA common.WmsLogisticsBizcontent
	reqA.Orders = request
	bizcontent, err := json.Marshal(reqA)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	//此处可能还要加工Json
	_, body, err := client.Execute(method, params)
	if err != nil {
		var OutData common.WmsLogisticsReturn
		json.Unmarshal(body, &OutData)
		return &OutData, err
	}
	var OutData common.WmsLogisticsReturn
	err = json.Unmarshal(body, &OutData)
	return &OutData, err
}
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
	dto.Oaid = reqData.Oaid
	logs.Debug("提交的内容", reqData)
	dto.OrderNo = reqData.OrderNo
	dto.PlatTradeNo = reqData.PlatTradeNo
	if len(request.LogisticsServices) > 0 {
		dto.LogisticsServices = request.LogisticsServices
	}
	// 月结账号
	dto.CustomerCode = reqData.CustomerCode
	dto.CustomerName = reqData.CustomerName
	// 是否保价
	dto.IsInsurance = reqData.IsInsurance
	//保价金额
	dto.SupporPayMoney = reqData.SupporPayMoney
	// 销售平台
	dto.BusinessPlat = reqData.BusinessPlat
	if dto.BusinessPlat == "" {
		dto.BusinessPlat = "OTHERS"
	}

	dto.NumPackage = "1"
	dto.OrderType = reqData.OrderType
	if dto.OrderType == "" {
		dto.OrderType = "JH_Normal"
	}
	dto.CodPayMoney = reqData.CodPayMoney

	if request.BusinessType != "0" {
		dto.BusinessType = fmt.Sprintf("%d", request.BusinessType)
	} else {
		dto.BusinessType = "JH_Normal"
	}
	//快递支付方式(立即付款=0，货到付款=1，发件人月结付款=2，收件人月结付款=3，预付款=4，银行转账=5，欠款=6，现金付款=7，第三方付费=8，寄方付=9，收方付=10)

	if request.PayMode > 0 {
		dto.PayMode = fmt.Sprintf("%d", request.PayMode)
	} else {
		dto.PayMode = "9"
	}
	dto.ShipperNo = request.ShipperNo
	dto.OrderSource = "OTHERS"
	dto.LogisticsProductName = "JH_02"

	// 发件人
	dto.Sender = new(LogisticsAddress)
	dto.Sender.Name = reqData.SendName
	dto.Sender.Phone = reqData.SendPhone
	//dto.Sender.Mobile = reqData.SendName
	dto.Sender.Country = request.ShippingAddress.Country
	dto.Sender.Province = common.GetStrCn(request.ShippingAddress.Province)
	dto.Sender.City = common.GetStrCn(request.ShippingAddress.City)
	dto.Sender.Area = common.GetStrCn(request.ShippingAddress.Area)
	dto.Sender.Town = request.ShippingAddress.Town
	dto.Sender.Address = request.ShippingAddress.AddressDetail
	// 收件人
	dto.Receiver = new(LogisticsAddress)
	dto.Receiver = new(LogisticsAddress)
	dto.Receiver.Name = reqData.ConsigneeName
	dto.Receiver.Phone = reqData.ConsigneePhone
	dto.Receiver.Country = common.GetStrCn(reqData.ConsigneeAddress.Country)
	dto.Receiver.Province = common.GetStrCn(reqData.ConsigneeAddress.Province)
	dto.Receiver.City = common.GetStrCn(reqData.ConsigneeAddress.City)
	dto.Receiver.Area = reqData.ConsigneeAddress.Area
	dto.Receiver.Town = reqData.ConsigneeAddress.Town
	dto.Receiver.Address = reqData.ConsigneeAddress.AddressDetail
	dto.LogisticType = reqData.ProductType
	//fmt.Println("dto.LogisticType:", dto.LogisticType)

	// 京东接口需要传递ProviderCode
	dto.ProviderCode = reqData.ProductType
	dto.SiteCode = reqData.SiteCode
	dto.OrderSource = "0030001"
	dto.CpCode = request.CpCode
	// EMS相关参数
	dto.MainSubPayMode = "1"
	dto.TransType = "0"
	dto.TransTypeCode = "1"
	dto.ProdCode = "0300000000"
	dto.IsNanJi = "1"
	dto.NeedEncrypt = 1

	if client.Params.PlatId == "548" || client.Params.PlatId == "566" {
		// 菜鸟、拼多多需要先获取电子面单模板
		templateReq := new(common.WaybillTemplateRequest)
		//订单信息(所有模版=ALL，客户拥有的模版=OWNER)
		templateReq.TemplatesType = "ALL"
		// 快递公司类别
		templateReq.LogisticType = reqData.ProductType
		templateRes, body, err := client.GetWaybillTemplates(templateReq, extData...)
		if err != nil {
			return nil, body, err
		}
		templateLen := len(templateRes.Results)
		if templateLen == 0 {
			return nil, body, fmt.Errorf("电子面单模板信息为空，平台id：%s", client.Params.PlatId)
		}
		//取最后一个模板
		dto.TemplateUrl = templateRes.Results[templateLen-1].Url
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
		//fmt.Println(method, err)
		return res, body, err
	}
	res.WaybillApplyNewInfo = make([]*common.WaybillApplyNewInfo, 1)
	waybill := resJson.Get("results").GetIndex(0)
	waybillInfo := new(common.WaybillApplyNewInfo)
	isSuccess, _ := waybill.Get("issuccess").String()
	if isSuccess == "0" {
		// 不成功，返回错误信息
		message, _ := waybill.Get("message").String()
		return res, body, errors.New(message)
	}
	waybillInfo.WaybillCode, _ = waybill.Get("logisticno").String()
	waybillInfo.PackageCenterName, _ = waybill.Get("destcode").String()
	destName, _ := waybill.Get("destname").String()
	if destName != "" {
		waybillInfo.PackageCenterName = fmt.Sprintf("%s-%s", waybillInfo.PackageCenterName, destName)
	}

	waybillInfo.OriginName, _ = waybill.Get("originname").String()

	waybillInfo.OriginCrossCode, _ = waybill.Get("origincrosscode").String()
	originTableTrolleyCode, _ := waybill.Get("origintabletrolleycode").String()
	if originTableTrolleyCode != "" {
		waybillInfo.OriginCrossCode = fmt.Sprintf("%s-%s", waybillInfo.OriginCrossCode, originTableTrolleyCode)
	}

	waybillInfo.ShortAddress, _ = waybill.Get("markers").String()
	destCrossCode, _ := waybill.Get("destcrosscode").String()
	destTableTrolleyCode, _ := waybill.Get("desttabletrolleycode").String()
	if destCrossCode != "" || destTableTrolleyCode != "" {
		waybillInfo.ShortAddress = fmt.Sprintf("%s-%s-%s", waybillInfo.ShortAddress, destCrossCode, destTableTrolleyCode)
	}
	waybillInfo.ShortAddress = strings.Trim(waybillInfo.ShortAddress, "-")

	packages, hasPackages := waybill.CheckGet("packages")
	if hasPackages {
		pkg := packages.GetIndex(0)
		if pkg != nil {
			waybillInfo.PrintConfig, _ = pkg.Get("printinfo").String()
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
	//淘宝解密
	Oaid string `json:"oaid,omitempty"`
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
	CodPayMoney float64 `json:"codpaymoney,omitempty"`
	// 订单包裹物品金额
	PackageMoney string `json:"packagemoney,omitempty"`
	Weight       string `json:"weight,omitempty"`
	//保价金额
	SupporPayMoney float64 `json:"supporpaymoney,omitempty"`
	Length         string  `json:"length,omitempty"`
	Width          string  `json:"width,omitempty"`
	Height         string  `json:"height,omitempty"`
	Volume         string  `json:"volume,omitempty"`
	IsPickUp       string  `json:"ispickup,omitempty"`
	ProductType    string  `json:"producttype,omitempty"`
	LogisticType   string  `json:"logistictype,omitempty"`
	// 承运公司编码
	CpCode      string            `json:"cpcode,omitempty"`
	DmsSorting  string            `json:"dmssorting,omitempty"`
	NeedEncrypt uint8             `json:"needencrypt,omitempty"`
	Sender      *LogisticsAddress `json:"sender,omitempty"`
	Receiver    *LogisticsAddress `json:"receiver,omitempty"`

	Goods []*LogisticsGoods `json:"goods"`
	// 月结账号
	ShipperNo    string `json:"shipperno,omitempty"`
	WareCode     string `json:"warecode,omitempty"`
	BatchNo      string `json:"batchno,omitempty"`
	IsLiquid     string `json:"isliquid,omitempty"`
	IsHasBattery string `json:"ishasbattery,omitempty"`
	// 是否保险，默认0
	IsInsurance           uint8  `json:"isinsurance,omitempty"`
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

//菠萝派解密接口
type BusinessBatchTBDecryptBizcontent struct {
	Randomnumber string                          `json:"randomnumber"` //!必填	通用	64	淘宝随机字符串	tbxLGzL2r67me4zhYLHtDNvxxqPfjlgkAdU88pSPT55=
	Orders       []*BusinessBatchTBDecryptOrders `json:"orders"`       //!必填	通用	-	订单集合	-
	Ismask       json.Number                     `json:"ismask"`       //!必填	拼多多	64	是否需要调用脱敏解密接口(不需要=0，需要=...	1
}
type BusinessBatchTBDecryptOrders struct {
	Platorderno string                        `json:"platorderno"` //!必填	通用	32	平台订单号	TS04594434433
	Oaid        string                        `json:"oaid"`        //可选	通用	256	解密标识	2w2RYE45iahnF4aiaJ7pHKCJ3Hwnbgnq2PH3AfpQVyWZNHKS9wNgAAOUfCVt9XZMetogNHwc
	Items       []*BusinessBatchTBDecryptItem `json:"items"`       //!必填	通用	-	待加密密钥集合	-
}
type BusinessBatchTBDecryptItem struct {
	Secret string `json:"Secret"` //!必填	通用	256	待解密字符串	$136$rplMdffh4+x9GhkdlddxHg==$1$
	Type   string `json:"Type"`   //!必填	通用	256	密钥类别(收货人=RECEIVER_NAME，买家nick=...
}
type BusinessBatchTBDecryptReturn struct {
	Code             string                               `json:"code"`             //!必填	通用	64	返回码	10000
	Msg              string                               `json:"msg"`              //!必填	通用	64	返回消息	Success
	Subcode          string                               `json:"subcode"`          //!必填	通用	200	子集编码	LXO.JD.REQUEST_FAIL
	Submessage       string                               `json:"submessage"`       //!必填	通用	200	子级消息	订单已出库
	Polyapitotalms   json.Number                          `json:"polyapitotalms"`   //!必填	通用	64	菠萝派总耗时	102
	Polyapirequestid string                               `json:"polyapirequestid"` //!必填	通用	64	请求菠萝派编号	20161222154212742
	Orders           []*BusinessBatchTBDecryptReturnOrder `json:"orders"`
}
type BusinessBatchTBDecryptReturnOrder struct {
	Issuccess     json.Number                         `json:"issuccess"`   //!必填	通用	1	是否成功(0表示失败；1表示成功)	0
	Message       string                              `json:"message"`     //可选	通用	256	是否成功	订单已出库
	Platorderno   string                              `json:"platorderno"` //!必填	通用	32	平台订单号	TS04594434433
	Items         []*BusinessBatchTBDecryptReturnItem `json:"items"`       //必填	通用	-	待解密结果集合	-
	Addressdetail string                              `json:"addressdetail"`
	Area          string                              `json:"area"`
	City          string                              `json:"city"`
	Country       string                              `json:"country"`
	Messagefrom   string                              `json:"messagefrom"`
	Mobile        string                              `json:"mobile"`
	Name          string                              `json:"name"`
	Oaid          string                              `json:"oaid"`
	Phone         string                              `json:"phone"`
	Province      string                              `json:"province"`
	Subcode       string                              `json:"subcode"`
	Submessage    string                              `json:"submessage"`
	Town          string                              `json:"town"`
	Zip           string                              `json:"zip"`
}
type BusinessBatchTBDecryptReturnItem struct {
	Issuccess json.Number `json:"issuccess"` //！必填	通用	1	是否成功(0表示失败；1表示成功)	0
	Message   string      `json:"message"`   //可选	通用	256	是否成功	订单已出库
	Secret    string      `json:"secret"`    //!必填	通用	256	待解密字符串	$136$rplMdffh4+x9GhkdlddxHg==$1$
	Type      string      `json:"type"`      //!必填	通用	256	密钥类别(收货人=RECEIVER_NAME，买家nick=...
	Secretstr string      `json:"secretstr"` //!必填	通用	256	解密后字符串	15044558868
}
