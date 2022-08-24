package common

import "encoding/xml"

//奇门库存查询接收结构体
type QimenInventoryQuery struct {
	XMLName      xml.Name                          `xml:"request"`
	CriteriaList *QimenInventoryQuery_CriteriaList `xml:"criteriaList"`
	ExtendProps  interface{}                       `xml:"extendProps"` //Map	false		扩展属性
	Remark       string                            `xml:"remark"`      //	false	备注	备注
}
type QimenInventoryQuery_Criteria struct {
	WarehouseCode string `xml:"warehouseCode"` //false	W1234	仓库编码
	OwnerCode     string `xml:"ownerCode"`     //false	H1234	货主编码
	ItemCode      string `xml:"itemCode"`      //true	I1234	商品编码
	ItemId        string `xml:"itemId"`        //false	C1234	仓储系统商品ID
	InventoryType string `xml:"inventoryType"` //false	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
	Remark        string `xml:"remark"`        //false	备注	备注

}
type QimenInventoryQuery_CriteriaList struct {
	Criteria []*QimenInventoryQuery_Criteria `xml:"criteria"` //Criteria[]	false		查询准则
}

//奇门库存查询返回结构体
type QimenInventoryQueryResponse struct {
	XMLName xml.Name                            `xml:"response"`
	Flag    string                              `xml:"flag"`    //success	响应结果:success|failure
	Code    string                              `xml:"code"`    //0	响应码
	Message string                              `xml:"message"` //invalid appkey	响应信息
	Items   []*QimenInventoryQueryResponse_Item `xml:"items"`   //		商品的库存信息列表
}
type QimenInventoryQueryResponse_Item struct {
	WarehouseCode string `xml:"warehouseCode"` //	C1234	仓库编码
	ItemCode      string `xml:"itemCode"`      //	I1234	商品编码
	ItemId        string `xml:"itemId"`        //	W1234	仓储系统商品ID
	InventoryType string `xml:"inventoryType"` //	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS= 箱损;ZT=在途库存)
	Quantity      int    `xml:"quantity"`      //	11	未冻结库存数量
	LockQuantity  int    `xml:"lockQuantity"`  //	1	冻结库存数量
	BatchCode     string `xml:"batchCode"`     //	P1234	批次编码
	ProductDate   string `xml:"productDate"`   //	2017-09-09	商品生产日期(YYYY-MM-DD)
	ExpireDate    string `xml:"expireDate"`    //	2017-09-09	商品过期日期(YYYY-MM-DD)
	ProduceCode   string `xml:"produceCode"`   //	P1234	生产批号
}

//奇门订单流水通知结构体
type QimenOrderProcessReportRequest struct {
	XMLName     xml.Name                                `xml:"request"`
	Order       *QimenOrderProcessReportRequest_Order   `xml:"order"`       //false		订单信息
	Process     *QimenOrderProcessReportRequest_Process `xml:"process"`     //false		订单处理信息
	ExtendProps interface{}                             `xml:"extendProps"` //false		扩展属性
	Remark      string                                  `xml:"remark"`      //false		备注
}
type QimenOrderProcessReportRequest_Order struct {
	OrderCode     string `xml:"orderCode"`     //	true	D1234	单据号
	OrderId       string `xml:"orderId"`       //	false	W1234	仓储系统单据号
	OrderType     string `xml:"orderType"`     //	false	JYCK	单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;B2BRK=B2B入 库;B2BCK=B2B出库;QTCK=其他出库;SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK= 调拨入库;QTRK= 其他入 库;XTRK= 销退入库;HHRK= 换货入库;CNJG= 仓内加工单)
	WarehouseCode string `xml:"warehouseCode"` //	false	W1234	仓库编码
	Remark        string `xml:"remark"`        //	false	备注	备注
}
type QimenOrderProcessReportRequest_Process struct {
	ProcessStatus string `xml:"processStatus"` //	true	ACCEPT	单据状态(ACCEPT=仓库接单;PARTFULFILLED-部分收货完成;FULFILLED=收货完成;PRINT = 打印;PICK=捡货;CHECK = 复核 ;PACKAGE= 打包;WEIGH= 称重;READY=待提货;DELIVERED=已发货;REFUSE=买家拒签;EXCEPTION =异常;CLOSED= 关闭;CANCELED= 取 消;REJECT=仓库拒单;SIGN=签收;TMSCANCELED=快递拦截;OTHER=其他;PARTDELIVERED=部分发货完成;TMSCANCELFAILED=快递拦截失败;只传英 文编码)
	OperatorCode  string `xml:"operatorCode"`  //	false	O1234	当前状态操作员编码
	OperatorName  string `xml:"operatorName"`  //	false	老王	当前状态操作员姓名
	OperateTime   string `xml:"operateTime"`   //	false	2016-09-09 12:00:00	当前状态操作时间(YYYY-MM-DD HH:MM:SS)
	OperateInfo   string `xml:"operateInfo"`   //	false	处理中	操作内容
	Remark        string `xml:"remark"`        //	false	备注信息	备注
	ExpressCode   string `xml:"expressCode"`   //	false	123456789	运单号
}
