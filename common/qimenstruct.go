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
type WmsGoodsSkuSyncReturn struct {
	Results []*WmsGoodsSkuSyncReturnItem `json:"results"`
}
type WmsGoodsSkuSyncReturnItem struct {
	Issuccess        int    `json:"issuccess"`        //必填	通用	1	是否成功(0表示失败；1表示成功)	0
	Code             string `json:"code"`             //必填	通用	256	错误编码	Failed
	Message          string `json:"message"`          //可选	通用	256	是否成功	Success
	Platproductid    string `json:"platproductid"`    //可选	通用	32	平台商品ID（平台商品ID、平台子规格ID、外...	11447876487218
	Skuid            string `json:"skuid"`            //可选	通用	32	平台子规格ID（平台商品ID、平台子规格ID、...	114478764
	Outerid          string `json:"outerid"`          //可选	通用	32	外部商家编码（平台商品ID、平台子规格ID、...	WR6685851555
	Outskuid         string `json:"outskuid"`         //可选	通用	30	外部商家SKU编号（平台商品ID、平台子规格I...	SK43008558525565
	Quantity         int    `json:"quantity"`         //必填	通用	4	更新后的数量	180
	Transactionid    string `json:"transactionid"`    //可选	亚马逊	45	交易序号	222425696737
	Platstoreid      string `json:"platstoreid"`      //可选	京东到家、有赞	15	门店ID	01
	Whsecode         string `json:"whsecode"`         //可选	融易购、达令网、美囤妈妈、唯品会	15	商品仓库编号	KU002
	Vipcooperationno int    `json:"vipcooperationno"` //可选	唯品会	15	合作编码	23
	Requestid        int    `json:"requestid"`        //选	亚马逊	4	查询批次号，仅限异步模式(如亚马逊)	180
}
type WmsGoodsSkuSync struct {
	Goods []*WmsGoodsSkuSyncGoods `json:"goods"`
}
type WmsGoodsSkuSyncGoods struct {
	Platproductid    string `json:"platproductid"`    //必填	通用	32	平台商品ID	11447876487218
	Skuid            string `json:"skuid"`            //必填	通用	32	平台子规格ID	114478764
	Outerid          string `json:"outerid"`          //必填	通用	32	外部商家编码	WR6685851555
	Quantity         int    `json:"quantity"`         //必填	通用	4	库存数量	180
	Syncstocktype    string `json:"syncstocktype"`    //必填	通用	25	库存更新方式(全量更新=JH_01，增量更新=JH...	JH_01
	Shoptype         string `json:"shoptype"`         //可选	通用	25	店铺类型(普通=JH_001，分销=JH_002，经销=...	JH_001
	Outskuid         string `json:"outskuid"`         //必填	通用	30	外部商家SKU编号	SK43008558525565
	Producttype      string `json:"producttype"`      //必填	当当网	20	商品类型+..	1
	Transactionid    string `json:"transactionid"`    //可选	亚马逊	45	交易序号	222425696737
	Whsecode         string `json:"whsecode"`         //可选	融易购、达令网、美囤妈妈、唯品会	15	商品仓库编号	KU002
	Vipcooperationno int    `json:"vipcooperationno"` //可选	唯品会	15	合作编码	23
	Status           string `json:"status"`           //可选	淘宝	30	同步商品状态	up
	Platstoreid      string `json:"platstoreid"`      //可选	京东到家	15	门店ID	01
	Circuitbreakflag string `json:"circuitbreakflag"` //可选	唯品会JIT	20	熔断确认标记（1代表熔断确认）	1
	Storetype        string `json:"storetype"`        //可选	淘宝	15	仓库类型	1
	Scitemid         string `json:"scitemid"`         //必填	淘宝	32	后端商品ID	5165423156
	Privatefield     string `json:"privatefield"`     //必填	爱库存	15	私有字段(爱库存用于放置活动号)	23
	Operator         string `json:"operator"`         //必填	京东到家	64	操作人	张三
	Ischainstore     string `json:"ischainstore"`     //必填	有赞	4	是否为连锁总店	1
	Nodeshopid       string `json:"nodeshopid"`       //必填	有赞	32	连锁门店ID	1
	Detailquantity   string `json:"detailquantity"`   //可选	菠萝派自建商城	128	库存数量详情	1:498.00;2:1.00;3:0.00;4:30.00;5:0.00;7:0;8:0;9:0;0:0
	Whseid           string `json:"whseid"`           //可选	微盟智慧零售	128	仓库id	423213
	Requestid        int    `json:"requestid"`        //必填	亚马逊	4	查询批次号，仅限异步模式(如亚马逊)	180
	Countrycode      string `json:"countrycode"`      //必填	亚马逊、LaZaDa	15	国家编码(中国=JH_01，美国=JH_02，德国=JH...	UK
}
