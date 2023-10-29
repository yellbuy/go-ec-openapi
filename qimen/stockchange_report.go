package qimen

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/yellbuy/go-ec-openapi/common"
)

// 奇门库存异动通知接口
func (client *Client) QimenStockChangeReport(dto *QimenStocChangeReportRequest) (body []byte, err error) {
	var bytes []byte
	bytes, err = xml.Marshal(dto)
	if err != nil {
		return nil, err
	}
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
	method := "taobao.qimen.stockchange.report"
	//fmt.Println("奇门出库单推送报文", string(bytes))
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}

// 奇门库存异动通知接口
func (client *Client) QimenInventoryReport(dto *QimenInventoryReport) (body []byte, err error) {
	var bytes []byte
	bytes, err = xml.Marshal(dto)
	if err != nil {
		return nil, err
	}
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
	method := "taobao.qimen.inventory.report"
	//fmt.Println("奇门出库单推送报文", string(bytes))
	body, err = client.Execute(method, params, bytes)
	if err != nil {
		fmt.Println(method, err)
	}
	return
}

type QimenStocChangeReportRequest struct {
	XMLName     xml.Name                    `xml:"request"`
	Items       *QimenStocChangeReportItems `xml:"items"`       //Item[]	false		item
	ExtendProps interface{}                 `xml:"extendProps"` //Map	false		扩展属性
}
type QimenStocChangeReportItem struct {
	OwnerCode     string                            `xml:"ownerCode"`     //!true	H1234	货主编码
	WarehouseCode string                            `xml:"warehouseCode"` //!true	CH1234	仓库编码
	OrderCode     string                            `xml:"orderCode"`     //!true	OR1234	引起异动的单据编码
	OrderType     string                            `xml:"orderType"`     //false	JYCK	单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;B2BRK=B2B入 库;B2BCK=B2B出库;QTCK=其他出库;SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK= 调拨入库;QTRK= 其他入 库;XTRK= 销退入库;HHRK= 换货入库;CNJG= 仓内加工单)
	OutBizCode    string                            `xml:"outBizCode"`    //!true	OZ1234	外部业务编码(消息ID;用于去重;用来保证因为网络等原因导致重复传输;请求不会被重复处理)
	ItemCode      string                            `xml:"itemCode"`      //!true	I1234	商品编码
	ItemId        string                            `xml:"itemId"`        //false	CH1234	仓储系统商品ID
	InventoryType string                            `xml:"inventoryType"` //false	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存)
	Quantity      json.Number                       `xml:"quantity"`      //!true	12	商品变化量(可为正为负)
	BatchCode     string                            `xml:"batchCode"`     //false	PC1234	批次编码
	ProductDate   string                            `xml:"productDate"`   //false	2017-09-09	商品生产日期(YYYY-MM-DD)
	ExpireDate    string                            `xml:"expireDate"`    //false	2017-09-09	商品过期日期(YYYY-MM-DD)
	ProduceCode   string                            `xml:"produceCode"`   //false	P1234	生产批号
	ChangeTime    string                            `xml:"changeTime"`    //false	2017-09-09 12:00:00	异动时间(YYYY-MM-DD HH:MM:SS)
	Batchs        []*QimenStocChangeReportItemBatch `xml:"batchs"`        //false		batchs
	Remark        string                            `xml:"remark"`        //false	备注信息	备注
	IsLocked      string                            `xml:"isLocked"`      //false	Y	isLocked
}
type QimenStocChangeReportItems struct {
	Item []*QimenStocChangeReportItem `xml:"item"`
}
type QimenStocChangeReportItemBatch struct {
	BatchCode     string      `xml:"batchCode"`     //false	PC1234	批次编号
	ProductDate   string      `xml:"productDate"`   //false	2017-09-09	生产日期(YYYY-MM-DD)
	ExpireDate    string      `xml:"expireDate"`    //false	2017-09-09	过期日期(YYYY-MM-DD)
	ProduceCode   string      `xml:"produceCode"`   //false	PD1234	生产批号
	InventoryType string      `xml:"inventoryType"` //false	ZP	库存类型(ZP=正品;CC=残次;JS=机损 XS= 箱损;ZT=在途库存)
	Quantity      json.Number `xml:"quantity"`      //false	12	异动数量(要求batchs节点下所有的异动数量之和等于orderline中的异动数量)
}

type QimenInventoryReport struct {
	XMLName        xml.Name                   `xml:"request"`
	TotalPage      int                        `xml:"totalPage"`      //!Number	true	12	总页数
	CurrentPage    int                        `xml:"currentPage"`    //!Number	true	1	当前页(从1开始)
	PageSize       int                        `xml:"pageSize"`       //!Number	true	12	每页记录的条数
	WarehouseCode  string                     `xml:"warehouseCode"`  //!String	true	W1234	仓库编码
	CheckOrderCode string                     `xml:"checkOrderCode"` //!String	true	P1234	盘点单编码
	CheckOrderId   string                     `xml:"checkOrderId"`   //String	false	PS1234	仓储系统的盘点单编码
	OwnerCode      string                     `xml:"ownerCode"`      //!String	true	H1234	货主编码
	CheckTime      string                     `xml:"checkTime"`      //String	false	2016-09-09 12:00:00	盘点时间(YYYY-MM-DD HH:MM:SS)
	OutBizCode     string                     `xml:"outBizCode"`     //!String	true	OZ1234	外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理)
	Remark         string                     `xml:"remark"`         //String	false	备注信息	备注
	Items          *QimenInventoryReportItems `xml:"items"`          //Item[]	false		商品库存信息列表
	AdjustType     string                     `xml:"adjustType"`     //!String	true	CHECK	变动类型：CHECK=盘点 ADJUST=调整
}
type QimenInventoryReportItems struct {
	Item []*QimenInventoryReport_Item `xml:"item"`
}
type QimenInventoryReport_Item struct {
	ItemCode      string `xml:"itemCode"`      //!String	true	I1234	商品编码
	ItemId        string `xml:"itemId"`        //String	false	ID1234	仓储系统商品ID
	InventoryType string `xml:"inventoryType"` //String	false	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS= 箱损;ZT=在途库存;默认为ZP)
	Quantity      int    `xml:"quantity"`      //!Number	true	12	盘盈盘亏商品变化量(盘盈为正数;盘亏为负数)
	BatchCode     string `xml:"batchCode"`     //String	false	P1234	批次编码
	ProductDate   string `xml:"productDate"`   //String	false	2016-09-09	商品生产日期(YYYY-MM-DD)
	ExpireDate    string `xml:"expireDate"`    //String	false	2016-09-09	商品过期日期(YYYY-MM-DD)
	ProduceCode   string `xml:"produceCode"`   //String	false	P1234	生产批号
	SnCode        string `xml:"snCode"`        //String	false	X1234	商品序列号
	Remark        string `xml:"remark"`        //String	false	备注信息	备注
	TotalQty      int    `xml:"totalQty"`      //Number	false	100	库存商品总量
}
