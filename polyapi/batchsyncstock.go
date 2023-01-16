package polyapi

type Differ_JH_Business_BatchSyncStock struct {
	Goods        []*Differ_JH_Business_BatchSyncStock_GoodInfo //必填	通用	-	商品信息集合	-
	Requestid    int                                           //必填	亚马逊	4	查询批次号，仅限异步模式(如亚马逊)	180
	Countrycode  string                                        //必填	亚马逊、LaZaDa	15	国家编码(中国=JH_01，美国=JH_02，德国=JH...	UK
	Ischainstore string                                        //必填	有赞	4	是否为连锁总店	1
}
type Differ_JH_Business_BatchSyncStock_GoodInfo struct {
	Platproductid    string //必填	通用	32	平台商品ID	11447876487218
	Skuid            string //必填	通用	32	平台子规格ID	114478764
	Outerid          string //必填	通用	32	外部商家编码	WR6685851555
	Quantity         int    //必填	通用	4	库存数量	180
	Syncstocktype    string //必填	通用	25	库存更新方式(全量更新=JH_01，增量更新=JH...	JH_01
	Shoptype         string //可选	通用	25	店铺类型(普通=JH_001，分销=JH_002，经销=...	JH_001
	Outskuid         string //必填	通用	30	外部商家SKU编号	SK43008558525565
	Rroducttype      string //必填	当当网	20	商品类型+..	1
	Transactionid    string //可选	亚马逊	45	交易序号	222425696737
	Whsecode         string //可选	融易购、达令网、美囤妈妈、唯品会	15	商品仓库编号	KU002
	Vipcooperationno int    //可选	唯品会	15	合作编码	23
	Status           string //可选	淘宝	30	同步商品状态	up
	Platstoreid      string //可选	京东到家	15	门店ID	01
	Circuitbreakflag string //可选	唯品会JIT	20	熔断确认标记（1代表熔断确认）	1
	Storetype        string //可选	淘宝	15	仓库类型	1
	Scitemid         string //必填	淘宝	32	后端商品ID	5165423156
	Privatefield     string //必填	爱库存	15	私有字段(爱库存用于放置活动号)	23
	Operator         string //必填	京东到家	64	操作人	张三
	Ischainstore     string //必填	有赞	4	是否为连锁总店	1
	Nodeshopid       string //必填	有赞	32	连锁门店ID	1
	Detailquantity   string //可选	菠萝派自建商城	128	库存数量详情	1:498.00;2:1.00;3:0.00;4:30.00;5:0.00;7:0;8:0;9:0;0:0
	Whseid           string //可选	微盟智慧零售	128	仓库id	423213
}

type Differ_JH_Business_BatchSyncStock_RETURN struct {
	Code             string                                              //必填	通用	64	返回码	10000
	Msg              string                                              //必填	通用	64	返回消息	Success
	Subcode          string                                              //必填	通用	200	子集编码	LXO.JD.REQUEST_FAIL
	Submessage       string                                              //必填	通用	200	子级消息	订单已出库
	Polyapitotalms   int                                                 //必填	通用	64	菠萝派总耗时	102
	Polyapirequestid string                                              //必填	通用	64	请求菠萝派编号	20161222154212742
	Requestid        int                                                 //可选	亚马逊	4	查询批次号，仅限异步模式(如亚马逊)	180
	Results          []*Differ_JH_Business_BatchSyncStock_RETURN_results //	必填	通用	-	批量库存同步返回结果集合
}
type Differ_JH_Business_BatchSyncStock_RETURN_results struct {
	Issuccess        int    //必填	通用	1	是否成功(0表示失败；1表示成功)	0
	Code             string //必填	通用	256	错误编码	Failed
	Message          string //可选	通用	256	是否成功	Success
	Platproductid    string //可选	通用	32	平台商品ID（平台商品ID、平台子规格ID、外...	11447876487218
	Skuid            string //可选	通用	32	平台子规格ID（平台商品ID、平台子规格ID、...	114478764
	Outerid          string //可选	通用	32	外部商家编码（平台商品ID、平台子规格ID、...	WR6685851555
	Outskuid         string //可选	通用	30	外部商家SKU编号（平台商品ID、平台子规格I...	SK43008558525565
	Quantity         int    //必填	通用	4	更新后的数量	180
	Transactionid    string //可选	亚马逊	45	交易序号	222425696737
	Platstoreid      string //可选	京东到家、有赞	15	门店ID	01
	Whsecode         string //可选	融易购、达令网、美囤妈妈、唯品会	15	商品仓库编号	KU002
	Vipcooperationno int    //可选	唯品会	15	合作编码	23
}

func (client *Client) BatchSyncStock(request *Differ_JH_Business_BatchSyncStock, extData ...string) (*Differ_JH_Business_BatchSyncStock_RETURN, error) {
	method := "Differ.JH.Business.BatchSyncStock" //菠萝派批量同步接口
	bizcontent, err := json.Marshal(request)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	_, body, err := client.Execute(method, params)
	var OutData Differ_JH_Business_BatchSyncStock_RETURN
	if err != nil {
		if len(body) > 0 {
			json.Unmarshal(body, &OutData)
		}
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	if err != nil {
		logs.Debug("库存同步接口[" + err.Error() + "]")
	}
	return OutData, err
}
