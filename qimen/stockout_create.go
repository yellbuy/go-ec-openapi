package qimen

import (
	"encoding/json"
	"encoding/xml"
)

func StockOutCreateParse(body []byte) (res *WmsQimenStockout, err error) {
	res = new(WmsQimenStockout)
	err = xml.Unmarshal(body, res)
	return
}

//奇门出库单确认结构体
type WmsQimenStockoutConfirm struct {
	XMLName       xml.Name                             `xml:"request"`
	DeliveryOrder WmsQimenStockoutConfirmDeliveryOrder `xml:"deliveryOrder"` //false		deliveryOrder
	Packages      WmsQimenStockoutConfirmPackages      `xml:"packages"`      //false		packages
	OrderLines    WmsQimenStockoutConfirmOrderLine     `xml:"orderLines"`    //OrderLine[]	//false		orderLines
	//ExtendProps   map[string]interface{}               `xml:"extendProps"`   //false		扩展属性
}
type WmsQimenStockoutConfirmPackages struct {
	Package []WmsQimenStockoutConfirmPackage `xml:"package"`
}
type WmsQimenStockoutConfirmOrderLine struct {
	OrderLine []WmsQimenStockoutConfirmOrderLines `xml:"orderLine"`
}
type WmsQimenStockoutConfirmDeliveryOrder struct {
	TotalOrderLines       json.Number `xml:"totalOrderLines"`       //false	11	单据总行数
	DeliveryOrderCode     string      `xml:"deliveryOrderCode"`     //!true	Ox123456	出库单号
	DeliveryOrderId       string      `xml:"deliveryOrderId"`       //false	Dx123456	仓储系统出库单号
	WarehouseCode         string      `xml:"warehouseCode"`         //!true	Wx123456	仓库编码
	OrderType             string      `xml:"orderType"`             //!true	PTCK	出库单类型
	Status                string      `xml:"status"`                //false	NEW	出库单状态
	OutBizCode            string      `xml:"outBizCode"`            //false	23456	外部业务编码
	ConfirmType           json.Number `xml:"confirmType"`           //false	1	支持出库单多次发货的状态位
	LogisticsCode         string      `xml:"logisticsCode"`         //false	SF	物流公司编码
	LogisticsName         string      `xml:"logisticsName"`         //false	顺丰	物流公司名称
	ExpressCode           string      `xml:"expressCode"`           //false	Q123456	运单号
	OrderConfirmTime      string      `xml:"orderConfirmTime"`      //false	2015-09-12 12:00:00	订单完成时间
	ResponsibleDepartment string      `xml:"responsibleDepartment"` //false	财务部	该笔出库单的费用承担部门或责任部门
	SubOrderType          string      `xml:"subOrderType"`          //false	hss	出库单确认其他出库单的子类型，用于 orderType设置为其他 出库单时设置
}
type WmsQimenStockoutConfirmPackage struct {
	LogisticsName       string                                          `xml:"logisticsName"`       //false	顺丰	物流公司名称
	ExpressCode         string                                          `xml:"expressCode"`         //false	O987654	运单号
	PackageCode         string                                          `xml:"packageCode"`         //false	L123456	包裹编号
	Length              string                                          `xml:"length"`              //false	12	包裹长度(厘米)
	Width               string                                          `xml:"width"`               //false	11	包裹宽度(厘米)
	Height              string                                          `xml:"height"`              //false	11	包裹高度(厘米)
	Weight              string                                          `xml:"weight"`              //false	11	包裹重量(千克)
	Volume              string                                          `xml:"volume"`              //false	12	包裹体积(升)
	PackageMaterialList []WmsQimenStockoutConfirmPackagePackageMaterial `xml:"packageMaterialList"` //false		packageMaterialList
	Items               WmsQimenStockoutConfirmPackageItem              `xml:"items"`               //false		items
}
type WmsQimenStockoutConfirmPackageItem struct {
	Item []WmsQimenStockoutConfirmPackageItems `xml:"item"`
}
type WmsQimenStockoutConfirmPackagePackageMaterial struct {
	Type     string      `xml:"type"`     //false	XL	包材型号
	Quantity json.Number `xml:"quantity"` //false	11	包材的数量
}
type WmsQimenStockoutConfirmPackageItems struct {
	ItemCode string      `xml:"itemCode"` //!true	JO123456	商品编码
	ItemId   string      `xml:"itemId"`   //false	TO67890	商品仓储系统编码
	Quantity json.Number `xml:"quantity"` //!true	11	包裹内该商品的数量
}
type WmsQimenStockoutConfirmOrderLines struct {
	OutBizCode    string                                    `xml:"outBizCode"`    //false	O123456	外部业务编码
	OrderLineNo   string                                    `xml:"orderLineNo"`   //false	1	单据行号
	ItemCode      string                                    `xml:"itemCode"`      //true	SH123456	商品编码
	ItemId        string                                    `xml:"itemId"`        //false	Q123456	商品仓储系统编码
	ItemName      string                                    `xml:"itemName"`      //false	小都进	商品名称
	InventoryType string                                    `xml:"inventoryType"` //false	ZP	库存类型
	ActualQty     json.Number                               `xml:"actualQty"`     //!true	11	实发商品数量
	BatchCode     string                                    `xml:"batchCode"`     //false	P12	批次编号
	ProductDate   string                                    `xml:"productDate"`   //false	2015-09-12	生产日期
	ExpireDate    string                                    `xml:"expireDate"`    //false	2015-09-12	过期日期
	ProduceCode   string                                    `xml:"produceCode"`   //false	P23	生产批号
	Batchs        []WmsQimenStockoutConfirmOrderLinesBatchs `xml:"batchs"`        //false		batchs
	Unit          string                                    `xml:"unit"`          //false	个/盒/箱等	单位
	SnList        WmsQimenStockoutConfirmOrderSnList        `xml:"snList"`        //false		snList
}
type WmsQimenStockoutConfirmOrderLinesBatchs struct {
	BatchCode     string      `xml:"batchCode"`     //false	P234	批次编号
	ProductDate   string      `xml:"productDate"`   //false	2015-09-12	生产日期
	ExpireDate    string      `xml:"expireDate"`    //false	2015-09-12	过期日期
	ProduceCode   string      `xml:"produceCode"`   //false	P23456	生产批号
	InventoryType string      `xml:"inventoryType"` //false	ZP	库存类型
	ActualQty     json.Number `xml:"actualQty"`     //false	123	实发数量
}
type WmsQimenStockoutConfirmOrderSnList struct {
	Sn []string `xml:"sn"` //false		sn
}

