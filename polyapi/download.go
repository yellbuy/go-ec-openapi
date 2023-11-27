package polyapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/logs"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

// 产品下载
func (client *Client) DownloadProductList(pageIndex, pageSize int, status, productToken string, extData ...string) (res []*common.Product, hasNextPage bool, nextToken string, body []byte, err error) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err1 := recover(); err1 != nil {
			fmt.Println(err1) // 这里的err其实就是panic传入的内容，55
			//time.Sleep(time.Duration(30) * time.Second)
			err = errors.New("接口异常")
		}
	}()
	res = make([]*common.Product, 0)
	reqJson := simplejson.New()
	reqJson.Set("pageindex", pageIndex)
	reqJson.Set("pagesize", pageSize)
	reqJson.Set("status", status)
	reqJson.Set("requestid", productToken)

	if len(extData) > 0 {
		reqJson.Set("platvalue", extData[0])
	}
	if len(extData) > 1 {
		reqJson.Set("polyapitoken", extData[1])
	}
	if len(extData) > 2 {
		reqJson.Set("shoptype", extData[2])
	} else {
		if client.Params.PlatId == "13" {
			// 苏宁易购的参数
			reqJson.Set("shoptype", "0")
		} else {
			reqJson.Set("shoptype", "JH_001")
		}
	}
	reqJson.Set("producttype", "")
	bizcontent, resErr := reqJson.Encode()
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return res, hasNextPage, nextToken, body, err
	}

	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return res, hasNextPage, nextToken, body, err
	}

	// 通过奇门代理平台
	//method := "Differ.JH.Other.DelegateQimenDownloadProduct"
	// 通过polyapi自有平台
	method := "Differ.JH.Business.DownloadProduct"
	resJson := simplejson.New()
	resJson, body, err = client.Execute("Differ.JH.Business.DownloadProduct", params)
	if err != nil {
		fmt.Println(method, err)
		return res, hasNextPage, nextToken, body, err
	}
	nextToken, _ = resJson.Get("requestid").String()
	hasNextPageStr, _ := resJson.Get("ishasnextpage").String()
	hasNextPage = hasNextPageStr == "1"
	goodsList, err := resJson.Get("goodslist").Array()
	if err != nil {
		fmt.Println(method, err)
		return res, hasNextPage, nextToken, body, err
	}
	for index := range goodsList {
		goods := resJson.Get("goodslist").GetIndex(index)
		product := new(common.Product)
		product.ProductId, _ = goods.Get("platproductid").String()
		product.ProductName, _ = goods.Get("name").String()
		product.ProductCode, _ = goods.Get("outerid").String()
		product.Price, _ = goods.Get("price").String()
		product.Num, _ = goods.Get("num").String()
		product.WhseCode, _ = goods.Get("whsecode").String()
		product.Attrbutes, _ = goods.Get("attrbutes").String()
		product.CategoryId, _ = goods.Get("categoryid").String()
		product.Status, _ = goods.Get("status").String()
		product.PropertyAlias, _ = goods.Get("propertyalias").String()
		product.PictureUrl, _ = goods.Get("pictureurl").String()

		skuList, _ := goods.Get("skus").Array()
		product.SkuList = make([]*common.Sku, len(skuList))
		for j := range skuList {
			skuJson := goods.Get("skus").GetIndex(j)
			sku := new(common.Sku)
			sku.SkuId, _ = skuJson.Get("skuid").String()
			sku.SkuCode, _ = skuJson.Get("skuouterid").String()
			sku.SkuPrice, _ = skuJson.Get("skuprice").String()
			sku.SkuQuantity, _ = skuJson.Get("skuquantity").String()
			// if client.Params.PlatId == "1" {
			// 	// 淘宝平台，取商品的属性别名
			// 	sku.SkuName = product.PropertyAlias
			// } else {
			// 	sku.SkuName, _ = skuJson.Get("skuname").String()
			// }
			sku.SkuName, _ = skuJson.Get("skuname").String()
			sku.SkuProperty, _ = skuJson.Get("skuproperty").String()
			sku.SkuPictureUrl, _ = skuJson.Get("skupictureurl").String()
			sku.SkuName2, _ = skuJson.Get("skuname2").String()
			product.SkuList[j] = sku
		}
		res = append(res, product)
	}
	return res, hasNextPage, nextToken, body, err
}

