// https://qimen.taobao.com/qimen/index.htm#/api-doc?fromPage=api-list&&apiServiceId=&officialApiId=25995&_k=crdysi
// taobao.qimen.entryorder.confirm 入库单确认接口

package qimen

import (
	"encoding/xml"
	"fmt"

	"github.com/yellbuy/go-ec-openapi/common"
)

// 采购单回传
func (client *Client) EntryOrderConfirm(dto *EntryOrderConfirmReqDto) (body []byte, err error) {
	var bytes []byte
	bytes, err = xml.Marshal(dto)

	req := make(map[string]interface{})
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		body = nil
		return
	}

	// 通过奇门代理平台
	method := "taobao.qimen.entryorder.confirm"
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}
func (client *Client) EntryOrderConfirmV2(dto *EntryOrderConfirmReqDtoV2) (body []byte, err error) {
	var bytes []byte
	bytes, err = xml.Marshal(dto)

	req := make(map[string]interface{})
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		body = nil
		return
	}

	// 通过奇门代理平台
	method := "taobao.qimen.entryorder.confirm"
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}
func (client *Client) ReturnOrderConfirmV2(dto *ReturnOrderConfirmRequest) (body []byte, err error) {
	var bytes []byte
	bytes, err = xml.Marshal(dto)

	req := make(map[string]interface{})
	//fmt.Println("bizcontent：", string(bizcontent))
	params, resErr := common.InterfaceToParameter(req)
	if resErr != nil {
		fmt.Println(resErr)
		err = resErr
		body = nil
		return
	}

	// 通过奇门代理平台
	method := "taobao.qimen.returnorder.confirm"
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}

