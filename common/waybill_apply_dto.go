package common

//电子面单模板请求dto
type WaybillTemplateRequest struct {
	//订单信息(所有模版=ALL，客户拥有的模版=OWNER)
	TemplatesType string `json:"templatestype"`
	// 快递公司类别
	LogisticType string `valid:"Required" json:"logistictype"`
}

//订单下载返回参数
type DownloadOrderListReturn struct {
	// code	String	必填	通用	64	返回码	10000
	// msg	String	必填	通用	64	返回消息	Success
	// subcode	String	必填	通用	200	子集编码	LXO.JD.REQUEST_FAIL
	// submessage	String	必填	通用	200	子级消息	订单已出库
	// polyapitotalms	int	必填	通用	64	菠萝派总耗时	102
	// polyapirequestid	String	必填	通用	64	请求菠萝派编号	20161222154212742
	// ishasnextpage	int	必填	通用	4	是否有下一页(判断下一页请以此参数为准)(...	0
	// numtotalorder	int	必填	通用	4	订单总数量，注意是符合条件的总数量，不是当前页的订单数量	1
	Code             string                          `valid:"MaxSize(64)" json:"code"`
	Msg              string                          `valid:"MaxSize(64)" json:"msg"`
	Subcode          string                          `valid:"MaxSize(200)" json:"subcode"`
	Submessage       string                          `valid:"MaxSize(200)" json:"submessage"`
	Polyapitotalms   string                          `json:"polyapitotalms"`
	Polyapirequestid string                          `valid:"MaxSize(64)" json:"polyapirequestid"`
	Ishasnextpage    string                          `json:"ishasnextpage"`
	Numtotalorder    string                          `json:"numtotalorder"`
	Orders           []DownLoadOrderListOrdersReturn `json:"orders"`
	Nexttoken        string                          `json:"nexttoken"`
}
type DownLoadOrderListOrdersReturn struct {
	// shoptype	string	可选	通用	25	店铺类型(普通=JH_001，分销=JH_002，经销=...+..	JH_001
	// platorderno	string	必填	通用	20	平台订单号	6442452481264
	// tradestatus	string	必填	通用	25	订单交易状态，展示菠萝派统一的交易状态编...+..	JH_01
	// tradestatusdescription	string	可选	通用	200	平台交易状态，展示平台原始的交易状态编码...	卖家关闭
	// tradetime	datetime	可选	通用	20	订单创建时间(格式:yyyy-MM-dd HH:mm:ss)	1900-01-01 00:00:00
	// modifytime	datetime	可选	通用	20	订单更新时间 (格式:yyyy-MM-dd HH:mm:ss，...	1900-01-01 00:00:00
	// collagetime	string	可选	通用	20	拼团成团时间 (格式:yyyy-MM-dd HH:mm:ss，...	1900-01-01 00:00:00
	// username	string	必填	通用	32	会员名，可用于区分不同会员	J1274551574
	// nick	string	必填	通用	32	买家昵称	J1274551574
	// buyermobile	string	必填	通用	32	买家电话	18888888888
	// receivername	string	可选	通用	20	收件人姓名	张三
	// country	string	可选	通用	2	国家二位简码	CN
	// province	string	必填	通用	64	州/省	浙江省
	// city	string	必填	通用	64	城市	杭州市
	// area	string	必填	通用	64	区县	西湖区
	// town	string	可选	通用	64	镇/街道	三墩镇
	// address	string	必填	通用	64	地址	尚坤生态创业园A211
	// payorderno	string	可选	通用	64	支付单号	2014022821001001020249575376
	// paytype	PayTypes	可选	通用	64	支付方式默认JH_Other(其他=JH_Other，支付...	JH_Alipay
	// shouldpaytype	string	可选	通用	64	支付方式	网上付款
	// zip	string	可选	通用	10	邮编	310000
	// phone	string	可选	通用	20	电话(手机、电话二选一)	0571-89845712
	// mobile	string	可选	通用	20	手机(手机、电话二选一)	15067888888
	// email	string	可选	通用	64	Email	differ@test.com
	// customerremark	string	可选	通用	200	买家备注	包装好
	// sellerremark	string	可选	通用	200	卖家备注	我会的
	// postfee	decimal	可选	通用	20	邮资	0
	// postinsurancefee	decimal	可选	通用	20	运费险	0
	// goodsfee	decimal	必填	通用	20	商品总金额（已减去商品优惠金额）	500
	// servicefee	decimal	必填	通用	20	服务费(单位：元)	32
	// totalamount	decimal	必填	通用	20	订单总金额(卖家实际收到金额，包含平台优...	500
	// realpaymoney	decimal	必填	通用	20	实际支付金额（用户支付金额，已减去优惠金...	400
	// favourablemoney	decimal	必填	通用	20	商家优惠金额（不包含平台优惠）	100
	// platdiscountmoney	decimal	必填	通用	64	平台优惠金额（开发票给平台时，可用此金额...	10
	// taxamount	decimal	必填	通用	64	订单总税费	102.0
	// tariffamount	decimal	必填	通用	64	订单总关税额	12.0
	// addedvalueamount	decimal	必填	通用	64	订单总增值税额	0.0
	// consumptiondutyamount	decimal	必填	通用	64	订单总消费额	90.0
	// commissionvalue	decimal	必填	通用	20	佣金或分销商提成金额	10
	// paytime	datetime?	可选	通用		支付时间(格式:yyyy-MM-dd HH:mm:ss)	1900-01-01 00:00:00
	// sendtype	SendTypes	可选	通用	64	货运方式，返回菠萝派统一的货运方式类型(...	JH_COD
	// sendstyle	string	可选	通用	64	货运方式说明，返回平台原始的货运方式类型	普通快递
	// codservicefee	decimal	必填	通用	20	COD服务费	0
	// sellerflag	string	必填	通用	32	卖家订单标签(无=JH_None，灰=JH_Gray，红=...	JH_Red
	// cardtype	string	必填	通用	4	证件类型(身份证=JH_01，护照=JH_02，港澳...	JH_01
	// idcard	string	可选	通用	20	证件号码	312055199001014872
	// idcardtruename	string	可选	通用	20	证件真实姓名	张三
	// idcardimgs	string	可选	通用	20	证件照(多个照片之间以','分隔)	http://www.ygjj.com/bookpic2/2012-02-21/new453046-20120221164344662253.jpg
	// whsecode	string	可选	通用	32	商品所在仓库编号	KU002
	// ishwgflag	int	必填	通用	2	是否为海外购(是=1；否=0)	1
	// deliverytype	string	可选	通用	15	海外购供货商发货方式(国内供货商发货=JH_0...+..	01
	// shopid	string	可选	通用	15	门店ID，用于支持多门店业务	32
	// mdbid	string	可选	通用	15	门店帮，用于支持多门店业务（京东）	23900
	// salespin	string	可选	通用	15	导购员pin（京东）	jd_5555563
	// isneedinvoice	int	可选	通用	2	是否需要发票(需要=1 ，不需要=0)	0
	// invoicetype	string	可选	通用	64	发票类型(不开票=JH_NONE，纸质普票=JH_01...	不开票
	// invoicebusinesstype	string	可选	通用	64	发票业务类型(个人=JH_01，商家/企业=JH_02...	JH_01
	// invoicetitle	string	可选	通用	64	发票抬头	杭州笛佛软件有限公司
	// invoicecontent	string	可选	通用	64	发票内容	明细
	// taxpayerident	string	可选	通用	64	纳税人识别号，税号	913305647250399591T
	// registeredaddress	string	可选	通用	200	注册地址	宁波高新区
	// registeredphone	string	可选	通用	200	注册电话	0571-87902677
	// depositbank	string	可选	通用	200	开户银行	中国建设银行宁波支行
	// bankaccount	string	可选	通用	200	银行账户	33300285136010502987
	// fetchtime	datetime	必填	通用	32	到店自提时间，若平台返回时间不是DateTime...	2016-11-01 10:00:00
	// fetchtimedesc	datetime	必填	通用	32	到店自提时间，注意部分平台时间不是DateTi...	2016-11-01 10:00:00
	// ordersource	string	可选	通用	64	订单来源	订单来源
	// customattr	string	可选	通用	64	自定义属性，预留字段
	// transportday	string	可选	通用	64	发货日期
	// 平台专有参数
	// goodinfos	GoodInfo[]	必填	通用		商品信息集合
	// coupondetails	CouponDetail[]	可选	通用		优惠明细，只用来展示参考，不能用来计算优...
	// serviceorders	List	可选	通用		服务子订单列表
	Shoptype               string            `json:"shoptype"`
	Platorderno            string            `json:"platorderno"`
	Tradestatus            string            `json:"tradestatus"`
	Tradestatusdescription string            `json:"tradestatusdescription"`
	Tradetime              string            `json:"tradetime"`
	Modifytime             string            `json:"modifytime"`
	Collagetime            string            `json:"collagetime"`
	Username               string            `json:"username"`
	Nick                   string            `json:"nick"`
	Buyermobile            string            `json:"buyermobile"`
	Receivername           string            `json:"receivername"`
	Country                string            `json:"country"`
	Province               string            `json:"province"`
	City                   string            `json:"city"`
	Area                   string            `json:"area"`
	Town                   string            `json:"town"`
	Address                string            `json:"address"`
	Payorderno             string            `json:"payorderno"`
	Paytype                string            `json:"paytype"`
	Shouldpaytype          string            `json:"shouldpaytype"`
	Zip                    string            `json:"zip"`
	Phone                  string            `json:"phone"`
	Mobile                 string            `json:"mobile"`
	Email                  string            `json:"email"`
	Customerremark         string            `json:"customerremark"`
	Sellerremark           string            `json:"sellerremark"`
	Postfee                string            `json:"postfee"`
	Postinsurancefee       string            `json:"postinsurancefee"`
	Goodsfee               string            `json:"goodsfee"`
	Servicefee             string            `json:"servicefee"`
	Totalamount            string            `json:"totalamount"`
	Realpaymoney           string            `json:"realpaymoney"`
	Favourablemoney        string            `json:"favourablemoney"`
	Platdiscountmoney      string            `json:"platdiscountmoney"`
	Taxamount              string            `json:"taxamount"`
	Tariffamount           string            `json:"tariffamount"`
	Addedvalueamount       string            `json:"addedvalueamount"`
	Consumptiondutyamount  string            `json:"consumptiondutyamount"`
	Commissionvalue        string            `json:"commissionvalue"`
	Paytime                string            `json:"paytime"`
	Sendtype               string            `json:"sendtype"`
	Sendstyle              string            `json:"sendstyle"`
	Codservicefee          string            `json:"codservicefee"`
	Sellerflag             string            `json:"sellerflag"`
	Cardtype               string            `json:"cardtype"`
	Idcard                 string            `json:"idcard"`
	Idcardtruename         string            `json:"idcardtruename"`
	Idcardimgs             string            `json:"idcardimgs"`
	Whsecode               string            `json:"whsecode"`
	Ishwgflag              string            `json:"ishwgflag"`
	Deliverytype           string            `json:"deliverytype"`
	Shopid                 string            `json:"shopid"`
	Mdbid                  string            `json:"mdbid"`
	Salespin               string            `json:"salespin"`
	Isneedinvoice          string            `json:"isneedinvoice"`
	Invoicetype            string            `json:"invoicetype"`
	Invoicebusinesstype    string            `json:"invoicebusinesstype"`
	Invoicetitle           string            `json:"invoicetitle"`
	Invoicecontent         string            `json:"invoicecontent"`
	Taxpayerident          string            `json:"taxpayerident"`
	Registeredaddress      string            `json:"registeredaddress"`
	Registeredphone        string            `json:"registeredphone"`
	Depositbank            string            `json:"depositbank"`
	Bankaccount            string            `json:"bankaccount"`
	Fetchtime              string            `json:"fetchtime"`
	Fetchtimedesc          string            `json:"fetchtimedesc"`
	Ordersource            string            `json:"ordersource"`
	Customattr             string            `json:"customattr"`
	Transportday           string            `json:"transportday"`
	Goodinfos              []GoodInfoV2      `json:"goodinfos"`
	Coupondetails          []CouponDetailV2  `json:"coupondetails"`
	Serviceorders          []ServiceOrdersV2 `json:"serviceorders"`
}
type ServiceOrdersV2 struct {
	Serviceid    string `json:"serviceid"`
	Servicename  string `json:"servicename"`
	Servicetype  string `json:"servicetype"`
	Serviceprice string `json:"serviceprice"`
	Servicenum   string `json:"servicenum"`
}
type CouponDetailV2 struct {
	Sku_id     string `json:"sku_id"`
	Coupontype string `json:"coupontype"`
	Type       string `json:"type"`
	Price      string `json:"price"`
	Code       string `json:"code"`
	Couponnum  string `json:"couponnum"`
}
type GoodInfoV2 struct {
	Productid             string `json:"productid"`
	Suborderno            string `json:"suborderno"`
	Taxsuborderno         string `json:"taxsuborderno"`
	Tradegoodsno          string `json:"tradegoodsno"`
	Tradegoodsname        string `json:"tradegoodsname"`
	Tradegoodsspec        string `json:"tradegoodsspec"`
	Goodscount            string `json:"goodscount"`
	Price                 string `json:"price"`
	Refundcount           string `json:"refundcount"`
	Discountmoney         string `json:"discountmoney"`
	Taxamount             string `json:"taxamount"`
	Tariffamount          string `json:"tariffamount"`
	Addedvalueamount      string `json:"addedvalueamount"`
	Consumptiondutyamount string `json:"consumptiondutyamount"`
	Refundstatus          string `json:"refundstatus"`
	Status                string `json:"status"`
	Remark                string `json:"remark"`
	Outskuid              string `json:"outskuid"`
	Platgoodsid           string `json:"platgoodsid"`
	Platskuid             string `json:"platskuid"`
	Outitemid             string `json:"outitemid"`
	Subgoods              string `json:"subgoods"`
	Isgift                string `json:"isgift"`
	Ishwgflag             string `json:"ishwgflag"`
	Deliverytype          string `json:"deliverytype"`
	Payorderid            string `json:"payorderid"`
	Packageorderid        string `json:"packageorderid"`
	Activityamount        string `json:"activityamount"`
	Specialamount         string `json:"specialamount"`
	Couponamount          string `json:"couponamount"`
	Productitemid         string `json:"productitemid"`
	Goodscount2           string `json:"goodscount2"`
	Isplatstorageorder    string `json:"isplatstorageorder"`
	Pictureurl            string `json:"pictureurl"`
	Goodtype              string `json:"goodtype"`
	Estimatecontime       string `json:"estimatecontime"`
	Fenxiaoprice          string `json:"fenxiaoprice"`
	Fenxiaopayment        string `json:"fenxiaopayment"`
	Suborderitemno        string `json:"suborderitemno"`
	Goodsorderattr        string `json:"goodsorderattr"`
	Ispresale             string `json:"ispresale"`
	Serialnumbers         string `json:"serialnumbers"`
	Cantsendreason        string `json:"cantsendreason"`
	Cansendgoods          string `json:"cansendgoods"`
	Fenxiaofreight        string `json:"fenxiaofreight"`
}