// 订单下载
func (client *Client) DownloadOrderList(pageIndex, pageSize int, startTime, endTime, timeType, orderStatus, orderToken string, extData ...string) (res []*common.OrderInfo, hasNextPage bool, nextToken string, body []byte, err error) {
	defer func() {
		e := recover()
		if e == nil {
			return
		} else {
			err = fmt.Errorf("%s", e)
		}
	}()
	res = make([]*common.OrderInfo, 0)
	hasNextPage = false
	reqJson := simplejson.New()
	reqJson.Set("isuseinterface", true)
	reqJson.Set("pageindex", pageIndex)
	reqJson.Set("pagesize", pageSize)
	reqJson.Set("starttime", startTime)
	reqJson.Set("endtime", endTime)
	reqJson.Set("timetype", timeType)
	reqJson.Set("orderstatus", orderStatus)
	reqJson.Set("nexttoken", orderToken)
	reqJson.Set("randomnumber", orderToken)
	reqJson.Set("outUserName", "JXOS")

	if len(extData) > 0 && extData[0] != "" {
		reqJson.Set("platvalue", extData[0])
	}
	if len(extData) > 1 && extData[1] != "" {
		reqJson.Set("polyapitoken", extData[1])
	}
	if len(extData) > 2 && extData[2] != "" {
		reqJson.Set("shoptype", extData[2])
	} else {
		reqJson.Set("shoptype", "SOP")
	}
	bizcontent, resErr := reqJson.Encode()
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return res, hasNextPage, nextToken, body, err
	}

	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		return res, hasNextPage, nextToken, body, err
	}
	// 通过奇门代理平台
	//method := "Differ.JH.Other.DelegateQimenGetOrder"
	// 通过polyapi自有平台
	method := "Differ.JH.Business.GetOrder"
	resJson := simplejson.New()
	resJson, body, err = client.Execute(method, params)
	if err != nil {
		fmt.Println(method, err)
		return res, hasNextPage, nextToken, body, err
	}
	nextToken, _ = resJson.Get("nexttoken").String()
	hasNextPageStr, _ := resJson.Get("ishasnextpage").String()
	hasNextPage = hasNextPageStr == "1"
	logs.Debug("订单内容：", resJson)
	res, err = orderParse(resJson)
	return res, hasNextPage, nextToken, body, err
}

// 菠萝派订单解析，Key键值全小写
func orderParse(resJson *simplejson.Json) ([]*common.OrderInfo, error) {
	res := make([]*common.OrderInfo, 0)
	orderList, err := resJson.Get("orders").Array()
	if err != nil {
		fmt.Println("PolyApiOrderParse err:", err)
		return res, err
	}
	//fmt.Println("PolyApiOrderParse len:", len(orderList))
	for index := range orderList {
		order := resJson.Get("orders").GetIndex(index)
		orderInfo := new(common.OrderInfo)
		orderInfo.PlatOrderNo, _ = order.Get("platorderno").String()
		orderInfo.TradeStatus, _ = order.Get("tradestatus").String()
		orderInfo.Nick, _ = order.Get("nick").String()
		orderInfo.Mobile, _ = order.Get("mobile").String()
		orderInfo.Phone, _ = order.Get("phone").String()
		orderInfo.ReceiverName, _ = order.Get("receivername").String()
		orderInfo.Country, _ = order.Get("country").String()
		orderInfo.Province, _ = order.Get("province").String()
		orderInfo.City, _ = order.Get("city").String()
		orderInfo.Area, _ = order.Get("area").String()
		orderInfo.Town, _ = order.Get("town").String()
		orderInfo.Address, _ = order.Get("address").String()
		orderInfo.Zip, _ = order.Get("zip").String()
		orderInfo.CustomerRemark, _ = order.Get("customerremark").String()
		orderInfo.SellerRemark, _ = order.Get("sellerremark").String()
		orderInfo.ShipTypeName, _ = order.Get("shiptypename").String()
		orderInfo.PayOrderNo, _ = order.Get("payorderno").String()
		orderInfo.GoodsFee, _ = order.Get("goodsfee").String()
		orderInfo.TotalAmount, _ = order.Get("totalamount").String()
		orderInfo.PayTime, _ = order.Get("paytime").String()
		orderInfo.TradeTime, _ = order.Get("tradetime").String()

		goodsList, _ := order.Get("goodinfos").Array()
		orderInfo.GoodsInfoList = make([]*common.GoodsInfo, len(goodsList))
		for j := range goodsList {
			goodsJson := order.Get("goodinfos").GetIndex(j)
			goods := new(common.GoodsInfo)
			goods.SubOrderNo, _ = goodsJson.Get("suborderno").String()
			goods.PlatGoodsId, _ = goodsJson.Get("platgoodsid").String()
			goods.PlatSkuId, _ = goodsJson.Get("platskuid").String()
			goods.OutItemId, _ = goodsJson.Get("outitemid").String()
			goods.RefundStatus, _ = goodsJson.Get("refundstatus").String()
			goodsCount, err := goodsJson.Get("goodscount").String()
			if err != nil {
				goodsCountInt, _ := goodsJson.Get("goodscount").Int()
				goodsCount = strconv.Itoa(goodsCountInt)
			}
			goods.GoodsCount = goodsCount
			goods.TradeGoodsName, _ = goodsJson.Get("tradegoodsname").String()
			goods.TradeGoodsSpec, _ = goodsJson.Get("tradegoodsspec").String()
			goods.Price, _ = goodsJson.Get("price").String()
			goods.DiscountMoney, _ = goodsJson.Get("discountmoney").String()
			orderInfo.GoodsInfoList[j] = goods
		}
		res = append(res, orderInfo)
	}
	return res, nil
}

