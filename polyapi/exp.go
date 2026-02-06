package polyapi

import (
	"encoding/json"

	"github.com/spf13/cast"
	"github.com/yellbuy/go-ec-openapi/common"
)

type WayBillConditionPost struct {
	Cpcode     string `json:"cpcode"` //必填	通用	32	承运公司编码	POSTB
	BizVersion string
}
type JDWayBillConditionPost struct {
	ShopType     string `json:"shopType"`
	ShipperState string `json:"shipperState"`
}
type WayBillConditionReturn struct {
	Code             string                     //必填	通用	64	返回码	10000
	Msg              string                     //必填	通用	64	返回消息	Success
	Subcode          string                     //必填	通用	200	子集编码	LXO.JD.REQUEST_FAIL
	Submessage       string                     //必填	通用	200	子级消息	订单已出库
	Polyapitotalms   json.Number                //必填	通用	64	菠萝派总耗时	102
	Polyapirequestid string                     //必填	通用	64	请求菠萝派编号	20161222154212742
	Results          []*WayBillConditionResults `json:"results"`
}
type JDWayBillConditionReturn struct {
	Code             string                        //必填	通用	64	返回码	10000
	Msg              string                        //必填	通用	64	返回消息	Success
	Subcode          string                        //必填	通用	200	子集编码	LXO.JD.REQUEST_FAIL
	Submessage       string                        //必填	通用	200	子级消息	订单已出库
	Polyapitotalms   json.Number                   //必填	通用	64	菠萝派总耗时	102
	Polyapirequestid string                        //必填	通用	64	请求菠萝派编号	20161222154212742
	Results          []*JDWayBillCondition_Results `json:"results"`
}

func (c *JDWayBillConditionReturn) ToWayBillConditionReturn() *WayBillConditionReturn {
	outData := new(WayBillConditionReturn)
	outData.Code = c.Code
	outData.Msg = c.Msg
	outData.Subcode = c.Subcode
	outData.Submessage = c.Submessage
	outData.Polyapitotalms = c.Polyapitotalms
	outData.Polyapirequestid = c.Polyapirequestid
	outData.Results = make([]*WayBillConditionResults, 0)
	for _, val := range c.Results {
		results := new(WayBillConditionResults)
		results.Cpcode = val.ProviderCode
		results.Cptype = val.OperationType
		results.Results = make([]*WaybillApplySubscriptionInfo, 0)
		resultsNode := new(WaybillApplySubscriptionInfo)
		resultsNode.Sitename = val.SiteName
		resultsNode.Sitecode = val.SiteCode
		resultsNode.Shippers = make([]*WaybillApplySubscriptionInfoShipperInfo, 0)
		resultsNode.Shippers = append(resultsNode.Shippers, val.Address)
		resultsNode.Quantity = cast.ToInt(val.Amount)
		results.Results = append(results.Results, resultsNode)
		outData.Results = append(outData.Results, results)
	}
	return outData
}

type JDWayBillCondition_Results struct {
	IsSuccess               string                                   `json:"isSuccess"`               // 是否成功(0表示失败；1表示成功)
	Code                    string                                   `json:"code"`                    // 错误编码 成功-10000
	Message                 string                                   `json:"message"`                 // 失败说明
	SubCode                 string                                   `json:"subCode"`                 // 返回子码
	SubMessage              string                                   `json:"subMessage"`              // 返回子级消息
	MessageFrom             string                                   `json:"messageFrom"`             // 指示此次返回是菠萝派的返回还是平台的返回，默认 POLY:POLY=POLY,PLAT=PLAT,TWICE=TWICE
	ProviderId              string                                   `json:"providerId"`              // 承运商ID
	ProviderCode            string                                   `json:"providerCode"`            // 承运商编码
	ProviderName            string                                   `json:"providerName"`            // 承运商名称
	ProvideType             string                                   `json:"provideType"`             // 承运商类型:快递公司=JH_001,物流公司=JH_002,安装公司=JH_003,生鲜冷链承运商=JH_004,其他=JH_99
	ProviderState           string                                   `json:"providerState"`           // 承运商状态
	OperationType           string                                   `json:"operationType"`           // 经营类型
	Name                    string                                   `json:"name"`                    // 承运商联系人
	Phone                   string                                   `json:"phone"`                   // 承运商联系电话
	Mobile                  string                                   `json:"mobile"`                  // 承运商联系手机
	SupportCOD              string                                   `json:"supportCOD"`              // 是否支持货到付款
	SiteCode                string                                   `json:"siteCode"`                // 网点编码
	SiteName                string                                   `json:"siteName"`                // 网点名称
	ShipperNo               string                                   `json:"shipperNo"`               // 月结卡号
	Amount                  string                                   `json:"amount"`                  // 剩余单号量
	Address                 *WaybillApplySubscriptionInfoShipperInfo `json:"address"`                 // 地址信息
	LogisticsServiceID      string                                   `json:"logisticsserviceid"`      // 物流服务ID
	LogisticsServiceName    string                                   `json:"logisticsservicename"`    // 物流服务名称
	DeliveryAddress         string                                   `json:"deliveryaddress"`         // 送货地址
	WarehouseName           string                                   `json:"warehousename"`           // 仓库名称
	LogisticsTimeliness     string                                   `json:"logisticstimeliness"`     // 物流时效
	ExpressLogisticsService string                                   `json:"expresslogisticsservice"` // 快递物流服务
	TrialResult             string                                   `json:"trialresult"`             // 试用结果
	OrderNo                 string                                   `json:"orderno"`                 // 订单编号
}