type ReturnOrderConfirmRequest struct {
	XMLName     xml.Name                              `xml:"request"`
	ReturnOrder *ReturnOrderConfirmRequestReturnOrder `xml:"returnOrder"` //ReturnOrder	false		退货单信息
	OrderLines  *ReturnOrderConfirmRequestOrderLines  `xml:"orderLines"`  //OrderLine[]	false		订单信息
	// extendProps `xml:"extendProps"`                   //Map	false		扩展属性
}
type ReturnOrderConfirmRequestOrderLines struct {
	OrderLine []*ReturnOrderConfirmRequestOrderLine `xml:"orderLine"`
}
type ReturnOrderConfirmRequestOrderLine struct {
	Remark             string                                      `xml:"remark"`             //String	false	备注	备注
	OrderLineNo        string                                      `xml:"orderLineNo"`        //String	false	D1234	单据行号
	SourceOrderCode    string                                      `xml:"sourceOrderCode"`    //String	false	PD1224	交易平台订单
	SubSourceOrderCode string                                      `xml:"subSourceOrderCode"` //String	false	PL1234	交易平台子订单编码
	OwnerCode          string                                      `xml:"ownerCode"`          //!String	true	HZ1234	货主编码
	ItemCode           string                                      `xml:"itemCode"`           //!String	true	I1234	商品编码
	ItemId             string                                      `xml:"itemId"`             //String	false	CK1234	仓储系统商品编码(条件为提供后端（仓储系统）商品编码的仓储系统)
	InventoryType      string                                      `xml:"inventoryType"`      //String	false	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;默认为ZP)
	PlanQty            string                                      `xml:"planQty"`            //!Number	true	12	应收商品数量
	BatchCode          string                                      `xml:"batchCode"`          //String	false	P123	批次编码
	ProductDate        string                                      `xml:"productDate"`        //String	false	2016-09-09	生产日期(YYYY-MM-DD)
	ExpireDate         string                                      `xml:"expireDate"`         //String	false	2016-09-09	过期日期(YYYY-MM-DD)
	ProduceCode        string                                      `xml:"produceCode"`        //String	false	P1234	生产批号
	Batchs             []*ReturnOrderConfirmRequestOrderLineBatchs `xml:"batchs"`             //Batch[]	false		批次信息
	QrCode             string                                      `xml:"qrCode"`             //String	false	1;1;1	商品的二维码(类似电子产品的SN码;用来进行商品的溯源;多个二维码之间用分号;隔开)
	ActualQty          string                                      `xml:"actualQty"`          //!String	true	12	实收商品数量
	SnList             *ReturnOrderConfirmRequestOrderLineSnList   `xml:"snList"`             //SnList	false		sn列表
}
type ReturnOrderConfirmRequestOrderLineSnList struct {
	Sn []string `xml:"sn"` //String[]	false		sn编码
}
type ReturnOrderConfirmRequestOrderLineBatchs struct {
	Remark        string `xml:"remark"`        //String	false	备注	备注
	BatchCode     string `xml:"batchCode"`     //String	false	P1234	批次编号
	ProductDate   string `xml:"productDate"`   //String	false	2016-09-09	生产日期(YYYY-MM-DD)
	ExpireDate    string `xml:"expireDate"`    //String	false	2016-09-09	过期日期(YYYY-MM-DD)
	ProduceCode   string `xml:"produceCode"`   //String	false	P1234	生产批号
	InventoryType string `xml:"inventoryType"` //String	false	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;默认为ZP)
	ActualQty     string `xml:"actualQty"`     //Number	false	12	实收数量(要求batchs节点下所有的实收数量之和等于orderline中的实收数量)
}
type ReturnOrderConfirmRequestReturnOrder struct {
	ReturnOrderCode  string                                          `xml:"returnOrderCode"`  //!String	true	R1234	ERP的退货入库单编码
	RreturnOrderId   string                                          `xml:"returnOrderId"`    //String	false	R1234	仓库系统订单编码
	WarehouseCode    string                                          `xml:"warehouseCode"`    //!String	true	W1234	仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	OutBizCode       string                                          `xml:"outBizCode"`       //String	false	OZ1234	外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不会 被重复处理)
	OrderType        string                                          `xml:"orderType"`        //String	false	THRK	单据类型(THRK=退货入库;HHRK=换货入库;只传英文编码)
	OrderConfirmTime string                                          `xml:"orderConfirmTime"` //String	false	2016-09-09 12:00:00	确认入库时间(YYYY-MM-DD HH:MM:SS)
	LogisticsCode    string                                          `xml:"logisticsCode"`    //!String	true	SF	物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通、ZTO=中通(ZTO)、HTKY=百世汇通、 UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、 AMAZON=亚马逊物流、OTHER=其他;只传英文编码)
	LogisticsName    string                                          `xml:"logisticsName"`    //String	false	顺丰	物流公司名称
	ExpressCode      string                                          `xml:"expressCode"`      //String	false	YD1234	运单号
	ReturnReason     string                                          `xml:"returnReason"`     //String	false	破损退货	退货原因
	Remark           string                                          `xml:"remark"`           //String	false	备注信息	备注
	SenderInfo       *ReturnOrderConfirmRequestReturnOrderSenderInfo `xml:"senderInfo"`       //SenderInfo	false		发件人信息
	ConfirmType      int                                             `xml:"confirmType"`
}
type ReturnOrderConfirmRequestReturnOrderSenderInfo struct {
	company       string `xml:"company"`       //String	false	淘宝	公司名称
	name          string `xml:"name"`          //!String	true	老王	姓名
	zipCode       string `xml:"zipCode"`       //String	false	043300	邮编
	tel           string `xml:"tel"`           //String	false	81020340	固定电话
	mobile        string `xml:"mobile"`        //!String	true	13214567869	移动电话
	email         string `xml:"email"`         //String	false	345@gmail.com	电子邮箱
	countryCode   string `xml:"countryCode"`   //String	false	051532	国家二字码
	province      string `xml:"province"`      //!String	true	浙江省	省份
	city          string `xml:"city"`          //!String	true	杭州	城市
	area          string `xml:"area"`          //String	false	余杭	区域
	town          string `xml:"town"`          //String	false	横加桥	村镇
	detailAddress string `xml:"detailAddress"` //String	true	杭州市余杭区989号	详细地址
	remark        string `xml:"remark"`        //String	false	备注	备注
}
type EntryOrderConfirmReqDto struct {
	XMLName    xml.Name        `xml:"request"`
	Text       string          `xml:",chardata"`
	EntryOrder EntryOrder      `xml:"entryOrder"`
	OrderLines EntryOrderLines `xml:"orderLines"`
}
type EntryOrderConfirmReqDtoV2 struct {
	XMLName    xml.Name           `xml:"request"`
	Text       string             `xml:",chardata"`
	EntryOrder *EntryOrderV2      `xml:"entryOrder"`
	OrderLines *EntryOrderLinesV2 `xml:"orderLines"`
}
type EntryOrderV2 struct {
	OrderCode             string `xml:"orderCode"`             //String	false	订单编码	订单编码
	OrderId               string `xml:"orderId"`               //String	false	后端订单id	后端订单id
	OrderType             string `xml:"orderType"`             //String	false	订单类型	订单类型
	WarehouseName         string `xml:"warehouseName"`         //String	false	仓库名称	仓库名称
	TotalOrderLines       string `xml:"totalOrderLines"`       //Number	false	12	单据总行数(当单据需要分多个请求发送时;发送方需要将totalOrderLines填入;接收方收到后;根据实际接收到的 条数和 totalOrderLines进行比对;如果小于;则继续等待接收请求。如果等于;则表示该单据的所有请求发送完 成)
	EntryOrderCode        string `xml:"entryOrderCode"`        //String	true	E1234	入库单号
	OwnerCode             string `xml:"ownerCode"`             //String	true	O1234	货主编码
	PurchaseOrderCode     string `xml:"purchaseOrderCode"`     //String	false	C123455	采购单号(当orderType=CGRK时使用)
	WarehouseCode         string `xml:"warehouseCode"`         //String	true	W1234	仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	EntryOrderId          string `xml:"entryOrderId"`          //String	false	E1234	仓储系统入库单ID
	EntryOrderType        string `xml:"entryOrderType"`        //String	false	SCRK	入库单类型(SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入 库;B2BRK=B2B入库)
	OutBizCode            string `xml:"outBizCode"`            //String	true	O1234	外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求 不会被重复处理)
	ConfirmType           string `xml:"confirmType"`           //Number	false	0	支持出入库单多次收货(多次收货后确认时:0:表示入库单最终状态确认;1:表示入库单中间状态确认;每次入库传入的数量为增 量;特殊情况;同一入库单;如果先收到0;后又收到1;允许修改收货的数量)
	Status                string `xml:"status"`                //String	true	NEW	入库单状态(NEW-未开始处理;ACCEPT-仓库接单;PARTFULFILLED-部分收货完成;FULFILLED-收货完成;EXCEPTION-异 常;CANCELED-取消;CLOSED-关闭;REJECT-拒单;CANCELEDFAIL-取消失败;只传英文编码)
	OperateTime           string `xml:"operateTime"`           //String	false	2017-09-09 12:00:00	操作时间(YYYY-MM-DD HH:MM:SS;当status=FULFILLED;operateTime为入库时间)
	Remark                string `xml:"remark"`                //String	false	备注信息	备注
	Freight               string `xml:"freight"`               //String	false	奇门仓储字段,说明,string(50),,	邮费
	SubOrderType          string `xml:"subOrderType"`          //String	false	hss	入库单确认的其他入库子类型，用于entryOrderType设置为其他入库时设置
	ResponsibleDepartment string `xml:"responsibleDepartment"` //String	false	财务部	该笔入库单的费用承担部门或责任部门
	ShopNick              string `xml:"shopNick"`              //String	true	旗舰店	店铺名称
	ShopCode              string `xml:"shopCode"`              //String	true	ssss	店铺编码
}
type EntryOrderLinesV2 struct {
	OrderLine []*EntryOrderLineV2 `xml:"orderLine"`
}
type EntryOrderLineV2 struct {
	OutBizCode    string      `xml:"outBizCode"`    //String	false	O123	外部业务编码(消息ID;用于去重;当单据需要分批次发送时使用)
	OrderLineNo   string      `xml:"orderLineNo"`   //String	false	EL123	入库单的行号
	OwnerCode     string      `xml:"ownerCode"`     //String	true	O123	货主编码
	ItemCode      string      `xml:"itemCode"`      //String	true	I123	商品编码
	ItemId        string      `xml:"itemId"`        //String	false	CI123	仓储系统商品ID
	ItemName      string      `xml:"itemName"`      //String	false	IN123	商品名称
	PlanQty       string      `xml:"planQty"`       //Number	true	12	应收商品数量
	InventoryType string      `xml:"inventoryType"` //String	false	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;默认为ZP)
	ActualQty     string      `xml:"actualQty"`     //Number	true	12	实收数量
	ProductDate   string      `xml:"productDate"`   //String	false	2017-09-09	商品生产日期(YYYY-MM-DD)
	ExpireDate    string      `xml:"expireDate"`    //String	false	2017-09-09	商品过期日期(YYYY-MM-DD)
	ProduceCode   string      `xml:"produceCode"`   //String	false	P1234	生产批号
	BatchCode     string      `xml:"batchCode"`     //String	false	PCODE123	批次编码
	Batchs        []*BatchsV2 `xml:"batchs"`        //Batch[]	false		批次信息
	Remark        string      `xml:"remark"`        //String	false	备注信息	备注
	Unit          string      `xml:"unit"`          //String	false	盒/箱/个等	单位
	SnList        []*SnListV2 `xml:"snList"`        //SnList	false		sn列表
}
type BatchsV2 struct {
	BatchCode     string `xml:"batchCode"`     //String	false	P1234	批次编号
	ProductDate   string `xml:"productDate"`   //String	false	2016-09-09	生产日期(YYYY-MM-DD)
	ExpireDate    string `xml:"expireDate"`    //String	false	2016-09-09	过期日期(YYYY-MM-DD)
	ProduceCode   string `xml:"produceCode"`   //String	false	P1234	生产批号
	InventoryType string `xml:"inventoryType"` //String	false	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS= 箱损;默认为ZP;)
	ActualQty     string `xml:"actualQty"`     //Number	false	12	实收数量(要求batchs节点下所有的实收数量之和等于orderline中的实收数量)
}
type SnListV2 struct {
	Sn []string `xml:"sn"` //String[]	false		sn编码
}