//奇门出库单结构体
type WmsQimenStockout struct {
	XMLName       xml.Name                       `xml:"request"`
	DeliveryOrder *WmsQimenStockoutDeliveryOrder `xml:"deliveryOrder"` //DeliveryOrder	false		出库单信息
	OrderLines    *WmsQimenStockOrderLine        `xml:"orderLines"`    //OrderLine[]	false		单据信息
	ExtendProps   map[string]interface{}         `xml:"extendProps"`   //Map	false		扩展属性
}
type WmsQimenStockOrderLine struct {
	OrderLine []*QimenOrderLines `xml:"orderLine"`
}
type WmsQimenStockoutDeliveryOrder struct {
	TotalOrderLines      json.Number                                  `xml:"totalOrderLines"`      //false	12	单据总行数(当单据需要分多个请求发送时;发送方需要将totalOrderLines填入;接收方收到后;根据实际接收到的条数和totalOrderLines进行比对;如果小于;则继续等待接收请求。如果等于;则表示该单据的所有请求发送完成.)
	DeliveryOrderCode    string                                       `xml:"deliveryOrderCode"`    //true	TB1234	出库单号
	OrderType            string                                       `xml:"orderType"`            //true	PTCK	出库单类型(PTCK=普通出库单;DBCK=调拨出库;B2BCK=B2B出库;QTCK=其他出库;CGTH=采购退货出库单;XNCK=虚拟出库单, JITCK=唯品出库)
	RelatedOrders        []*WmsQimenStockoutDeliveryOrderRelatedOrder `xml:"relatedOrders"`        //false		关联单据信息
	WarehouseCode        string                                       `xml:"warehouseCode"`        //true	CK1234	仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	CreateTime           string                                       `xml:"createTime"`           //true	2016-09-09 12:00:00	出库单创建时间(YYYY-MM-DD HH:MM:SS)
	ScheduleDate         string                                       `xml:"scheduleDate"`         //false	2017-09-09	要求出库时间(YYYY-MM-DD)
	LogisticsCode        string                                       `xml:"logisticsCode"`        //false	SF	物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通 、ZTO=中通(ZTO)、HTKY=百世汇通、UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、AMAZON=亚马逊物流、OTHER=其他;只传英文编码)
	LogisticsName        string                                       `xml:"logisticsName"`        //false	顺丰	物流公司名称(包括干线物流公司等)
	SupplierCode         string                                       `xml:"supplierCode"`         //false	TB	供应商编码
	SupplierName         string                                       `xml:"supplierName"`         //false	淘宝	供应商名称
	TransportMode        string                                       `xml:"transportMode"`        //false	自提	提货方式(到仓自提、快递、干线物流)
	PickerInfo           *QimenPickerInfo                             `xml:"pickerInfo"`           //false		提货人信息
	SenderInfo           *QimenSenderInfo                             `xml:"senderInfo"`           //false		发件人信息
	ReceiverInfo         *QimenReceiverInfo                           `xml:"receiverInfo"`         //false		收件人信息
	Remark               string                                       `xml:"remark"`               //false	备注信息	备注
	OrderSourceType      string                                       `xml:"orderSourceType"`      //false	VIP	出库单渠道类型,VIP=唯品会，FX=分销 ，SHOP=门店
	ReceivingTime        string                                       `xml:"receivingTime"`        //false	2016-09-09 12:00:00	到货时间(YYYY-MM-DD HH:MM:SS)
	ShippingTime         string                                       `xml:"shippingTime"`         //false	2016-09-09 12:00:00	送货时间(YYYY-MM-DD HH:MM:SS)
	TargetWarehouseName  string                                       `xml:"targetWarehouseName"`  //false	入库仓库名称, string (50)	入库仓库名称, string (50)
	TargetWarehouseCode  string                                       `xml:"targetWarehouseCode"`  //false	入库仓库编码, string (50) ，统仓统配等无需ERP指定仓储编码的情况填OTHER	入库仓库编码, string (50) ，统仓统配等无需ERP指定仓储编码的情况填OTHER
	TargetEntryOrderCode string                                       `xml:"targetEntryOrderCode"` //false	关联的入库单号（ERP分配）	关联的入库单号（ERP分配）
	WarehouseName        string                                       `xml:"warehouseName"`        //false	123	仓库名称
}
type WmsQimenStockoutDeliveryOrderRelatedOrder struct {
	OrderType string `xml:"orderType"` //false	CG	关联的订单类型(CG=采购单;DB=调拨单;CK=出库单;RK=入库单;只传英文编码)
	OrderCode string `xml:"orderCode"` //false	GL1234	关联的订单编号
}

