package common

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

type SuccessResDto struct {
	Response *successRes `json:"response"`
}

type ErrorResDto struct {
	Response *errorRes `json:"error_response"`
}

// 奇门下载成功响应内容
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

// 商品
type Product struct {
	ProductId          string      `json:"platproductid"`
	ProductName        string      `json:"name"`
	ProductCode        string      `json:"outerid"`
	Price              string      `json:"price"`
	Num                string      `json:"num"`
	PictureUrl         string      `json:"pictureurl"`
	WhseCode           string      `json:"whsecode"`
	Attrbutes          interface{} `json:"attrbutes"`
	CategoryId         string      `json:"categoryid"`
	Status             string      `json:"status"`
	StatusDesc         string      `json:"statusdesc"`
	SkuList            []*Sku      `json:"skus"`
	SendType           string      `json:"sendtype"`
	SkuType            string      `json:"skutype"`
	PropertyAlias      string      `json:"propertyalias"`
	IsPlatStorageOrder string      `json:"isplatstorageorder"`
	CooperationNo      string      `json:"cooperationno"`
}

// 规格
type Sku struct {
	SkuId         string `json:"skuid"`
	SkuCode       string `json:"skuouterid"`
	SkuPrice      string `json:"skuprice"`
	SkuQuantity   string `json:"skuquantity"`
	SkuName       string `json:"skuname"`
	SkuProperty   string `json:"skuproperty"`
	SkuType       string `json:"skutype"`
	SkuPictureUrl string `json:"skupictureurl"`
	SkuName2      string `json:"skuname2"`
}

type SubGoods struct {
	Productid      string `json:"productid"`
	Tradegoodsno   string `json:"tradegoodsno"`
	Tradegoodsname string `json:"tradegoodsname"`
	Goodscount     string `json:"goodscount"`
	Price          string `json:"price"`
	Outskuid       string `json:"outskuid"`
	Outitemid      string `json:"outitemid"`
	Platgoodsid    string `json:"platgoodsid"`
	Platskuid      string `json:"platskuid"`
}
type GoodsInfo struct {
	// 子订单号
	SubOrderNo  string `json:"suborderno"`
	PlatGoodsId string `json:"platgoodsid"`
	PlatSkuId   string `json:"platskuid"`
	ProductId   string `json:"productid"`
	//商品数量
	GoodsCount string `json:"goodscount"`
	//商品原价
	Price string `json:"price"`
	// 线上商品交易名字
	TradeGoodsNo string `json:"tradegoodsno"`
	//线上商品交易名字
	TradeGoodsName string `json:"tradegoodsname"`
	// 线上商品交易 规格 名字
	TradeGoodsSpec string `json:"tradegoodsspec"`
	// 商品退款状态（JH_07没有退款，返回不是这个状态的货品货品显示明面加个退字）
	RefundStatus string `json:"refundstatus"`
	RefundCount  string `json:"refundcount"`
	// 商品优惠金额
	DiscountMoney string `json:"discountmoney"`

	Status   string `json:"status"`
	Remark   string `json:"remark"`
	OutSkuId string `json:"outskuid"`

	OutItemId    string      `json:"outitemid"`
	SubGoodsList []*SubGoods `json:"subgoods"`
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
type OrderInfo struct {
	ShopType string `json:"shoptype"`
	//平台订单号   对应ERP原始订单号
	PlatOrderNo string `json:"platorderno"`
	// 订单交易状态  JH_02  等待卖家发货只下载此状态订单
	TradeStatus            string `json:"tradestatus"`
	TradeStatusDescription string `json:"tradestatusdescription"`
	TradeTime              string `json:"tradetime"`
	ModifyTime             string `json:"modifytime"`
	CollageTime            string `json:"collagetime"`
	Username               string `json:"username"`
	Nick                   string `json:"nick"`
	Zip                    string `json:"zip"`
	Phone                  string `json:"phone"`
	Mobile                 string `json:"mobile"`
	Email                  string `json:"email"`
	//买家备注
	CustomerRemark string `json:"customeremark"`
	//卖家备注
	SellerRemark string `json:"sellerremark"`
	BuyerMobile  string `json:"buyermobile"`
	// 收件人姓名
	ReceiverName string `json:"receivername"`
	Country      string `json:"country"`
	Province     string `json:"province"`
	City         string `json:"city"`
	// 区县
	Area string `json:"area"`
	Town string `json:"town"`
	//详细地址
	Address string `json:"address"`
	// 支付单号
	PayOrderNo    string `json:"payorderno"`
	PayType       string `json:"paytype"`
	ShouldPayType string `json:"shouldpaytype"`
	// 支付单号
	TotalAmount string `json:"totalamount"`
	// 实付金额
	RealPayMoney string `json:"realpaymoney"`
	// 邮费
	PostFee string `json:"postfee"`
	// 保险费
	PostInsuranceFee string `json:"postinsurancefee"`
	// 物品费用
	GoodsFee string `json:"goodsfee"`
	// 支付时间
	PayTime string `json:"paytime"`
	GoodsInfoList []*GoodsInfo `json:"goodinfos"`

	ServiceOrderList []*ServiceOrder `json:"serviceorders"`
}
