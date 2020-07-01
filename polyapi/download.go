package polyapi

import (
	"fmt"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

// 产品下载
func (client *Client) DownloadProductList(pageIndex, pageSize int, status, productToken string, extData ...string) (res []*common.Product, hasNextPage bool, nextToken string, body []byte, err error) {
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
		reqJson.Set("shoptype", "JH_001")
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
	res = make([]*common.OrderInfo, 0)
	hasNextPage = false
	reqJson := simplejson.New()
	reqJson.Set("pageindex", pageIndex)
	reqJson.Set("pagesize", pageSize)
	reqJson.Set("starttime", startTime)
	reqJson.Set("endtime", endTime)
	reqJson.Set("timetype", timeType)
	reqJson.Set("orderstatus", orderStatus)
	reqJson.Set("nexttoken", orderToken)
	reqJson.Set("randomnumber", orderToken)

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
	orderList, err := resJson.Get("orders").Array()
	if err != nil {
		fmt.Println(method, err)
		return res, hasNextPage, nextToken, body, err
	}
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
	return res, hasNextPage, nextToken, body, err
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