//奇门提货人信息通用结构体
type QimenPickerInfo struct {
	Company string `xml:"company"` //false	天猫	公司名称
	Name    string `xml:"name"`    //false	老王	姓名
	Tel     string `xml:"tel"`     //false	897765	固定电话
	Mobile  string `xml:"mobile"`  //false	123421234	移动电话
	Id      string `xml:"id"`      //false	1232344322	证件号
	CarNo   string `xml:"carNo"`   //false	XA1234	车牌号
}

//奇门发件人通用结构体
type QimenSenderInfo struct {
	Company       string `xml:"company"`       //false	淘宝	公司名称
	Name          string `xml:"name"`          //true	老王	姓名
	ZipCode       string `xml:"zipCode"`       //false	043300	邮编
	Tel           string `xml:"tel"`           //false	81020340	固定电话
	Mobile        string `xml:"mobile"`        //true	13214567869	移动电话
	Email         string `xml:"email"`         //false	345@gmail.com	电子邮箱
	CountryCode   string `xml:"countryCode"`   //false	051532	国家二字码
	Province      string `xml:"province"`      //true	浙江省	省份
	City          string `xml:"city"`          //true	杭州	城市
	Area          string `xml:"area"`          //false	余杭	区域
	Town          string `xml:"town"`          //false	横加桥	村镇
	DetailAddress string `xml:"detailAddress"` //true	杭州市余杭区989号	详细地址
	Id            string `xml:"id"`            //false	476543213245733	证件号
}