// 菠萝派订单解析
func PolyApiOrderParse(resJson *simplejson.Json) ([]*common.OrderInfo, error) {
	res := make([]*common.OrderInfo, 0)
	orderList, err := resJson.Get("Orders").Array()
	if err != nil {
		fmt.Println("PolyApiOrderParse err:", err)
		return res, err
	}
	//fmt.Println("PolyApiOrderParse len:", len(orderList))
	for index := range orderList {
		order := resJson.Get("Orders").GetIndex(index)
		orderInfo := new(common.OrderInfo)
		orderInfo.PlatOrderNo, _ = order.Get("PlatOrderNo").String()
		orderInfo.TradeStatus, _ = order.Get("TradeStatus").String()
		orderInfo.Nick, _ = order.Get("Nick").String()
		orderInfo.Mobile, _ = order.Get("Mobile").String()
		orderInfo.Phone, _ = order.Get("Phone").String()
		orderInfo.ReceiverName, _ = order.Get("ReceiverName").String()
		orderInfo.Country, _ = order.Get("Country").String()
		orderInfo.Province, _ = order.Get("Province").String()
		orderInfo.City, _ = order.Get("City").String()
		orderInfo.Area, _ = order.Get("Area").String()
		orderInfo.Town, _ = order.Get("Town").String()
		orderInfo.Address, _ = order.Get("Address").String()
		orderInfo.Zip, _ = order.Get("Zip").String()
		orderInfo.CustomerRemark, _ = order.Get("CustomerRemark").String()
		orderInfo.SellerRemark, _ = order.Get("SellerRemark").String()
		orderInfo.ShipTypeName, _ = order.Get("ShipTypeName").String()
		orderInfo.PayOrderNo, _ = order.Get("PayOrderNo").String()
		orderInfo.GoodsFee, _ = order.Get("GoodsFee").String()
		orderInfo.TotalAmount, _ = order.Get("TotalAmount").String()
		orderInfo.PayTime, _ = order.Get("PayTime").String()
		orderInfo.TradeTime, _ = order.Get("TradeTime").String()
		orderInfo.Oaid, _ = order.Get("Oaid").String()

		goodsList, _ := order.Get("GoodInfos").Array()
		orderInfo.GoodsInfoList = make([]*common.GoodsInfo, len(goodsList))
		for j := range goodsList {
			goodsJson := order.Get("GoodInfos").GetIndex(j)
			goods := new(common.GoodsInfo)
			goods.SubOrderNo, _ = goodsJson.Get("SubOrderNO").String()
			goods.PlatGoodsId, _ = goodsJson.Get("PlatGoodsID").String()
			goods.PlatSkuId, _ = goodsJson.Get("PlatSkuID").String()
			goods.OutItemId, _ = goodsJson.Get("OutItemId").String()
			goods.RefundStatus, _ = goodsJson.Get("RefundStatus").String()
			goodsCount, err := goodsJson.Get("GoodsCount").String()
			if err != nil {
				goodsCountInt, _ := goodsJson.Get("GoodsCount").Int()
				goodsCount = strconv.Itoa(goodsCountInt)
			}
			goods.GoodsCount = goodsCount
			goods.TradeGoodsName, _ = goodsJson.Get("TradeGoodsName").String()
			goods.TradeGoodsSpec, _ = goodsJson.Get("TradeGoodsSpec").String()
			goods.Price, _ = goodsJson.Get("Price").String()
			goods.DiscountMoney, _ = goodsJson.Get("DiscountMoney").String()
			orderInfo.GoodsInfoList[j] = goods
		}
		res = append(res, orderInfo)
	}
	return res, nil
}

func NewSuccessResDto(isSuccess bool, code int, message, itemId string) *SuccessResDto {
	dto := new(SuccessResDto)
	dto.Response = new(successRes)
	if isSuccess {
		dto.Response.Flag = "success"
	} else {
		dto.Response.Flag = "failure"
	}
	dto.Response.Code = code
	dto.Response.Message = message
	dto.Response.ItemId = itemId
	return dto
}

