package common

import "encoding/json"

//电子面单模板请求dto
type WaybillTemplateRequest struct {
	//订单信息(所有模版=ALL，客户拥有的模版=OWNER)
	TemplatesType string `json:"templatestype"`
	// 快递公司类别
	LogisticType string `valid:"Required" json:"logistictype"`
}

//订单下载返回参数
type DownloadOrderListReturn struct {
	// code	string	必填	通用	64	返回码	10000
	// msg	string	必填	通用	64	返回消息	Success
	// subcode	string	必填	通用	200	子集编码	LXO.JD.REQUEST_FAIL
	// submessage	string	必填	通用	200	子级消息	订单已出库
	// polyapitotalms	int	必填	通用	64	菠萝派总耗时	102
	// polyapirequestid	string	必填	通用	64	请求菠萝派编号	20161222154212742
	// ishasnextpage	int	必填	通用	4	是否有下一页(判断下一页请以此参数为准)(...	0
	// numtotalorder	int	必填	通用	4	订单总数量，注意是符合条件的总数量，不是当前页的订单数量	1
	Code             string                          `valid:"MaxSize(64)" json:"code"`
	Msg              string                          `valid:"MaxSize(64)" json:"msg"`
	Subcode          string                          `valid:"MaxSize(200)" json:"subcode"`
	Submessage       string                          `valid:"MaxSize(200)" json:"submessage"`
	Polyapitotalms   string                          `json:"polyapitotalms"`
	Polyapirequestid string                          `valid:"MaxSize(64)" json:"polyapirequestid"`
	Ishasnextpage    json.Number                     `json:"ishasnextpage"`
	Numtotalorder    json.Number                     `json:"numtotalorder"`
	Orders           []DownLoadOrderListOrdersReturn `json:"orders"`
	Nexttoken        string                          `json:"nexttoken"`
	Tid              int64
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
	Postfee                json.Number       `json:"postfee"`
	Postinsurancefee       json.Number       `json:"postinsurancefee"`
	Goodsfee               json.Number       `json:"goodsfee"`
	Servicefee             json.Number       `json:"servicefee"`
	Totalamount            json.Number       `json:"totalamount"`
	Realpaymoney           json.Number       `json:"realpaymoney"`
	Favourablemoney        json.Number       `json:"favourablemoney"`
	Platdiscountmoney      json.Number       `json:"platdiscountmoney"`
	Taxamount              json.Number       `json:"taxamount"`
	Tariffamount           json.Number       `json:"tariffamount"`
	Addedvalueamount       json.Number       `json:"addedvalueamount"`
	Consumptiondutyamount  json.Number       `json:"consumptiondutyamount"`
	Commissionvalue        json.Number       `json:"commissionvalue"`
	Paytime                string            `json:"paytime"`
	Sendtype               string            `json:"sendtype"`
	Sendstyle              string            `json:"sendstyle"`
	Codservicefee          json.Number       `json:"codservicefee"`
	Sellerflag             string            `json:"sellerflag"`
	Cardtype               string            `json:"cardtype"`
	Idcard                 string            `json:"idcard"`
	Idcardtruename         string            `json:"idcardtruename"`
	Idcardimgs             string            `json:"idcardimgs"`
	Whsecode               string            `json:"whsecode"`
	Ishwgflag              json.Number       `json:"ishwgflag"`
	Deliverytype           string            `json:"deliverytype"`
	Shopid                 string            `json:"shopid"`
	Mdbid                  string            `json:"mdbid"`
	Salespin               string            `json:"salespin"`
	Isneedinvoice          json.Number       `json:"isneedinvoice"`
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
	Oaid                   string            `json:"oaid"`
	Shopname               string            `json:"shopname"`
	Tid                    int64             `json:"tid"`
	IsPremium              json.Number       `json:"ispremium"`
	PremiumAmount          json.Number       `json:"premiumamount"`
}
type ServiceOrdersV2 struct {
	Serviceid    string      `json:"serviceid"`
	Servicename  string      `json:"servicename"`
	Servicetype  string      `json:"servicetype"`
	Serviceprice json.Number `json:"serviceprice"`
	Servicenum   json.Number `json:"servicenum"`
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
	Productid             string      `json:"productid"`
	Suborderno            string      `json:"suborderno"`
	Taxsuborderno         string      `json:"taxsuborderno"`
	Tradegoodsno          string      `json:"tradegoodsno"`
	Tradegoodsname        string      `json:"tradegoodsname"`
	Tradegoodsspec        string      `json:"tradegoodsspec"`
	Goodscount            json.Number `json:"goodscount"`
	Price                 json.Number `json:"price"`
	Refundcount           json.Number `json:"refundcount"`
	Discountmoney         json.Number `json:"discountmoney"`
	Taxamount             json.Number `json:"taxamount"`
	Tariffamount          json.Number `json:"tariffamount"`
	Addedvalueamount      json.Number `json:"addedvalueamount"`
	Consumptiondutyamount json.Number `json:"consumptiondutyamount"`
	Refundstatus          string      `json:"refundstatus"`
	Status                string      `json:"status"`
	Remark                string      `json:"remark"`
	Outskuid              string      `json:"outskuid"`
	Platgoodsid           string      `json:"platgoodsid"`
	Platskuid             string      `json:"platskuid"`
	Outitemid             string      `json:"outitemid"`
	Subgoods              string      `json:"subgoods"`
	Isgift                string      `json:"isgift"`
	Ishwgflag             json.Number `json:"ishwgflag"`
	Deliverytype          string      `json:"deliverytype"`
	Payorderid            string      `json:"payorderid"`
	Packageorderid        string      `json:"packageorderid"`
	Activityamount        json.Number `json:"activityamount"`
	Specialamount         json.Number `json:"specialamount"`
	Couponamount          json.Number `json:"couponamount"`
	Productitemid         string      `json:"productitemid"`
	Goodscount2           string      `json:"goodscount2"`
	Isplatstorageorder    json.Number `json:"isplatstorageorder"`
	Pictureurl            string      `json:"pictureurl"`
	Goodtype              string      `json:"goodtype"`
	Estimatecontime       string      `json:"estimatecontime"`
	Fenxiaoprice          string      `json:"fenxiaoprice"`
	Fenxiaopayment        json.Number `json:"fenxiaopayment"`
	Suborderitemno        string      `json:"suborderitemno"`
	Goodsorderattr        string      `json:"goodsorderattr"`
	Ispresale             json.Number `json:"ispresale"`
	Serialnumbers         string      `json:"serialnumbers"`
	Cantsendreason        string      `json:"cantsendreason"`
	Cansendgoods          string      `json:"cansendgoods"`
	Fenxiaofreight        string      `json:"fenxiaofreight"`
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

//退款检测发送字段结构体
type BatchCheckRefundStatusBizcontent struct {
	Orders []CheckRefundOrderInfo `json:"orders"`
}

//退款检测订单明细结构体
type CheckRefundOrderInfo struct {
	Platorderno string `json:"platorderno"`
	Shoptype    string `json:"shoptype"`
	Countrycode string `json:"countrycode"`
}

//退款检测返回公共结构体
type CheckRefundReturn struct {
	Code             string                  `json:"code"`
	Msg              string                  `json:"msg"`
	Subcode          string                  `json:"subcode"`
	Submessage       string                  `json:"submessage"`
	Polyapitotalms   json.Number             `json:"polyapitotalms"`
	Polyapirequestid string                  `json:"polyapirequestid"`
	Results          []CheckRefundReturnData `json:"results"`
}
type CheckRefundReturnData struct {
	Issuccess               json.Number                                 `json:"issuccess"`               //必填	通用	1	是否成功(0表示失败；1表示成功)	0
	Code                    string                                      `json:"code"`                    //必填	通用	256	错误编码	Failed
	Message                 string                                      `json:"Message"`                 //可选	通用	256	是否成功	订单已出库
	Platorderno             string                                      `json:"platorderno"`             //必填	通用	64	平台订单号	SE88956989966
	Refundstatus            string                                      `json:"refundstatus"`            //必填	通用	25	退款状态(没有退款=JH_07，买家已经申请退款等待卖家同意=JH_01，卖家已经同意退款等待买家退货=JH_02，买家已经退货等待卖家确认收货=JH_03，卖家拒绝退款=JH_04，退款关闭=JH_05，退款成功=JH_06，退款中=JH_08，部分退款=JH_09，待审核=JH_10，其他=JH_99)
	Refundstatusdescription string                                      `json:"refundstatusdescription"` //可选	通用	32	退款状态说明	退款中
	Tradestatus             string                                      `json:"tradestatus"`             //必填	通用	25	订单交易状态(等待买家付款=JH_01，等待卖家发货=JH_02，等待买家确认收货=JH_03，交易成功=JH_04，交易关闭=JH_05，已暂停=JH_06，已锁定=JH_07，卖家部分发货=JH_08，订单冻结=JH_09，缺货订单=JH_10，投诉订单=JH_12，已拆单=JH_13，退换货订单=JH_14，待开票订单(亚马逊)=JH_15，处理中订单（库巴国美）=JH_17，未发货取消（唯品会JITX）=JH_18，已发货取消（唯品会JITX）=JH_19，已揽收取消（唯品会JITX）=JH_20，无需发货(格格家履约)=JH_21，取消待审核（孩子王与孩子王一件代发专用）=JH_22，已拒收（每日一淘专用）=JH_23，禁发订单（唯品会JITX专用）=JH_24，其他（返参使用）=JH_98，所有订单=JH_99)
	Childrenrefundstatus    []CheckRefundReturnChildrenRefundStatusItem `json:"childrenrefundstatus"`
	ShopId                  int64
}
type CheckRefundReturnChildrenRefundStatusItem struct {
	Refundno                string `json:"refundno"`                //可选	通用	32	退款订单号	RF158956655
	Suborderno              string `json:"suborderno"`              //可选	通用	64	子订单号,通常用于与抓单的子订单号匹配，...	622321321323345
	Productname             string `json:"productname"`             //可选	通用	512	商品名称	清风纸巾
	Platproductid           string `json:"platproductid"`           //必填	通用	64	平台商品ID或SKUID(SKUID优先)	622321321323323
	Tradegoodsno            string `json:"tradegoodsno"`            //必填	通用	64	外部商家编码或外部SKU编码(SKU编码优先)	6723213213
	Refundstatus            string `json:"refundstatus"`            //可选	通用	25	退款状态(没有退款=JH_07，买家已经申请退款等待卖家同意=JH_01，卖家已经同意退款等待买家退货=JH_02，买家已经退货等待卖家确认收货=JH_03，卖家拒绝退款=JH_04，退款关闭=JH_05，退款成功=JH_06，退款中=JH_08，部分退款=JH_09，待审核=JH_10，其他=JH_99)
	Refundstatusdescription string `json:"refundstatusdescription"` //可选	通用	32	退款状态说明	退款中
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
	Oaid string `json:"oaid"`
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
	Issuccess  json.Number `json:"issuccess"`
	Errorcode  string      `json:"errorcode"`
	Message    string      `json:"message"`
	Orderno    string      `json:"orderno"`
	Logisticno string      `json:"logisticno"`
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
	Oaid               string            `json:"oaid"`
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
type WmsLogisticsBizcontent struct {
	Orders []*WmsLogisticsPostOrder
}
type WmsLogisticsPostOrder struct {
	Orderno     string //!必填	通用	32	订单号	FO58965555996
	Plattradeno string //可选	通用	32	平台原始单号(多个原始单号以英文“,”分隔...	YS567417751123
	// Ismultiplepieces            json.Number              //可选	通用	4	是否为子母件(子母件=1，非子母件=0；默认0...	0
	Numpackage json.Number //!必填	通用	4	包裹数量(默认填写值为1，有子母件时“IsMu...	1
	// Logisticno                  string                   //可选	通用	32	运单号(仅限于先预约单号的平台，如果是子...	XO434334244543423
	Businesstype string      //!必填	通用	32	平台的业务类型(标准业务=JH_Normal)+..	JH_Normal
	Businessplat string      //可选	通用	20	电商平台标识码+..	Taobao
	Paymode      string      //可选	通用	32	快递支付方式(立即付款=0，货到付款=1，发...	货到付款
	Ordertype    string      //可选	通用	32	订单类型(普通订单=JH_Normal，退货单=JH_R...+..	JH_Normal
	Codpaymoney  json.Number //可选	通用	12	货到付款金额(OrderType=JH_COD时必传)	150.2
	// Packagemoney                json.Number              //可选	通用	12	订单包裹物品金额	238.05
	// Weight                      json.Number              //!必填	通用	4	订单重量(单位：克)	500
	// Supporpaymoney              json.Number              //可选	通用	12	保价赔付金额	220.50
	// Length                      json.Number              //!必填	通用	4	包裹长(单位：CM)	18
	// Width                       json.Number              //!必填	通用	4	包裹宽(单位：CM)	18
	// Height                      json.Number              //!必填	通用	4	包裹高(单位：CM)	18
	// Volume                      json.Number              //可选	通用	8	包裹体积	12
	// Ispickup                    json.Number              //可选	通用	1	是否上门揽收(上门揽收=1，不上门揽收=0；...	0
	// Producttype                 string                   //可选	通用	20	产品类型+..	服装
	Logistictype      string                 //可选	通用	32	物流类型(标准物流=JH_Normal，经济物流=JH...+..	JH_Normal
	Cpcode            string                 //!必填	通用	32	承运公司编码	POSTB
	Dmssorting        json.Number            //!必填	通用	32	是否使用智分宝预分拣(仓库WMS系统对接落地...	1
	Needencrypt       json.Number            //!必填	通用	32	设定取号返回的云打印报文是否加密(是=1，...	1
	Sender            *WmsLogisticsHumanInfo //!必填	发件人信息
	Receiver          *WmsLogisticsHumanInfo //!必填	收件人信息
	LogisticsServices string
	// Pickup                      *WmsLogisticsHumanInfo   //可选	通用	-	揽收信息(当Ispickup=1时必填)
	// Return                      *WmsLogisticsHumanInfo   //可选	通用	-	退货人
	Goods []*WmsLogisticsGoodsInfo //!必填	通用	-	商品信息集合
	// Labelformat                 string                   //!必填	通用	10	打印尺寸格式+..
	Shipperno string //?必填	UPS、联邦快递、顺丰国际、顺丰国际俄罗斯、京东阿尔法	32	支付账号(在顺丰指月结卡号)	6124851842487424
	// Warecode                    string                   //?必填	Panex、Wish邮、中邮小包、中通国际物流、菜鸟、万邑通ISP、唯品会	32	仓库代码+..	0003
	// Warehouselineid             json.Number              //?必填	Panex	32	货站自选线路+..	0003
	// Batchno                     string                   //?必填	唯品会	32	批次号+..
	// Isliquid                    json.Number              //?必填	Panex	0	是否液体(是=1; 否=0)	0
	// Ishasbattery                json.Number              //?必填	Panex、中邮小包、互联易、顺丰国际俄罗斯、Wish邮、速卖通物流、递四方速递、易宇通物流、去发货、E速宝	0	是否包含电池(是=1; 否=0)	0
	// Isinsurance                 json.Number              //?必填	联邦快递、京东阿尔法、JD、递四方速递	32	是否保价(保价=1; 不保价=0)	0
	// Isvip                       json.Number              //?必填	中通	32	是否VIP尊享件(是=1; 否=0)	0
	// Istopayment                 json.Number              //?必填	中通	32	是否到付件(是=1; 否=0)	0
	// Deliverytype                string                   //?必填	德邦物流、运东西	32	送货方式(自提=JH_Picked，送货(不含上楼)=...	JH_Picked
	// Backsignbill                string                   //?必填	顺丰、德邦物流、JD	32	签收回单(无需返单=JH_001，需要返单=JH_0...	0
	// Islesstruck                 json.Number              //?必填	德邦物流	32	是否零担下单(是=1;否=0)	0
	Customercode string //?必填	JD、京东阿尔法	32	商家编码(此商家编码需由商家向快递运营人...	1542154
	Customername string //?必填	京东阿尔法	32	商家名称(此商家名称需由商家向快递运营人...	1542154
	// Imagestyle                  string                   //?必填	联邦快递	32	图像类型(两联单=JH_TwoDuplicate，三联单=...	JH_TwoDuplicate
	Inputsite string //?必填	申通、京东阿尔法	32	录入网点	上海陈行公司
	Sitecode  string //?必填	安能快递、转运四方快递、京东阿尔法	32	网点名称	1354512
	// Ischeckrange                json.Number              //?必填	天天快递	32	是否检测超区(是=1;否=0)	0
	// Temprangetype               string                   //?必填	顺丰	32	温度范围(冷藏 = 1；冷冻 = 3)	1
	// Mainsubpaymode              string                   //?必填	EMS	1	一票多单计费方式(集中主单计费 = 1；平均...	1
	// Transtype                   string                   //?必填	EMS、中通国际物流、Tourline、JD	1	运输方式(陆运 = 0；航空 = 1；陆转航 = 2...	1
	// Isbbc                       json.Number              //?必填	中通国际物流、Tourline	32	是否是BBC订单(是=1;否=0)	0
	// Transtypecode               string                   //?必填	EMS、中通国际物流、Tourline	1	运输方式编码	1
	// Prodcode                    string                   //?必填	EMS、邮政国际	1	产品代码(文件 = 0100000000；信函 = 02000...	1
	// Isnanji                     string                   //?必填	EMS	1	是否集散邮件(是 =1； 否 = 0)	1
	// Currencytype                string                   //?必填	顺丰国际俄罗斯、转运四方快递、邮政国际	255	货币类型简写	USD
	// Platwebsite                 string                   //?必填	顺丰国际俄罗斯	255	平台网址(当Businesstype为顺丰国际电商专...	Http://Www.Taobao.Com
	// Crosscodeid                 string                   //?必填	贝海直邮	255	货站ID(纽约分拨中心=1 ,旧金山分拨中心=2 ...	1
	// Definedorderinfo            string                   //?必填	韵达、韵达国际、贝海直邮	255	订单自定义信息	红色、大码
	// Definedgoodsinfo            string                   //?必填	韵达、韵达国际	255	商品自定义信息	红色、大码
	// Payorderno                  string                   //?必填	贝海直邮	255	支付单号	NO011244
	// Payamount                   json.Number              //?必填	贝海直邮	16	支付金额	55
	// Senderaccount               string                   //?必填	联邦快递	16	寄件账号(9 位数字)	1545515674521
	// Is5kgpacking                json.Number              //?必填	联邦快递	4	是否5公斤装	1
	// Is3pl                       string                   //?可选	EWE物流	4	是否3PL订单(是=Y;否=N)	Y
	// Isusestock                  string                   //?可选	EWE物流	4	是否使用库存商品（不使用=0;使用=1）	0
	// Iseconomic                  json.Number              //?可选	EWE物流	4	是否走经济线	False
	// Specialhandling             string                   //?必填	联邦快递	4	特殊处理(留站自提=留站自提；签单返还（运...	1
	// Codtype                     string                   //?必填	德邦物流	4	代收货款类型(三日退=3，即日退=1)	1
	// Packageservice              string                   //?必填	德邦物流、拼多多金虹桥	4	包装类型（直接用中文）： 纸、纤、木（包...	纸箱
	// Isout                       json.Number              //?必填	德邦物流	1	是否外发	0
	// Receiveraccount             string                   //?必填	德邦物流	4	代收账号	6225848752785448
	// Receiveraccountname         string                   //?必填	德邦物流	4	代收账户开启名	张三
	// Isfresh                     json.Number              //?必填	顺丰	4	是否保鲜(是=1; 否=0)	1
	// Remark                      string                   //?必填	菠萝派快递、品骏、顺丰国际IBS	255	客服备注	周末送，工作日没人
	// Customerremark              string                   //?必填	菜鸟	255	客服备注	发申通
	Ordersource string //?可选	京东阿尔法、拼多多金虹桥、捷网物流	64	订单来源	订单来源
	// Providerid                  string                   //?可选	京东阿尔法	20	承运商ID(京东阿尔法承运商ID与物流类别两...	20160927770
	Providercode     string //?必填	京东阿尔法、唯品会JIT	20	承运商编码	20160927770
	Expresspaymethod string //?必填	京东阿尔法	20	快递费付款方式(顺丰必传)	1:寄方付
	Expresstype      string //?必填	京东阿尔法	20	快件产品类别(顺丰必传)	1.顺丰次日
	// Undeliverabledecision       string                   //?可选	速卖通物流、递四方速递、邮政国际	20	不可达处理	0
	// Servicename                 string                   //?可选	速卖通物流、EWE物流	64	服务名称	0
	// Cumstomscode                string                   //?可选	中通国际物流	64	海关编号	NBCUSTOMS
	// Totallogisticsno            string                   //?可选	中通国际物流	64	提货单号	88091011
	// Stockflag                   string                   //?必填	中通国际物流	1	是否集货(是 =1； 否 = 0)	1
	// Ebpcode                     string                   //?必填	中通国际物流	64	电商企业代码	1
	// Ecpname                     string                   //?必填	中通国际物流	64	电商平台名称(海关备案名称)	1
	// Ecpcodeg                    string                   //?必填	中通国际物流	64	电商平台代码(国检)	1
	// Ecpnameg                    string                   //?必填	中通国际物流	64	电商平台名称(国检)	1
	// Agentcode                   string                   //?可选	中通国际物流	32	代理企业编号	154643265
	// Agentname                   string                   //?可选	中通国际物流	32	代理企业名称	154643265
	// Totalshippingfee            json.Number              //?必填	中通国际物流、贝海直邮	64	订单运费	0
	// Feeunit                     string                   //?可选	中通国际物流	64	费用单位	元
	// Clearcode                   json.Number              //?必填	中通国际物流	64	关区代码	2991
	// Sellerid                    string                   //?必填	菜鸟	64	卖家Id	2991
	// User_id                     string                   //?必填	菜鸟	32	使用者ID	12
	// Logistics_services          string                   //?可选	菜鸟	32	服务值,如不需要特殊服务，该值为空	12
	// Object_id                   string                   //?必填	菜鸟	32	请求ID	12
	TemplateUrl string `json:"template_url"` //?必填	菜鸟、拼多多金虹桥	32	标准模板模板URL	Http://Xxx.Com
	// Order_channels_type         string                   //?必填	菜鸟、易宇通物流、去发货、捷网物流	32	订单渠道平台编码	TB
	// Trade_order_list            string                   //?必填	菜鸟	32	订单号,数量限制100	1222221
	Logisticsproductname string //?必填	菜鸟	32	菜鸟物流产品名称（JH_01:智选物流;JH_02:...	JH_01
	// Deptno                      string                   //?必填	JD	50	事业部编号	00001
	// Businessnetworkno           string                   //?必填	德邦物流	50	发货部门编码	00001
	// Sendertc                    string                   //?必填	JD	50	始发转运中心名称	00001
	// Pickupdate                  string                   //?必填	JD	50	上门揽件时间
	// Installflag                 json.Number              //?必填	JD	4	是否安维
	// Thirdcategoryno             string                   //?必填	JD	20	三级分类编码(安维必填)
	// Brandno                     string                   //?必填	JD	50	品牌ID(安维必填)
	// Productsku                  string                   //?必填	JD	50	商品Sku(安维必填)
	// Platcode                    string                   //?必填	JD	50	订单平台编码（必填）
	// Sequenceno                  string                   //?必填	唯品会	50	顺序号（必填）
	// Chinashipname               string                   //?必填	贝海直邮	10	国内配送公司（必填）
	// Taxpaytype                  string                   //可选	顺丰国际IBS	32	税金支付方式（必填）(立即付款=0，货到付...	货到付款
	// Taxsetaccounts              string                   //?必填	顺丰国际IBS	10	税金结算账号
	// Praceltype                  string                   //?必填	去发货、E速宝、捷网物流	1	包裹类型
	// Addressid                   string                   //可选	EbayDIS	255	发货地址ID+..	1
	// Shoptype                    string                   //可选	通用	25	店铺类型+..	JH_001
	// Consignpreferenceid         string                   //可选	EbayDIS	255	交运偏好ID+..	1
	// Notifycouriertype           string                   //?必填	顺丰	10	通知顺丰收派员收件方式
	// Mallmaskid                  string                   //?必填	拼多多金虹桥	64	代打店铺Id	184003167
	// Goodsdescription            string                   //可选	拼多多金虹桥	64	货品描述	测试货品
	// Openboxservice              string                   //可选	京东大件	64	开箱服务（京东大件用，0:否 1:开箱通电 2:...	0
	// Shopnick                    string                   //可选	奇门海外物流	64	店铺名称（奇门海外专用）	0
	// Isneedsignatureconfirmation json.Number              //?必填	顺丰	4	是否使用签收确认(是=1; 否=0)	1
}
type WmsLogisticsHumanInfo struct {
	Name            string //!必填	通用	32	姓名	张三
	Company         string //!必填	通用	32	公司	微软科技
	Phone           string //!必填	通用	25	电话	0571-85585585
	Mobile          string //!必填	通用	11	手机	13896985526
	Country         string //!必填	通用	2	国家二位简码(AD=安道尔，AE=阿联酋，AF=阿...	CN
	Province        string //!必填	通用	32	州省	浙江
	City            string //!必填	通用	32	城市	杭州
	Area            string //!必填	通用	50	区县	江干区
	Town            string //!必填	通用	50	镇（街道）	三里亭
	Address         string //!必填	通用	200	地址	浙江杭州市江干区秋涛路255号10楼302
	Zip             string //!必填	通用	20	邮编	311703
	Email           string //!必填	通用	100	电子邮箱	Sender@163.Com
	Userid          string //!必填	Ebay、联邦快递、万邑通ISP	32	用户ID	14443
	Certificatetype string //!必填	韵达国际、中通国际物流、邮政国际	32	证件类型(身份证=ID，护照=PP, 军官证=JG, ...	ID
	Certificate     string //!必填	韵达国际、中环国际速递、中通国际物流、邮政国际	32	证件号(身份证号)	321201199801051678
	Certificatename string //!必填	UEQ快递	32	证件姓名(身份姓名)	张三
	Addresscode     string //!必填	万邑通ISP、速卖通物流	8	地址编号	N001
	Linker          string //可选	邮政国际	255	用户联系人+..	Zhang Hanhan
	Taxpayerident   string //可选	通用	255	增值税税号	4444555655
}
type WmsLogisticsReturn struct {
	Code             string      //必填	通用	64	返回码	10000
	Msg              string      //必填	通用	64	返回消息	Success
	Subcode          string      //必填	通用	200	子集编码	LXO.JD.REQUEST_FAIL
	Submessage       string      //必填	通用	200	子级消息	订单已出库
	Polyapitotalms   json.Number //必填	通用	64	菠萝派总耗时	102
	Polyapirequestid string      //必填	通用	64	请求菠萝派编号	20161222154212742
	Results          []*WmsLogisticsPostOrderResultItem
}
type WmsLogisticsGoodsInfo struct {
	Cnname          string      //!必填	通用	200	货品中文名	乔丹运动鞋
	Enname          string      //可选	通用	200	货品英文名	Nike Air Jordan
	Count           json.Number //!必填	通用	4	货品数量	122
	Currencytype    string      //可选	通用	10	货币类型(3位简码，默认CNY)	USD
	Price           json.Number //可选	通用	12	申报价值，默认0	10000.55
	Weight          json.Number //可选	通用	4	货品重量(克)	500
	Unit            string      //!必填	通用	10	计件单位	包
	Taxationid      string      //可选	通用	30	行邮税税号	622321321323323
	Productid       string      //可选	通用	50	平台产品ID，线上发货必填ID	622321321323323
	Sourceordercode string      //!必填	奇门海外物流	4	交易平台订单	1
	Innercount      json.Number //可选	贝海直邮	0	内件数量	0
	Length          string      //可选	中通国际物流	0	商品长度	0
	Width           string      //可选	中通国际物流	0	商品宽度	0
	Height          string      //可选	中通国际物流	0	商品高度	0
	Dutymoney       json.Number //可选	中通国际物流	12	预付税费，默认0	10000.55
	Isblinsure      json.Number //可选	中通国际物流	0	是否保价(0-否 1-是) (如:没有内容,默认填...	False
	Remark          string      //可选	云途物流、递四方速递	0	商品备注	0
	Isaneroidmarkup json.Number //!必填	速卖通物流	0	是否属于非液体化妆品(是=1; 否=0)	0
	Isonlybattery   json.Number //!必填	速卖通物流、EbayDIS	0	是否纯电池(是=1; 否=0)	0
	Productbrand    string      //可选	中环国际速递、顺丰	32	商品品牌	海尔
	Productattrs    string      //可选	中环国际速递、万邑通ISP、顺丰、中通国际物流	32	商品规格	2L
	Productmaterial string      //可选	中通国际物流	32	商品材质	2L
	Hscode          string      //可选	燕文快递、互联易、顺丰国际	32	商品海关编码(燕文快递发货方式为 “香港Fe...	HG7865239955
	Goodurl         string      //可选	顺丰国际俄罗斯	255	电商专递货物Url(当Businesstype为“顺丰国...	Http://Www.Taobao.Com/9040904422.Html
	Categoryid      string      //可选	贝海直邮、Panex	255	分类ID+..	1
	Categoryid2     string      //可选	Panex	255	分类ID2+..	1
	Plattradeno     string      //可选	Ebay、万邑通ISP	255	商品交易ID+..	1
	Origincountry   string      //可选	增速跨境、EbayDIS	2	商品原产国二位简码(AD=安道尔，AE=阿联酋...	CN
	Outerid         string      //可选	贝海直邮	255	外部商家编码+..	1
	Position        string      //可选	GOOD快递	255	仓位+..	1
	Supportbattery  string      //可选	EbayDIS	255	带电类型(内置电池Or支持安装电池)+..	1
	Description     string      //可选	邮政国际	255	商品描述（内件成分）+..	1
	Elecquaid       string      //可选	EbayDIS	255	带电资质证书ID，若包裹带电则必填(EbayDIS...+..	1
}
type WmsLogisticsPostOrderResultItem struct {
	Issuccess              json.Number       //!必填	通用	-	是否成功(0表示失败；1表示成功)	0
	Errorcode              string            //!必填	通用	-	错误码	-
	Message                string            //!必填	通用	-	失败说明	-
	Orderno                string            //!必填	通用	-	客户订单号	SE88956989966
	Logisticno             string            //!必填	通用	-	运单号	HT8895SW389966
	Logisticchildno        string            //!必填	通用	32	子运单号(子母单)	HT8895SW389966
	Origincode             string            //可选		通用	-	原寄地区域代码	ED3
	Destcode               string            //可选		通用	-	目的地区域代码	WX5
	Markers                string            //可选		通用	-	大头笔(韵达、韵达国际返回PDF字符串)	浙
	Markercode             string            //可选		通用	-	大头笔编码	32432
	Packages               []*WmsPackageInfo //!必填	通用	-	包裹信息(菜鸟多包裹专用字段)(菜鸟返回包裹信息优先取值)
	Plattradeno            string            //可选		LaZaDa物流、E速宝	255	商品交易ID+..	1
	Logisticcompanycode    string            //?必填	菜鸟	-	快递公司编码	STO
	Originname             string            //?必填	JD、百世汇通、安能快递、德邦物流、跨越速运	32	原寄地区名称	上海
	Destname               string            //?必填	JD、百世汇通、安能快递、德邦物流、跨越速运	32	目的地区名称	杭州
	Originsitename         string            //?必填	德邦物流	32	原寄站点名称	上海转运站
	Destsitename           string            //?必填	JD、安能快递、德邦物流、跨越速运	32	目的站点名称	上海站
	Destsiteid             string            //?必填	JD、百世汇通、安能快递	32	目的站点编码	05
	Trackid                string            //?必填	云途物流、天翼快递	32	追踪单号	03
	Origincrosscode        string            //?必填	JD	32	始发道口号	03
	Origintabletrolleycode string            //?必填	JD	32	始发笼车号	01
	Desttabletrolleycode   string            //?必填	JD	32	目的笼车号	04
	Destcrosscode          string            //?必填	JD	32	目的滑道号	03
	Agingname              string            //?必填	JD	32	时效名称	无时效
	Aging                  string            //?必填	JD	32	时效	2
	Road                   string            //?必填	JD	32	路区	解
	Pkgcode                string            //可选		百世汇通、京东阿尔法、跨越速运	32	集包编码	03
	Pkgname                string            //可选		百世汇通、京东阿尔法、跨越速运、申通新、申通	32	集包地名称	03
	Resultlogisticno       string            //可选		顺丰	32	签回单服务运单号	WX5
	Agentlogisticno        string            //?必填	顺丰国际俄罗斯、顺丰国际	32	代理运单号	154643265
	Printno                string            //?必填	威速易、万邑通ISP	32	打印单号	154643265
	Receiveorgcode         string            //?必填	品骏	32	收货站点编码	154643265
	Receiveorgname         string            //?必填	品骏	32	收货站点名称	浙江杭州
	Receivesortingcode     string            //?必填	品骏	32	分拣编码	L16
	Printinfo              string            //?必填	菜鸟、京东大件	-	打印Json串	L16
	Printurl               string            //?必填	顺丰国际IBS	-	打印网址	L16
	Invoiceurl             string            //?必填	顺丰国际IBS	-	发票打印网址	L16
	Recommandreason        string            //?必填	菜鸟	-	原因	原因
	Lpnumber               string            //?必填	速卖通物流	-	多包裹情况下：{"运单号1":"LP码1","运单号...	123456
	Channelcode            string            //?必填	速卖通物流	-	多包裹情况下：{"运单号1":"渠道编码1","运...	123456
	Warehouseorderid       string            //?必填	速卖通物流	-	物流订单号，主要用于拆单发货预约单号的情...	123456
	Identification         string            //?必填	跨越速运、菜鸟	-	类型标识	商
	Procode                string            //可选		顺丰	-	类型标识	商
	Destroutelabel         string            //可选		顺丰	-	类型标识	商
	Destteamcode           string            //可选		顺丰	-	类型标识	001
	Codingmapping          string            //可选		顺丰	-	类型标识	S10
	Codingmappingout       string            //可选		顺丰	-	类型标识	S10
	Abflag                 string            //可选		顺丰	-	类型标识	商
	Facesheetid            string            //可选		菜鸟	-	面单号
	Fetchcode              string            //可选		菜鸟	-	取件码
	Offcode                string            //可选		菜鸟	-	核销码
	Producttype            string            //可选		通用	20	产品类型+..	服装
	Expressoperationmode   string            //可选		通用	20	运营模式+..	航
	Goodstype              string            //可选		通用	20	业务类型+..	服装
	Transtype              string            //?必填	JD	1	运输方式(1：陆运 2：航空)	1
	Newicon                string            //?必填	顺丰	1	NewIcon 标	1
	Newabflag              string            //?必填	顺丰	1	NewAbFlag 标	1
	Xbflag                 string            //?必填	顺丰	1	Xb 标	1
}
type WmsPackageInfo struct {
	Logisticno      string //!必填	通用	32	运单号	HT8895SW389966
	Logisticchildno string //!必填	通用	32	子运单号(子母单)	HT8895SW389966
	Printinfo       string //!必填	通用	2000	打印Json串	{"Data":{"Recipient":{"Address":{"City":"北京市","Detail":"花家地社区卫生服务站","District":"朝阳区","Province":"北京","Town":"望京街道"},"Mobile":"1326443654","Name":"Bar","Phone":"057123222"},"RoutingInfo":{"Consolidation":{},"Origin":{"Code":"POSTB"},"Sortation":{"Name":"杭州"
}

//订单发货接口结构体
type WmsBusinessSendBizcontent struct {
	Platorderno         string      //必填 通用 32 平台订单号 TS04594434433
	Logisticno          string      //必填 通用 32 快递运单号 WR6685851555
	Logistictype        string      //必填 通用 25 快递类别(JH前缀为国内快递 ，JHI为国际快...+.. JH_001
	Logisticname        string      //可选 通用 32 快递名称 顺风
	Sendtype            string      //可选 通用 32 订单发货类别(自己联系物流=JH_01，在线下...+.. JH_01
	Issplit             json.Number //可选 通用 4 是否拆单发货(拆单=1 ，不拆单=0) 0
	Shoptype            string      //可选 通用 25 店铺类型(好药师代表ERP标识)(普通=JH_001...+.. JH_001
	Sendername          string      //可选 通用 15 发货人姓名 张三
	Sendertel           string      //可选 通用 15 发货人联系电话 15047788954
	Senderaddress       string      //可选 通用 15 发货人地址(省市区之间以空格分隔) 上海 上海市 宝山区 逸仙路2816号华滋奔腾大厦A栋14楼
	Ishwgflag           json.Number //可选 通用 2 是否为海外购(是=1；否=0) 1
	Deliverytype        string      //可选 通用 15 海外购供货商发货方式,新蛋商城使用平台自...+.. 01
	Packagenumber       string      //可选 通用 15 拆单发货包裹总数量 3
	Logisticfee         string      //必填 蚂蚁销客 25 快递费 1
	Ismodsendinfo       json.Number //可选 通用 4 是否是修改已发货物流信息操作(是=1 ，否=0... 0
	Whsecode            string      //必填 寺库、唯品会JIT 32 商品所在仓库编号 KU002
	Countrycode         string      //必填 LaZaDa、Wish 15 国家编码(中国=JH_01，美国=JH_02，德国=JH... UK
	Sellernick          string      //必填 千米、每日一淘 15 电商云商家编号A开头 Jd121574147
	Packageorderid      string      //必填 苏宁易购 32 合并包裹单号(IsHwgFlag=1 且 DeliveryType... 1100202311
	Website             string      //必填 速卖通 32 网站地址(当速卖通LogisticType=Other时，... 1100202311
	Olderlogisticno     string      //必填 速卖通 32 旧的快递运单号 WR6685851555
	Olderlogistictype   string      //必填 通用 25 旧的快递类别(JH前缀为国内快递 ，JHI为国...+.. JH_001
	Invoiceno           string      //必填 大V店 32 发票号码 1100202311
	Invoicedate         string      //必填 大V店 32 开票日期 1100202311
	Tradetime           string      //必填 亚马逊自营 32 下单日期 1100202311
	Transactionid       string      //可选 雅虎 45 交易序号 222425696737
	Fxtid               string      //可选 Ebay、唯品会JIT 45 发货订单号 222425696737-1878361307012
	Isdomesticdelivery  bool        //可选 洋码头 5 是否国内段发货 222425696737-1878361307012
	Zylogisticcompany   string      //可选 淘宝 32 自有物流公司名称 新华物流
	Zylogisticphone     string      //可选 淘宝 32 自有物流公司电话 1236547890123
	Installlogisticname string      //可选 淘宝 32 安装物流公司名称 新华物流
	Installlogisticcode string      //可选 淘宝 32 安装物流公司编码 XBWL
	Feature             string      //可选 淘宝 32 商品识别码(淘宝3C商品必传) IdentCode=Tid:Aaa,Bbb;MachineCode=Tid2:Aaa
	Resid               string      //可选 淘宝 32 资源Id 12371273
	Rescode             string      //可选 淘宝 32 资源Code DISTRIBUTOR_123
	Orderstatus         string      //可选 国美自营 32 订单状态 1
	Isrequestwaybill    bool        //可选 京东 4 是否需要请求上传运单信息到青龙系统 False
	Waybillrequest      string      //可选 京东 100 青龙系统上传信息
	Receiveraddress     string      //必填 亚马逊自营 15 收货人地址 上海 上海市 宝山区 逸仙路2816号华滋奔腾大厦A栋14楼
	Ordertype           string      //必填 立白积分商城、蜜芽 15 订单类别
	Shopid              string      //必填 当当网 15 店铺ID
	Deliveryname        string      //必填 五星电器 15 送货人姓名
	Deliveryphone       string      //必填 五星电器 15 送货人电话
	Customerremark      string      //必填 米多商城 100 买家备注信息
	Giftcount           string      //必填 米多商城 100 礼品商品数量
	Ordernumber         string      //可选 米多商城 100 订单编号。活动订单时必传
	Verifycode          string      //必填 通用 25 提货码 324342342XX
	Eventcreatetime     string      //必填 通用 32 流转节点发生时间 2019-03-01 12:00:00
	Plantime            string      //必填 通用 32 配送日期 2019-03-01 12:00:00
	Desc                string      //必填 通用 256 流转节点的详细地址及操作描述 浙江省杭州市西湖区上车扫描
	Sequenceno          string      //必填 通用 256 配送序号 1
	Event               json.Number //必填 通用 256 事件编码 10
	Sipshopid           json.Number //必填 Shopee 64 SIP业务的子店铺ID 184003167
	Goods               []*WmsBusinessSendGoods
}
type WmsBusinessSendGoods struct {
	Platproductid   string      //可选 通用 32 平台商品ID 51654561
	Suborderno      string      //可选 通用 64 子订单号(从抓单中获取) 622321321323345
	Tradegoodsno    string      //可选 通用 64 外部商家编码(从抓单中获取) 6723213213
	Sublogistictype string      //可选 通用 25 快递类别(JH前缀为国内快递 ，JHI为国际快... JH_001
	Sublogisticname string      //可选 通用 30 子快递名称 顺风
	Sublogisticno   string      //可选 通用 32 子快递运单号 WR6685851555
	Barcode         string      //可选 通用 64 商品条形码(从抓单中获取) 6723213213
	Count           json.Number //可选 通用 11 发货数量 1
}
type WmsBusinessSendReturn struct {
	Code             string      //必填 通用 64 返回码 10000
	Msg              string      //必填 通用 64 返回消息 Success
	Subcode          string      //必填 通用 200 子集编码 LXO.JD.REQUEST_FAIL
	Submessage       string      //必填 通用 200 子级消息 订单已出库
	Polyapitotalms   json.Number //必填 通用 64 菠萝派总耗时 102
	Polyapirequestid string      //必填 通用 64 请求菠萝派编号 20161222154212742
	Codmoney         json.Number //必填 京东 64 货到付款金额 26.30
	Logisticnos      string      //必填 LaZaDa 64 物流单号 26.30
	Subplatorderno   string      //必填 苏宁特卖 32 订单行号，仅限异步模式 00890643547303
	Logisticname     string      //可选 通用 32 快递名称 顺风
	Logistictype     string      //必填 通用 25 快递类别(JH前缀为国内快递 ，JHI为国际快...+.. JH_001
}