//奇门收件人通用结构体
type QimenReceiverInfo struct {
	Company       string `xml:"company"`       //false	淘宝	公司名称
	Name          string `xml:"name"`          //true	老王	姓名
	ZipCode       string `xml:"zipCode"`       //false	043300	邮编
	Tel           string `xml:"tel"`           //false	808786543	固定电话
	Mobile        string `xml:"mobile"`        //true	13423456785	移动电话
	IdType        string `xml:"idType"`        //false	1	收件人证件类型(1-身份证、2-军官证、3-护照、4-其他)
	IdNumber      string `xml:"idNumber"`      //false	1234567	收件人证件号码
	Email         string `xml:"email"`         //false	878987654@qq.com	电子邮箱
	CountryCode   string `xml:"countryCode"`   //false	045565	国家二字码
	Province      string `xml:"province"`      //true	浙江省	省份
	City          string `xml:"city"`          //true	杭州	城市
	Area          string `xml:"area"`          //false	余杭	区域
	Town          string `xml:"town"`          //false	横加桥	村镇
	DetailAddress string `xml:"detailAddress"` //true	杭州市余杭区989号	详细地址
	Id            string `xml:"id"`            //false	4713242536	证件号
}
type QimenOrderLines struct {
	OutBizCode    string                 `xml:"outBizCode"`    //false	OB1234	外部业务编码(消息ID;用于去重;当单据需要分批次发送时使用)
	OrderLineNo   string                 `xml:"orderLineNo"`   //false	11	单据行号
	OwnerCode     string                 `xml:"ownerCode"`     //true	H1234	货主编码
	ItemCode      string                 `xml:"itemCode"`      //true	I1234	商品编码
	ItemId        string                 `xml:"itemId"`        //false	W1234	仓储系统商品编码
	InventoryType string                 `xml:"inventoryType"` //false	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS= 箱损;ZT=在途库存;默认为查所有类型的库存)
	ItemName      string                 `xml:"itemName"`      //false	淘公仔	商品名称
	PlanQty       json.Number            `xml:"planQty"`       //true	11	应发商品数量
	BatchCode     string                 `xml:"batchCode"`     //false	123	批次编码
	ProductDate   string                 `xml:"productDate"`   //false	2016-07-06	生产日期(YYYY-MM-DD)
	ExpireDate    string                 `xml:"expireDate"`    //false	2016-07-06	过期日期(YYYY-MM-DD)
	ProduceCode   string                 `xml:"produceCode"`   //false	P11233	生产批号
	PlatformCode  string                 `xml:"platformCode"`  //false	123456789	交易平台商品编码
	Unit          string                 `xml:"unit"`          //false	个/箱/盒等	单位
	ExtendProps   map[string]interface{} `xml:"extendProps"`   //false		扩展属性
}
