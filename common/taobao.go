package common

import "encoding/json"

type TaobaoQimenDeliveryorderBatchcreate struct {
	Method         string                                                             //!是	API接口名称
	App_key        string                                                             //!是	TOP分配给应用的AppKey
	Target_app_key string                                                             //否	被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效
	Sign_method    string                                                             //!是	签名的摘要算法，可选值为：Hmac，Md5，Hmac-Sha256。
	Sign           string                                                             //!是	API输入参数签名结果，签名算法介绍请点击这里
	Session        string                                                             //否	用户登录授权成功后，TOP颁发给应用的授权信息，详细介绍请点击这里。当此API的标签上注明：“需要授权”，则此参数必传；“不需要授权”，则此参数不需要传；“可选授权”，则此参数为可选
	Timestamp      string                                                             //!是	时间戳，格式为Yyyy-MM-Dd HH:Mm:Ss，时区为GMT+8，例如：2015-01-01 12:00:00。淘宝API服务端允许客户端请求最大时间误差为10分钟
	Format         string                                                             //否	响应格式。默认为Xml格式，可选值：Xml，Json。
	V              string                                                             //!是	API协议版本，可选值：2.0
	Partner_id     string                                                             //否	合作伙伴身份标识
	Simplify       json.Number                                                        //否	是否采用精简JSON返回格式，仅当Format=Json时有效，默认值为：False
	CustomerId     json.Number                                                        //否	三方对接过程中，业务自定义路由参数，更多在奇门仓储等场景中使用
	Request        TaobaoQimenDeliveryorderBatchcreateDeliveryOrderBatchCreateRequest //false
}
type TaobaoQimenDeliveryorderBatchcreateDeliveryOrderBatchCreateRequest struct {
	//orders      []*TaobaoQimenDeliveryorderBatchcreateOrder //订单信息
	extendProps string //Map	//false		扩展属性
}

// type TaobaoQimenDeliveryorderBatchcreateOrder struct {
// 	deliveryOrder *DeliveryOrder //false		发货单信息
// 	orderLines    []*OrderLine   //false		单据列表
// }
type TaobaoQimenDeliveryorderBatchcreateDeliveryOrder struct {
	DeliveryOrderCode    string //!True	TB1234	出库单号
	PreDeliveryOrderCode string //False	Old1234	原出库单号(ERP分配)
	PreDeliveryOrderId   string //False	Oragin1234	原出库单号(WMS分配)
	OrderType            string //!True	JYCK	出库单类型(JYCK=一般交易出库单;HHCK=换货出库单;BFCK=补发出库单;QTCK=其他出库单)
	WarehouseCode        string //!True	OTHER	仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	OrderFlag            string //False	COD	订单标记(用字符串格式来表示订单标记列表:例如COD=货到付款;LIMIT=限时配 送;PRESELL=预 售;COMPLAIN=已投诉;SPLIT=拆单;EXCHANGE=换货;VISIT=上 门;MODIFYTRANSPORT=是否 可改配送方式;CONSIGN = 物流宝代理发货;SELLER_AFFORD=是否卖家承担运费;FENXIAO=分销订 单)
	SourcePlatformCode   string //!True	TB	订单来源平台编码(TB=淘宝、TM=天猫、JD=京东、DD=当当、PP=拍拍、YX= 易讯、 EBAY=Ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、 JM=聚美、LF=乐蜂 、MGJ=蘑菇街、 JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿 里巴巴、POS=POS门店、 MIA=蜜芽、OTHER=其他(只传英文编 码))
	SourcePlatformName   string //False	淘宝	订单来源平台名称
	CreateTime           string //!True	2016-07-06 12:00:00	发货单创建时间(YYYY-MM-DD HH:MM:SS)
	PlaceOrderTime       string //!True	2016-07-06 12:00:00	前台订单/店铺订单的创建时间/下单时间
	PayTime              string //False	2016-07-06 12:00:00	订单支付时间(YYYY-MM-DD HH:MM:SS)
	PayNo                string //False	P1234	支付平台交易号
	OperatorCode         string //False	0123	操作员(审核员)编码
	OperatorName         string //False	老王	操作员(审核员)名称
	OperateTime          string //!True	2016-07-06 12:00:00	操作(审核)时间(YYYY-MM-DD HH:MM:SS)
	ShopNick             string //!True	淘宝店	店铺名称
	SellerNick           string //False	淘宝店主	卖家名称
	BuyerNick            string //False	淘公仔	买家昵称
	TotalAmount          string //False	123	订单总金额(订单总金额=应收金额+已收金额=商品总金额-订单折扣金额+快递费用 ;单位 元)
	ItemAmount           string //False	234	商品总金额(元)
	DiscountAmount       string //False	111	订单折扣金额(元)
	Freight              string //False	111	快递费用(元)
	ArAmount             string //False	111	应收金额(消费者还需要支付多少--货到付款时消费者还需要支付多少约定使用这个字 段;单位元 )
	GotAmount            string //False	111	已收金额(消费者已经支付多少;单位元)
	ServiceFee           string //False	111	COD服务费
	LogisticsCode        string //!True	SF	物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通 、ZTO=中 通(ZTO)、HTKY=百世汇通、UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全 峰、FAST=快捷 、POSTB=邮政小包、 GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、AMAZON=亚马逊物流、 OTHER=其他)
	LogisticsName        string //False	顺丰	物流公司名称
	ExpressCode          string //False	Y12345	运单号
	LogisticsAreaCode    string //False	043300	快递区域编码
	DeliveryRequirements *TaobaoQimenDeliveryorderBatchcreateDeliveryRequirements
}
type TaobaoQimenDeliveryorderBatchcreateDeliveryRequirements struct {
	ScheduleType      json.Number //False	1	投递时延要求(1=工作日;2=节假日;101=当日达;102=次晨达;103=次日达;104= 预约 达)
	ScheduleDay       string      //False	2016-07-06	要求送达日期(YYYY-MM-DD)
	ScheduleStartTime string      //False	12:00:00	投递时间范围要求(开始时间;格式：HH:MM:SS)
	ScheduleEndTime   string      //False	13:00:00	投递时间范围要求(结束时间;格式：HH:MM:SS)
	DeliveryType      string      //False	PTPS	发货服务类型(PTPS:普通配送;LLPS:冷链配送;HBP:环保配)
}
type TaobaoQimenDeliveryorderBatchcreateSenderInfo struct {
	Company       string //False	淘宝	公司名称
	Name          string //!True	老王	姓名
	ZipCode       string //False	043300	邮编
	Tel           string //False	81020340	固定电话
	Mobile        string //!True	13214567869	移动电话
	Email         string //False	345@Gmail.Com	电子邮箱
	CountryCode   string //False	051532	国家二字码
	Province      string //!True	浙江省	省份
	City          string //!True	杭州	城市
	Area          string //False	余杭	区域
	Town          string //False	横加桥	村镇
	DetailAddress string //!True	杭州市余杭区989号	详细地址
}

// type TaobaoQimenDeliveryorderBatchcreateReceiverInfo struct {
// }