type DownLoadOrderListPostBizcontent struct {
	// orderstatus	string	可选	通用	25	订单交易状态(当抓取订单列表时必传)。注：...+..	JH_01
	// platorderno	string	可选	通用	32	平台订单号(当抓取当个订单时必传)	PX4040334233
	// starttime	datetime	可选	通用	20	开始时间(格式：yyyy-MM-dd HH:mm:ss)(当抓...	2016-06-15 12:23:32
	// endtime	datetime	可选	通用	20	截止时间(格式：yyyy-MM-dd HH:mm:ss)(当抓...	2016-07-15 08:32:00
	// timetype	string	可选	通用	25	订单时间类别(当抓取订单列表时必传，若需...+..	JH_01
	// pageindex	int	可选	通用	4	页码(当抓取订单列表时必传。寺库、达令网、融易购、聚美海淘不支持分页。云集品抓单的同时会导出订单，已导出的订单下次不会再抓)	1
	// pagesize	int	可选	通用	4	每页条数(当抓取订单列表时必传。寺库、达令网、融易购、聚美海淘不支持分页；Higo、eyee、美团B2B每页固定大小为20；云集品固定大小为50，但是会过滤未成团的订单)	10
	// shoptype	string	可选	通用	25	店铺类型(普通=JH_001，分销=JH_002，经销=...+..	JH_001
	// 平台专有参数
	// ordertype	string	可选	通用	25	订单类型(普通订单=JH_001，寻仓订单（唯品...+..	JH_001
	// randomnumber	string	必填	通用	64	淘宝随机字符串	tbxLGzL2r67me4zhYLHtDNvxxqPfjlgkAdU88pSPT55=
	// isneedflag	int	必填	通用	64	是否需要下载旗帜(下载=1，不下载=0)	1
	// isnotneeddetail	int	必填	通用	64	是否不需要下载详情(下载=0，不下载=1)	1
	Orderstatus     string `valid:"MaxSize(25)" json:"orderstatus"`
	Starttime       string `valid:"MaxSize(64)" json:"starttime"`
	Endtime         string `valid:"MaxSize(64)" json:"endtime"`
	Timetype        string `valid:"MaxSize(25)" json:"timetype"`
	Pageindex       int    `json:"pageindex"`
	Pagesize        int    `json:"pagesize"`
	Shoptype        string `valid:"MaxSize(25)" json:"shoptype"`
	Ordertype       string `valid:"MaxSize(25)" json:"ordertype"`
	Randomnumber    string `valid:"MaxSize(64)" json:"randomnumber"`
	Isneedflag      int    `json:"isneedflag"`
	Isnotneeddetail int    `json:"isnotneeddetail"`
	Nexttoken       string `json:"nexttoken"`
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

// 电子面单取消结构体发送内容
type WaybillCancelSend struct {
	Orders []WaybillCancel `json:"orders"`
}

// 电子面单取消结构体
type WaybillCancel struct {
	Orderno      string  `valid:"Required" json:"orderno"` //客户订单号
	Logisticno   string  `json:"logisticno"`               //运单号
	Logistictype string  `json:"logistictype"`             //物流类型(EMS=JH_001，圆通=JH_002，韵达=J...
	Weight       float64 `json:"weight"`                   //订单货物总重量(单位千克)
	Volumn       string  `json:"volumn"`                   //订单货物总体积(单位厘米)
	Remark       string  `json:"remark"`                   //*取消原因
	Warecode     string  `json:"warecode"`                 //*仓库编码
	Providercode string  `json:"providercode"`             //*承运商编码//京东必填
	Operatorname string  `json:"operatorname"`             //*取消操作人//京东必填
	Sendercoutry string  `json:"sendercoutry"`             //*发货人国家（用于判断取消哪个线路的包裹）
}

// 电子面单取消返回结构体主表
type WaybillCancelReturn struct {
	Code             string                     `json:"code"`
	Msg              string                     `json:"msg"`
	Subcode          string                     `json:"subcode"`
	Submessage       string                     `json:"submessage"`
	Polyapitotalms   string                     `json:"polyapitotalms"`
	Polyapirequestid string                     `json:"polyapirequestid"`
	Results          []*WaybillCancelReturnData `valid:"Required" json:"results"`
}

// 电子面单取消返回结构体数据表
type WaybillCancelReturnData struct {
	Issuccess  string `json:"issuccess"`
	Errorcode  string `json:"errorcode"`
	Message    string `json:"message"`
	Orderno    string `json:"orderno"`
	Logisticno string `json:"logisticno"`
}

// 电子面单请求结构体
type WaybillApplyNewRequest struct {
	// 物流服务商编码，必填
	CpCode string `valid:"Required" json:"cp_code"`
	// 菠萝派业务类型
	BusinessType uint `json:"business_type"`
	// 菠萝派业务类型
	LogisticsServices string `json:"logistics_services"`
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