func NewErrorResDto(code int, message string, subCode int, subMsg string) *ErrorResDto {
	dto := new(ErrorResDto)
	dto.Response = new(errorRes)
	dto.Response.Code = code
	dto.Response.Message = message
	dto.Response.SubCode = subCode
	dto.Response.SubMsg = subMsg
	return dto
}
func (client *Client) PrintOrder(request *BatchPrintOrder_Order, extData ...string) (*PrintOrderReturn, error) {
	method := "Differ.JH.Logistics.PrintOrder" //菠萝派批量同步接口
	OutData := new(PrintOrderReturn)
	bizcontent, err := json.Marshal(request)
	if err != nil {
		return OutData, err
	}
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	if err != nil {
		return OutData, err
	}
	_, body, err := client.Execute(method, params)
	// logs.Debug("物流打印", string(body))
	if err != nil {
		if len(body) > 0 {
			json.Unmarshal(body, &OutData)
		}
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	if err != nil {
		logs.Debug("物流打印接口错误[" + err.Error() + "]")
	}
	return OutData, err
}
func (client *Client) BatchPrintOrder(request *BatchPrintOrder) (*LogisticBatchPrintOrderResponseResultItemInfo, error) {
	method := "Differ.JH.Logistics.BatchPrintOrder" //菠萝派批量同步接口
	OutData := new(LogisticBatchPrintOrderResponseResultItemInfo)
	bizcontent, err := json.Marshal(request)
	if err != nil {
		return OutData, err
	}
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	if err != nil {
		return OutData, err
	}
	_, body, err := client.Execute(method, params)
	logs.Debug("物流打印(批量)", string(body))
	if err != nil {
		if len(body) > 0 {
			json.Unmarshal(body, &OutData)
		}
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	if err != nil {
		logs.Debug("物流打印(批量)接口错误[" + err.Error() + "]")
	}
	return OutData, err
}

// Polyapi接口商品下载响应
type GoodsDownloadResponseDto struct {
	Ishasnextpage    string  `json:"ishasnextpage"`
	Totalcount       string  `json:"totalcount"`
	Goodslist        []Goods `json:"goodslist"`
	Requestid        string  `json:"requestid"`
	Code             string  `json:"code"`
	Msg              string  `json:"msg"`
	Subcode          string  `json:"subcode"`
	Submessage       string  `json:"submessage"`
	Polyapitotalms   string  `json:"polyapitotalms"`
	Polyapirequestid string  `json:"polyapirequestid"`
}
type Sku struct {
	Skuid         string `json:"skuid"`
	Skuouterid    string `json:"skuouterid"`
	Skuprice      string `json:"skuprice"`
	Skuquantity   string `json:"skuquantity"`
	Skuname       string `json:"skuname"`
	Skuproperty   string `json:"skuproperty"`
	Skutype       string `json:"skutype"`
	Skupictureurl string `json:"skupictureurl"`
	Skuname2      string `json:"skuname2"`
}
type Goods struct {
	PlatProductId      string      `json:"platproductid"`
	Name               string      `json:"name"`
	Outerid            string      `json:"outerid"`
	Price              string      `json:"price"`
	Num                string      `json:"num"`
	Pictureurl         string      `json:"pictureurl"`
	Whsecode           string      `json:"whsecode"`
	Attrbutes          interface{} `json:"attrbutes"`
	Categoryid         string      `json:"categoryid"`
	Status             string      `json:"status"`
	Statusdesc         string      `json:"statusdesc"`
	SkuList            []Sku       `json:"skus"`
	SendType           string      `json:"sendtype"`
	Skutype            string      `json:"skutype"`
	Propertyalias      string      `json:"propertyalias"`
	Isplatstorageorder string      `json:"isplatstorageorder"`
	Cooperationno      string      `json:"cooperationno"`
}

// 奇门接口成功响应
type SuccessResDto struct {
	Response *successRes `json:"response"`
}

// 奇门接口错误响应
type ErrorResDto struct {
	Response *errorRes `json:"error_response"`
}

// 奇门接口下载成功响应内容
type successRes struct {
	//区名称（三级地址）
	Flag    string `json:"flag"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	ItemId  string `json:"itemId"`
}
type errorRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	SubCode int    `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

type SubGoods struct {
	ProductId      string `json:"productid"`
	TradeGoodsNo   string `json:"tradegoodsno"`
	TradeGoodsName string `json:"tradegoodsname"`
	GoodsCount     string `json:"goodscount"`
	Price          string `json:"price"`
	OutSkuId       string `json:"outskuid"`
	OutTtemId      string `json:"outitemid"`
	PlatGoodsId    string `json:"platgoodsid"`
	PlatSkuId      string `json:"platskuid"`
}
type GoodInfo struct {
	ProductId             string      `json:"productid"`
	SubOrderNo            string      `json:"suborderno"`
	TaxSubOrderNo         string      `json:"taxsuborderno"`
	TradeGoodsNo          string      `json:"tradegoodsno"`
	TradeGoodsName        string      `json:"tradegoodsname"`
	TradeGoodsSpec        string      `json:"tradegoodsspec"`
	GoodsCount            string      `json:"goodscount"`
	Price                 string      `json:"price"`
	RefundCount           string      `json:"refundcount"`
	DiscountMoney         string      `json:"discountmoney"`
	TaxAmount             string      `json:"taxamount"`
	TariffAmount          string      `json:"tariffamount"`
	AddedValueAmount      string      `json:"addedvalueamount"`
	ConsumptionDutyAmount string      `json:"consumptiondutyamount"`
	RefundStatus          string      `json:"refundstatus"`
	Status                string      `json:"status"`
	Remark                string      `json:"remark"`
	OutSkuId              string      `json:"outskuid"`
	PlatGoodsId           string      `json:"platgoodsid"`
	PlatSkuId             string      `json:"platskuid"`
	OutItemId             string      `json:"outitemid"`
	SubGoodsList          []*SubGoods `json:"subgoods"`
	IsGift                string      `json:"isgift"`
	IsHwgFlag             string      `json:"ishwgflag"`
	DeliveryType          string      `json:"deliverytype"`
	PayorderId            string      `json:"payorderid"`
	PackageOrderId        string      `json:"packageorderid"`
	ActivityAmount        string      `json:"activityamount"`
	SpecialAmount         string      `json:"specialamount"`
	CouponAmount          string      `json:"couponamount"`
	ProductItemId         string      `json:"productitemid"`
	GoodsCount2           string      `json:"goodscount2"`
	IsPlatStorageOrder    string      `json:"isplatstorageorder"`
	PictureUrl            string      `json:"pictureurl"`
	GoodType              string      `json:"goodtype"`
	EstimateConTime       string      `json:"estimatecontime"`
	Fenxiaoprice          string      `json:"fenxiaoprice"`
	SubOrderItemNo        string      `json:"suborderitemno"`
	GoodsOrderAttr        string      `json:"goodsorderattr"`
}
type CouponDetail struct {
	SkuId      string `json:"sku_id"`
	CouponType string `json:"coupontype"`
	Type       string `json:"type"`
	Price      string `json:"price"`
	CouponNum  string `json:"couponnum"`
}
type ServiceOrder struct {
	ServiceId    string `json:"serviceid"`
	ServiceName  string `json:"servicename"`
	ServiceType  string `json:"servicetype"`
	ServicePrice string `json:"serviceprice"`
	ServiceNum   string `json:"servicenum"`
}
type ShopOrderInfo struct {
	ShopType                     string          `json:"shoptype"`
	PlatOrderNo                  string          `json:"platorderno"`
	TradeStatus                  string          `json:"tradestatus"`
	TradeStatusDescription       string          `json:"tradestatusdescription"`
	TradeTime                    string          `json:"tradetime"`
	ModifyTime                   string          `json:"modifytime"`
	CollageTime                  string          `json:"collagetime"`
	Username                     string          `json:"username"`
	Nick                         string          `json:"nick"`
	BuyerMobile                  string          `json:"buyermobile"`
	ReceiverName                 string          `json:"receivername"`
	Country                      string          `json:"country"`
	Province                     string          `json:"province"`
	City                         string          `json:"city"`
	Area                         string          `json:"area"`
	Town                         string          `json:"town"`
	Address                      string          `json:"address"`
	PayOrderNo                   string          `json:"payorderno"`
	PayType                      string          `json:"paytype"`
	ShouldPayType                string          `json:"shouldpaytype"`
	Zip                          string          `json:"zip"`
	Phone                        string          `json:"phone"`
	Mobile                       string          `json:"mobile"`
	Email                        string          `json:"email"`
	CustomeRremark               string          `json:"customerremark"`
	SellerRemark                 string          `json:"sellerremark"`
	PostFee                      string          `json:"postfee"`
	PostInsuranceFee             string          `json:"postinsurancefee"`
	GoodsFee                     string          `json:"goodsfee"`
	TotalAmount                  string          `json:"totalamount"`
	RealPayMoney                 string          `json:"realpaymoney"`
	FavourableMoney              string          `json:"favourablemoney"`
	PlatDiscountMoney            string          `json:"platdiscountmoney"`
	TaxAmount                    string          `json:"taxamount"`
	TariffAmount                 string          `json:"tariffamount"`
	AddedValueAmount             string          `json:"addedvalueamount"`
	ConsumptionDutyAmount        string          `json:"consumptiondutyamount"`
	CommissionValue              string          `json:"commissionvalue"`
	PayTime                      string          `json:"paytime"`
	SendType                     string          `json:"sendtype"`
	SendStyle                    string          `json:"sendstyle"`
	CodServiceFee                string          `json:"codservicefee"`
	SellerFlag                   string          `json:"sellerflag"`
	CardType                     string          `json:"cardtype"`
	IdCard                       string          `json:"idcard"`
	IdCardTrueName               string          `json:"idcardtruename"`
	IdCardImgs                   string          `json:"idcardimgs"`
	WhseCode                     string          `json:"whsecode"`
	IsHwgFlag                    string          `json:"ishwgflag"`
	DeliveryType                 string          `json:"deliverytype"`
	ShopId                       string          `json:"shopid"`
	MdbId                        string          `json:"mdbid"`
	SaleSpin                     string          `json:"salespin"`
	IsNeedInvoice                string          `json:"isneedinvoice"`
	InvoiceType                  string          `json:"invoicetype"`
	InvoiceBusinessType          string          `json:"invoicebusinesstype"`
	InvoiceTitle                 string          `json:"invoicetitle"`
	InvoiceContent               string          `json:"invoicecontent"`
	TaxPayerIdent                string          `json:"taxpayerident"`
	RegisteredAddress            string          `json:"registeredaddress"`
	RegisteredPhone              string          `json:"registeredphone"`
	DepositBank                  string          `json:"depositbank"`
	BankAccount                  string          `json:"bankaccount"`
	FetchTime                    string          `json:"fetchtime"`
	FetchTimeDesc                string          `json:"fetchtimedesc"`
	OrderSource                  string          `json:"ordersource"`
	CustomAttr                   string          `json:"customattr"`
	TransportDay                 string          `json:"transportday"`
	GoodInfoList                 []*GoodInfo     `json:"goodinfos"`
	CouponDetails                []*CouponDetail `json:"coupondetails"`
	CustomsCode                  string          `json:"customscode"`
	UserLevel                    string          `json:"userlevel"`
	ReminderCount                string          `json:"remindercount"`
	VerifyCode                   string          `json:"verifycode"`
	SellerOrderId                string          `json:"sellerorderid"`
	Qq                           string          `json:"qq"`
	OrderFlag                    string          `json:"orderflag"`
	TradeAttr                    string          `json:"tradeattr"`
	OrdertType                   string          `json:"ordertype"`
	IsStoreOrder                 string          `json:"isstoreorder"`
	IsYunStoreOrder              string          `json:"isyunstoreorder"`
	SendDate                     string          `json:"senddate"`
	LeftSendDate                 string          `json:"leftsenddate"`
	SortingCode                  string          `json:"sortingcode"`
	PartNer                      string          `json:"partner"`
	BfDeligoodglag               string          `json:"bfdeligoodglag"`
	CkyName                      string          `json:"ckyname"`
	CreateDate                   string          `json:"createdate"`
	ShouldPayMoney               string          `json:"shouldpaymoney"`
	ResellerId                   string          `json:"resellerid"`
	ResellerShopName             string          `json:"resellershopname"`
	ResellerMobile               string          `json:"resellermobile"`
	ActivityAmount               string          `json:"activityamount"`
	BalanceUsed                  string          `json:"balanceused"`
	SpecialAmount                string          `json:"specialamount"`
	CouponAmount                 string          `json:"couponamount"`
	CurrencyCode                 string          `json:"currencycode"`
	IsPresaleOrder               string          `json:"ispresaleorder"`
	ShipTypeName                 string          `json:"shiptypename"`
	BondedId                     string          `json:"bondedid"`
	BondedName                   string          `json:"bondedname"`
	LogisticNo                   string          `json:"logisticno"`
	LogisticName                 string          `json:"logisticname"`
	PayMethod                    string          `json:"paymethod"`
	FxtId                        string          `json:"fxtid"`
	ClerkName                    string          `json:"clerkname"`
	ClerkPhone                   string          `json:"clerkphone"`
	TransactionId                string          `json:"transactionid"`
	IsDaiXiao                    string          `json:"isdaixiao"`
	IsBrandSale                  string          `json:"isbrandsale"`
	IsForcewLb                   string          `json:"isforcewlb"`
	IsWlBorder                   string          `json:"iswlborder"`
	IsShip                       string          `json:"isship"`
	IsEncrypt                    string          `json:"isencrypt"`
	FulfillmentChannel           string          `json:"fulfillmentchannel"`
	ShipmentServiceLevelCategory string          `json:"shipmentservicelevelcategory"`
	IsshippedByAmazonTfm         string          `json:"isshippedbyamazontfm"`
	IsJzOrder                    string          `json:"isjzorder"`
	TradeAttrJson                string          `json:"tradeattrjson"`
	EndTime                      string          `json:"endtime"`
	ProductNum                   string          `json:"productnum"`
	CancelTime                   string          `json:"canceltime"`
	CancelReason                 string          `json:"cancelreason"`
	AuditResult                  string          `json:"auditresult"`
	AuditReason                  string          `json:"auditreason"`
	OrderNumber                  string          `json:"ordernumber"`
	InnertranSactionId           string          `json:"innertransactionid"`
	OriginalMobile               string          `json:"originalmobile"`
	ServiceOrderList             []*ServiceOrder `json:"serviceorders"`
}
type PrintOrderReturn struct {
	Code             string                                     `json:"code"`             //1 返回码 10000
	Msg              string                                     `json:"msg"`              //1 返回消息 Success
	SubCode          string                                     `json:"subCode"`          //1 返回子码
	SubMessage       string                                     `json:"subMessage"`       //1 返回子级消息
	PolyAPITotalMS   int                                        `json:"polyAPITotalMS"`   //1 请求接口总耗时(毫秒) 287
	PolyAPIRequestID string                                     `json:"polyAPIRequestID"` //1 请求菠萝派上下文编号 2016122913161617168574
	MessageFrom      string                                     `json:"messageFrom"`      //1指示此次返回是菠萝派的返回还是平台的返回 默认 POLY :POLY=POLY,PLAT=PLAT POLY
	Env              string                                     `json:"env"`              //1 指示此次返回结果来源于菠萝派哪个环境 1
	BufferDataType   string                                     `json:"bufferDataType"`   //1 二进制内容类别:PDF=0,GIF=0,JPG=0,PNG=0 0
	BufferData       string                                     `json:"bufferData"`       //1 打印内容(二进制) -
	UrlData          string                                     `json:"urlData"`          //1 打印内容网址 http://www.aaa.com/stream.pdf
	PrintData        string                                     `json:"printData"`        //0 打印内容 万邑通ISP
	HtmlData         string                                     `json:"htmlData"`         //0 打印内容 唯品会JIT物流、tokopedia
	LogisticNO       string                                     `json:"logisticNO"`       //0 运单号 天猫海外仓直购物流
	BatchPrintData   []*BatchPrintOrder_Return_BATCH_PRINT_DATA `json:"batchPrintData"`   //1 打印多个面单 batchPrintData[]
	TrackingId       string                                     `json:"trackingId"`       //0 跟踪id 亚马逊物流
	EncryptVersion   string                                     `json:"encryptVersion"`   //1 加密算法版本号 华为电子面单
}
type LogisticBatchPrintOrderResponseResultItemInfo struct {
	IsSuccess            bool                                       `json:"isSuccess"`              //1 是否成功(0表示失败；1表示成功) 0
	Code                 string                                     `json:"code"`                   //1 错误编码 成功-10000 40000
	Message              string                                     `json:"message"`                //0 失败说明 订单已出库
	SubCode              string                                     `json:"subCode"`                //1 返回子码
	SubMessage           string                                     `json:"subMessage"`             //1 返回子级消息
	MessageFrom          string                                     `json:"messageFrom"`            //1 指示此次返回是菠萝派的返回还是平台的返回 默认 POLY :POLY=POLY,PLAT=PLAT POLY
	BufferDataType       string                                     `json:"bufferDataType"`         //1 二进制内容类别:PDF=0,GIF=0,JPG=0,PNG=00
	BufferData           string                                     `json:"bufferData"`             //1 打印内容(二进制) -
	OrderNo              string                                     `json:"orderNo"`                //1 客户订单号 12313123123
	LogisticNo           string                                     `json:"logisticNo"`             //1 运单号(多个用英文逗号隔开) 12313123123
	UrlData              string                                     `json:"urlData"`                //1 打印内容网址 http://www.aaa.com/stream.pdf
	PrintData            string                                     `json:"printData"`              //0 打印内容 万邑通ISP
	BaseStringBufferData string                                     `json:"base64StringBufferData"` //0 打印内容(二进制)转Base64编码
	HtmlData             string                                     `json:"htmlData"`               //0 打印内容 唯品会JIT物流、tokopedia
	BatchPrintData       []*BatchPrintOrder_Return_BATCH_PRINT_DATA `json:"batchPrintData"`         //1 打印多个面单 batchPrintData[]
	TrackingId           string                                     `json:"trackingId"`             //0 跟踪id 亚马逊物流
	Token                string                                     `json:"token"`                  //0 token(顺丰丰桥专用) 顺丰丰桥
	EncryptVersion       string                                     `json:"encryptVersion"`         //1 加密算法版本号 华为电子面单
}
type BatchPrintOrder_Return_BATCH_PRINT_DATA struct {
	UrlData    string `json:"urlData"`    //1 打印内容网址 http://www.aaa.com/stream.pdf
	OrderNo    string `json:"orderNo"`    //1 订单号 FO58965555996
	HtmlData   string `json:"htmlData"`   //0 打印内容 tokopedia
	LogisticNo string `json:"logisticNo"` //1 物流单号 SFO58965555996
	ExtraInfo  string `json:"extraInfo"`  //1 拓展字段 123456789 顺丰丰桥
}
type BatchPrintOrder struct {
	Orders []*BatchPrintOrder_Order `json:"orders"`
}
type BatchPrintOrder_Order struct {
	OrderNo            string `json:"orderNo"`            //1 客户订单号 T20160922120
	LogisticNo         string `json:"logisticNo"`         //1 运单号 Y20160927770
	OutputFormat       string `json:"outputFormat"`       //1 打印输出格式:PDF格式=JH_PDF,PNG格式=JH_PNG,PDF格式的网址=JH_PDFUrl JH_PDF
	LabelFormat        string `json:"labelFormat"`        //1 打印尺寸格式:PDF格式=JH_A4,PDF格式的网址=JH_Label,PDF格式=JH_LABEL10_10,PDF格式=JH_LABEL10_15 JH_A4
	PrintObjectType    string `json:"printObjectType"`    //1 打印类别:标签=JH_Label,标签+配货=JH_LabelPH,标签+报关=JH_LabelBG,标签+配货+报关=JH_LabelPHBG,收据联=JH_Receiving,发货联=JH_Shipping,收据联和发货联=JH_ReceivingAndShipping JH_LABEL
	NumPackage         string `json:"numPackage"`         //1 包裹数量 1
	CountryCode        string `json:"countryCode"`        //1 国家二字简码(去发货专用) US 去发货
	LogisticsId        string `json:"logisticsId"`        //1 渠道编码，去发货系统的渠道编号(去发货专用) WELPMDWY396 去发货
	OrderItemNo        string `json:"orderItemNo"`        //1 商品编码 Y20160927770 LaZaDa
	SenderName         string `json:"senderName"`         //1 寄件人姓名 张三 申通快递
	SenderCompany      string `json:"senderCompany"`      //1 寄件人公司名称 笛佛网店管家 申通快递
	SenderProvince     string `json:"senderProvince"`     //1 寄件人省份 浙江省 申通快递
	SenderCity         string `json:"senderCity"`         //1 寄件人城市 杭州 申通快递
	SenderArea         string `json:"senderArea"`         //1 寄件人区/县 西湖区 申通快递
	SenderAddress      string `json:"senderAddress"`      //1 寄件人地址 紫金花路120号 申通快递
	SenderTel          string `json:"senderTel"`          //1 寄件人电话 15044444444 申通快递
	ReceiverName       string `json:"receiverName"`       //1 收件人姓名 李四 申通快递
	ReceiverCompany    string `json:"receiverCompany"`    //1 收件人公司名称 上海东商有限公司 申通快递
	ReceiverProvince   string `json:"receiverProvince"`   //1 收件人省份 上海 申通快递
	ReceiverCity       string `json:"receiverCity"`       //1 收件人城市 上海 申通快递
	ReceiverArea       string `json:"receiverArea"`       //1 收件人区/县 浦东新区 申通快递
	ReceiverAddress    string `json:"receiverAddress"`    //1 收件人地址 浦东156号 申通快递
	ReceiverTel        string `json:"receiverTel"`        //1 收件人电话 13587944872 申通快递
	CargoName          string `json:"cargoName"`          //1 货物内件品名 茶杯 申通快递
	SendData           string `json:"sendData"`           //1 发货日期 2016-7-01 12:34:01 申通快递
	Weight             string `json:"weight"`             //1 重量 620 申通快递
	Remark             string `json:"remark"`             //1 备注 轻拿轻放 通快递
	Carrier            string `json:"carrier"`            //1 运输公司 联邦 Ebay物流
	SellerUserId       string `json:"sellerUserId"`       //1 卖家eBay账户 u12451547 Ebay物流
	PrintNo            string `json:"printNo"`            //1 打印单号 154643265 万邑通ISP
	IsCloudPrint       bool   `json:"isCloudPrint"`       //1 是否使用云打印 true 速卖通物流
	NeedPrintDetail    bool   `json:"needPrintDetail"`    //1 是否需要打印详情 true 速卖通物流、飞裕达
	ExtendData         string `json:"extendData"`         //1 自定义分拣单信息 154643265 速卖通物流
	ProviderCode       string `json:"providerCode"`       //1 承运商编码 20160927770 唯品会JIT物流
	PackageNo          string `json:"packageNo"`          //1 包裹序号 20160927770 唯品会JIT物流
	FontSize           string `json:"fontSize"`           //1 面单字体大小 (澳邮物流专用) 0 澳邮物流
	PlatOrderNo        string `json:"PlatOrderNo"`        //1 平台订单号(京东印尼专用) 0 京东印尼
	LogisticChildNO    string `json:"logisticChildNO"`    //1 子运单号 SF2234567890100
	BackSignBillNo     string `json:"backSignBillNo"`     //0 签回单号 SF2234567890100
	NeedPrintLogo      bool   `json:"needPrintLogo"`      //0 是否打印LOGO true
	CustomTemplateCode string `json:"customTemplateCode"` //0 自定义模板编码 123 顺丰丰桥
	CustomTemplateData string `json:"customTemplateData"` //0 自定义模板 123 顺丰丰桥
	SubOrderNo         string `json:"subOrderNo"`         //0 子订单号/物流单号 12345 华为电子面单
	SIPShopId          string `json:"SIPShopId"`          //0 跨境一店通业务下的子店铺ID 12345 TikTok物流
	ShopType           string `json:"shopType"`           //0  店铺类型 JH_MIRAVIA 速卖通物流
	BusinessType       string `json:"businessType"`       //1 业务类型 IUOP 顺丰国际IBS
	PrintPicking       bool   `json:"printPicking"`       //1 是否打印拣货单 false 顺丰国际IBS
	IsPrintSubParent   string `json:"isPrintSubParent"`   //1 是否打印全部的子母单 false 顺丰国际IBS
	TemplateId         string `json:"templateId"`         //1 模板id 12312313
	PrintNum           int    `json:"printNum"`           //1 打印数量 2
}