type WayBillConditionResults struct {
	Cpcode           string                          `json:"cpcode"`  //必填	通用	32	承运公司编码	POSTB
	Cptype           string                          `json:"cptype"`  //必填	通用	32	物流服务商业务类型(直营=0，客户拥有的模...	POSTB
	Results          []*WaybillApplySubscriptionInfo `json:"results"` //必填	通用	-	承运网点信息集合
	ShopId           string                          `json:"shopid"`
	Logisticsaccount string                          `json:"logisticsaccount"`
}
type WaybillApplySubscriptionInfo struct {
	Usingquantity  interface{}                                `json:"usingquantity"`  //必填	通用	32	已用面单数量	40
	Sitecode       string                                     `json:"sitecode"`       //必填	通用	32	网点编码	1232
	Sitename       string                                     `json:"sitename"`       //必填	通用	32	网点名称	1232
	Sitestatus     string                                     `json:"sitestatus"`     //必填	通用	32	网点状态	1232
	Cancelquantity interface{}                                `json:"cancelquantity"` //必填	通用	32	取消的面单数量	40
	Printquantity  interface{}                                `json:"printquantity"`  //必填	通用	32	已经打印的面单总数	30
	Quantity       interface{}                                `json:"quantity"`       //必填	通用	32	面单数量(可用数量)	30
	Shippers       []*WaybillApplySubscriptionInfoShipperInfo `json:"shippers"`       //必填	通用	-	网点下发货信息集合	-
	Services       []*WaybillApplySubscriptionInfoServiceInfo `json:"services"`       //必填	通用	-	服务信息列表	-
}
type WaybillApplySubscriptionInfoShipperInfo struct {
	Province string `json:"province"` //必填	通用	32	州省	浙江
	City     string `json:"city"`     //必填	通用	32	城市	杭州
	Area     string `json:"area"`     //必填	通用	50	区县	江干区
	Town     string `json:"town"`     //必填	通用	50	镇（街道）	三里亭
	Address  string `json:"address"`  //必填	通用	200	地址	浙江杭州市江干区秋涛路255号10楼302
}
type WaybillApplySubscriptionInfoServiceInfo struct {
	Name        string                                                `json:"name"`        //必填	通用	32	服务名称	代收货款
	Code        string                                                `json:"code"`        //必填	通用	32	服务编码	SVC-COD
	Desc        string                                                `json:"desc"`        //必填	通用	32	服务描述	SVC-COD
	Isrequired  string                                                `json:"isrequired"`  //必填	通用	32	是否必选服务(不必须=0，必须=1)	1
	Serviceattr []*WaybillApplySubscriptionInfoServiceInfoServiceAttr `json:"serviceattr"` //可选	通用	-	服务属性	-
}
type WaybillApplySubscriptionInfoServiceInfoServiceAttr struct {
	Code     string `json:"code"`     //必填	通用	32	属性的值	1
	Name     string `json:"name"`     //必填	通用	32	属性的名称	1
	Type     string `json:"type"`     //必填	通用	32	属性的类别	1
	Typedesc string `json:"typedesc"` //必填	通用	32	枚举类型的枚举值	1
}

func (client *Client) GetJDWayBillCondition(postData *JDWayBillConditionPost) (*JDWayBillConditionReturn, error) {
	method := "Differ.JH.Logistics.GetProvider" //定义菠萝派退款检测批量接口
	bizcontent, err := json.Marshal(postData)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	// fmt.Println(string(bizcontent))
	params, err := common.InterfaceToParameter(req)
	// body1, _ := json.Marshal(params)
	// fmt.Println(string(body1))
	_, body, err := client.Execute(method, params)
	// fmt.Println(string(body))
	OutData := new(JDWayBillConditionReturn)
	if err != nil {
		return OutData, err
	}
	err = json.Unmarshal(body, OutData)
	return OutData, err
}
func (client *Client) GetWayBillCondition(postData *WayBillConditionPost) (*WayBillConditionReturn, error) {
	method := "Differ.JH.Logistics.GetWayBillCondition" //定义菠萝派退款检测批量接口
	bizcontent, err := json.Marshal(postData)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	// fmt.Println(string(bizcontent))
	params, err := common.InterfaceToParameter(req)
	// body1, _ := json.Marshal(params)
	// fmt.Println(string(body1))
	_, body, err := client.Execute(method, params)
	// fmt.Println(string(body))
	OutData := new(WayBillConditionReturn)
	if err != nil {
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	return OutData, err
}
func (client *Client) GetTemplates(postData *common.GetTemplates) (*common.TemplatesReturn, error) {
	method := "Differ.JH.Logistics.GetTemplates" //定义菠萝派退款检测批量接口
	bizcontent, _ := json.Marshal(postData)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	// fmt.Println(string(bizcontent))
	params, _ := common.InterfaceToParameter(req)
	// body1, _ := json.Marshal(params)
	// fmt.Println(string(body1))
	_, body, err := client.Execute(method, params)
	// fmt.Println(string(body))
	OutData := new(common.TemplatesReturn)
	if err != nil {
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	return OutData, err
}