type EntryOrder struct {
	Text            string `xml:",chardata"`
	TotalOrderLines string `xml:"totalOrderLines"`
	EntryOrderCode  string `xml:"entryOrderCode"`
	OwnerCode       string `xml:"ownerCode"`
	WarehouseCode   string `xml:"warehouseCode"`
	EntryOrderId    string `xml:"entryOrderId"`
	EntryOrderType  string `xml:"entryOrderType"`
	OutBizCode      string `xml:"outBizCode"`
	ConfirmType     int    `xml:"confirmType"`
	Status          string `xml:"status"`
	Freight         string `xml:"freight"`
	OperateTime     string `xml:"operateTime"`
	Remark          string `xml:"remark"`
}
type EntryOrderLines struct {
	Text          string           `xml:",chardata"`
	OrderLineList []EntryOrderLine `xml:"orderLine"`
}

type EntryOrderLine struct {
	Text        string `xml:",chardata"`
	OutBizCode  string `xml:"outBizCode"`
	OrderLineNo string `xml:"orderLineNo"`
	OwnerCode   string `xml:"ownerCode"`
	ItemCode    string `xml:"itemCode"`
	ItemId      string `xml:"itemId"`
	SnList      struct {
		Text string `xml:",chardata"`
		Sn   string `xml:"sn"`
	} `xml:"snList"`
	ItemName      string `xml:"itemName"`
	InventoryType string `xml:"inventoryType"`
	PlanQty       int    `xml:"planQty"`
	ActualQty     int    `xml:"actualQty"`
	BatchCode     string `xml:"batchCode"`
	ProductDate   string `xml:"productDate"`
	ExpireDate    string `xml:"expireDate"`
	ProduceCode   string `xml:"produceCode"`
	// Batchs        struct {
	// 	Text  string `xml:",chardata"`
	// 	Batch struct {
	// 		Text          string `xml:",chardata"`
	// 		BatchCode     string `xml:"batchCode"`
	// 		ProductDate   string `xml:"productDate"`
	// 		ExpireDate    string `xml:"expireDate"`
	// 		ProduceCode   string `xml:"produceCode"`
	// 		InventoryType string `xml:"inventoryType"`
	// 		ActualQty     string `xml:"actualQty"`
	// 	} `xml:"batch"`
	// } `xml:"batchs"`
	Remark string `xml:"remark"`
}
