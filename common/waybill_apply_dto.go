package common

//电子面单模板请求dto
type WaybillTemplateRequest struct {
	//订单信息(所有模版=ALL，客户拥有的模版=OWNER)
	TemplatesType string `json:"templatestype"`
	// 快递公司类别
	LogisticType string `valid:"Required" json:"logistictype"`
}

// 电子面单查询结果dto
type WaybillTemplateDto struct {
	// 承运公司编码
	CpCode  string                 `json:"cpcode"`
	Results []*WaybillTemplateInfo `json:"results"`
}

// 电子面单查询结果信息
type WaybillTemplateInfo struct {
	// 模板id
	Id string `json:"id"`
	// 模板名称
	Name string `json:"name"`
	// 模板url
	Url string `json:"url"`
	// 模板url
	TemplateType string `json:"templatetype"`
}

// 收\发货地址
type WaybillAddress struct {
	//区名称（三级地址）
	Area    string `json:"area"`
	Country string `json:"country"`
	// 省名称，必填（一级地址）
	Province string `valid:"Required" json:"province"`
	// 街道\镇名称（四级地址）
	Town string `json:"town"`
	// 详细地址，必填
	AddressDetail string `valid:"Required" json:"address_detail"`
	// 市名称（二级地址）
	City string `json:"city"`
}

// 订单数据
type TradeOrderInfo struct {
	// 请求参数

	// 收货人，必填
	ConsigneeName string `valid:"Required" json:"consignee_name"`
	// 订单渠道:TB，必填
	OrderChannelsType string `valid:"Required" json:"order_channels_type"`
	// 交易订单列表，必填:12321321,12321321s
	TradeOrderList []string `valid:"Required" json:"trade_order_list"`
	// 收\发货地址，必填
	ConsigneeAddress *WaybillAddress `valid:"Required" json:"consignee_address"`
	// 发货人联系方式
	SendPhone string `json:"send_phone"`
	// 发货人姓名
	SendName string `json:"send_name"`
	// 重量，单位：克
	Weight int `json:"weight"`
	// 包裹里面的商品名称，必填
	PackageItems []*PackageItem `valid:"Required" json:"package_items"`
	// 物流服务能力集合
	LogisticsServiceList []*LogisticsService `json:"logistics_service_list"`
	// 快递服务产品类型编码，必填:STANDARD_EXPRESS
	ProductType string `valid:"Required" json:"product_type"`
	// 使用者ID，必填:13123
	RealUserId int64 `valid:"Required" json:"real_user_id"`
	// 包裹体积（立方厘米）
	Volumn int64 `json:"volume"`
	// 包裹号(或者ERP订单号)
	PackageId string `json:"package_id"`

	// 额外的响应参数
	// 商品名称
	ItemName string `json:"item_name"`
	// 是否阿里系订单
	AliOrder bool `json:"ali_order"`
	// 	大头笔:余杭
	ShortAddress string `json:"short_address"`
	// 面单号
	WaybillCode string `json:"waybill_code"`
	// 收货人联系方式
	ConsigneePhone string `json:"consignee_phone"`
	// 订单渠道来源:123
	OrderType string `json:"order_type"`
	// 菠萝派货到付款金额,OrderType=JH_COD时必传
	CodPayMoney float64 `json:"codpaymoney,omitempty"`
	// 拼多多接口所需参数
	ObjectId string `json:"-"`

	// 菠萝派接口所需参数
	CustomerCode string `json:"-"`
	CustomerName string `json:"-"`
	OrderNo      string `json:"-"`
	PlatTradeNo  string `json:"-"`
	SiteCode     string `json:"-"`
	// 是否保价
	IsInsurance uint8 `json:"-"`
	//保价金额
	SupporPayMoney float64 `json:"-"`
	// 销售平台
	BusinessPlat string `json:"-"`
}

// 包裹里面的商品名称
type PackageItem struct {
	// 商品名称，必填
	ItemName string `valid:"Required" json:"item_name"`
	//商品数量，必填
	Count int `valid:"Required" json:"count"`
}

// 物流服务能力集合
type LogisticsService struct {
	// 服务类型值，json格式表示:{ "value": "100.00","currency": "CNY","ensure_type": "0"}
	ServiceValue4Json string `json:"service_value4_json"`
	// 	服务编码:SVC-DELIVERY-ENV
	ServiceCode string `json:"service_code"`
}

// 电子面单请求结构体
type WaybillApplyNewRequest struct {
	// 物流服务商编码，必填
	CpCode string `valid:"Required" json:"cp_code"`
	// 菠萝派业务类型
	BusinessType uint `json:"business_type"`

	// 菠萝派月结帐号
	ShipperNo string `json:"shipperno"`
	// 菠萝派支付方式
	PayMode uint8 `json:"pay_mode"`
	// 收\发货地址，必填
	ShippingAddress *WaybillAddress `valid:"Required" json:"shipping_address"`
	// 订单数据，必填
	TradeOrderInfoCols []*TradeOrderInfo `valid:"Required" json:"trade_order_info_cols"`
}

// 电子面单响应结构体
// taobao.wlb.waybill.i.get
type WaybillApplyNewCols struct {
	WaybillApplyNewInfo []*WaybillApplyNewInfo `json:"waybill_apply_new_info"`
}

// 电子面单响应结构体
type WaybillApplyDto struct {
	// 根据收货地址返回大头笔信息
	ShortAddress string `json:"short_address"`
	// 返回的面单号
	WaybillCode string `json:"waybill_code"`
	// 集包地代码
	PackageCenterCode string `json:"package_center_code"`
	// 集包地名称
	PackageCenterName string `json:"package_center_name"`
	// 物流始发站点
	OriginName string `json:"origin_name"`
	// 始发网点编码
	OriginCrossCode string `json:"origin_cross_code"`
	// 打印配置项，传给ali-print组件
	PrintConfig string `json:"print_config"`
	// 面单号对应的物流服务商网点（分支机构）代码
	ShippingBranchCode string `json:"shipping_branch_code"`
	// 包裹对应的派件（收件）物流服务商网点（分支机构）名称:余杭一部
	ConsigneeBranchName string `json:"consignee_branch_name"`
	// 面单号对于的物流服务商网点（分支机构）名称:西湖二部
	ShippingBranchName string `json:"shipping_branch_name"`
	// 包裹对应的派件（收件）物流服务商网点（分支机构）代码
	ConsigneeBranchCode string `json:"consignee_branch_code"`
}
type WaybillApplyNewInfo struct {
	WaybillApplyDto
	// 订单数据，必填
	TradeOrderInfoCols []*TradeOrderInfo `valid:"Required" json:"trade_order_info_cols"`
	// 兼容pdd的接口，电子面单内容
	PrintData string `json:"-"`
}

// type CommonResponseDto{
// 	//平台颁发的每次请求访问的唯一标识
// 	RequestId	string	`json:"request_id"`
// 	// 请求访问失败时返回的根节点
// 	ErrorResponse	string	`json:"error_response"`
// 	// 请求失败返回的错误码
// 	Code	string	`json:"code"`
// 	// 请求失败返回的错误信息
// 	Msg	string	`json:"msg"`
// 	// 请求失败返回的子错误码
// 	SubCode	String
// 	// 请求失败返回的子错误信息
// 	SubMsg	String
// 	***_response	String
// }
