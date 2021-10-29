package qimen

import (
	"encoding/json"
	"encoding/xml"
)

func StockOutCreateParse(body []byte) (res *StockOutCreateDto, err error) {
	res = new(StockOutCreateDto)
	err = xml.Unmarshal(body, res)
	return
}

//出库单创建接口
type StockOutCreateDto struct {
	XMLName       xml.Name               `xml:"request"`
	DeliveryOrder *StockOutDeliveryOrder `xml:"deliveryOrder"` //false		出库单信息
	OrderLines    *StockOutOrderLines    `xml:"orderLines"`    //OrderLine[]	false		单据信息
	ExtendProps   map[string]interface{} `xml:"extendProps"`   //Map	false		扩展属性
}
type StockOutDeliveryOrder struct {
	OwnerCode            string                  `xml:"ownerCode"`
	TotalOrderLines      json.Number             `xml:"totalOrderLines"`      //false	12	单据总行数(当单据需要分多个请求发送时;发送方需要将totalOrderLines填入;接收方收到后;根据实际接收到的条数和totalOrderLines进行比对;如果小于;则继续等待接收请求。如果等于;则表示该单据的所有请求发送完成.)
	DeliveryOrderCode    string                  `xml:"deliveryOrderCode"`    //true	TB1234	出库单号
	OrderType            string                  `xml:"orderType"`            //true	PTCK	出库单类型(PTCK=普通出库单;DBCK=调拨出库;B2BCK=B2B出库;QTCK=其他出库;CGTH=采购退货出库单;XNCK=虚拟出库单, JITCK=唯品出库)
	RelatedOrders        []*StockOutRelatedOrder `xml:"relatedOrders"`        //false		关联单据信息
	WarehouseCode        string                  `xml:"warehouseCode"`        //true	CK1234	仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	CreateTime           string                  `xml:"createTime"`           //true	2016-09-09 12:00:00	出库单创建时间(YYYY-MM-DD HH:MM:SS)
	ScheduleDate         string                  `xml:"scheduleDate"`         //false	2017-09-09	要求出库时间(YYYY-MM-DD)
	LogisticsCode        string                  `xml:"logisticsCode"`        //false	SF	物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通 、ZTO=中通(ZTO)、HTKY=百世汇通、UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、AMAZON=亚马逊物流、OTHER=其他;只传英文编码)
	LogisticsName        string                  `xml:"logisticsName"`        //false	顺丰	物流公司名称(包括干线物流公司等)
	SupplierCode         string                  `xml:"supplierCode"`         //false	TB	供应商编码
	SupplierName         string                  `xml:"supplierName"`         //false	淘宝	供应商名称
	TransportMode        string                  `xml:"transportMode"`        //false	自提	提货方式(到仓自提、快递、干线物流)
	PickerInfo           *StockOutPickerInfo     `xml:"pickerInfo"`           //false		提货人信息
	SenderInfo           *StockOutSenderInfo     `xml:"senderInfo"`           //false		发件人信息
	ReceiverInfo         *StockOutReceiverInfo   `xml:"receiverInfo"`         //false		收件人信息
	Remark               string                  `xml:"remark"`               //false	备注信息	备注
	OrderSourceType      string                  `xml:"orderSourceType"`      //false	VIP	出库单渠道类型,VIP=唯品会，FX=分销 ，SHOP=门店
	ReceivingTime        string                  `xml:"receivingTime"`        //false	2016-09-09 12:00:00	到货时间(YYYY-MM-DD HH:MM:SS)
	ShippingTime         string                  `xml:"shippingTime"`         //false	2016-09-09 12:00:00	送货时间(YYYY-MM-DD HH:MM:SS)
	TargetWarehouseName  string                  `xml:"targetWarehouseName"`  //false	入库仓库名称, string (50)	入库仓库名称, string (50)
	TargetWarehouseCode  string                  `xml:"targetWarehouseCode"`  //false	入库仓库编码, string (50) ，统仓统配等无需ERP指定仓储编码的情况填OTHER	入库仓库编码, string (50) ，统仓统配等无需ERP指定仓储编码的情况填OTHER
	TargetEntryOrderCode string                  `xml:"targetEntryOrderCode"` //false	关联的入库单号（ERP分配）	关联的入库单号（ERP分配）
	WarehouseName        string                  `xml:"warehouseName"`        //false	123	仓库名称
}
type StockOutRelatedOrder struct {
	OrderType string `xml:"orderType"` //false	CG	关联的订单类型(CG=采购单;DB=调拨单;CK=出库单;RK=入库单;只传英文编码)
	OrderCode string `xml:"orderCode"` //false	GL1234	关联的订单编号
}
type StockOutPickerInfo struct {
	Company string `xml:"company"` //false	天猫	公司名称
	Name    string `xml:"name"`    //false	老王	姓名
	Tel     string `xml:"tel"`     //false	897765	固定电话
	Mobile  string `xml:"mobile"`  //false	123421234	移动电话
	Id      string `xml:"id"`      //false	1232344322	证件号
	CarNo   string `xml:"carNo"`   //false	XA1234	车牌号
}
type StockOutSenderInfo struct {
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
type StockOutReceiverInfo struct {
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
type StockOutOrderLines struct {
	OrderLine []*StockOutOrderLine `xml:"orderLine"`
}
type StockOutOrderLine struct {
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
