package polyapi

import (
	"encoding/json"
	"fmt"

	"github.com/yellbuy/go-ec-openapi/common"
)

type WayBillConditionPost struct {
	Cpcode string `json:"cpcode"` //必填	通用	32	承运公司编码	POSTB
}
type WayBillConditionReturn struct {
	Code             string      //必填	通用	64	返回码	10000
	Msg              string      //必填	通用	64	返回消息	Success
	Subcode          string      //必填	通用	200	子集编码	LXO.JD.REQUEST_FAIL
	Submessage       string      //必填	通用	200	子级消息	订单已出库
	Polyapitotalms   json.Number //必填	通用	64	菠萝派总耗时	102
	Polyapirequestid string      //必填	通用	64	请求菠萝派编号	20161222154212742
	Results          []*WayBillConditionResults
}
type WayBillConditionResults struct {
	Cpcode  string                          `json:"cpcode"`  //必填	通用	32	承运公司编码	POSTB
	Cptype  string                          `json:"cptype"`  //必填	通用	32	物流服务商业务类型(直营=0，客户拥有的模...	POSTB
	Results []*WaybillApplySubscriptionInfo `json:"results"` //必填	通用	-	承运网点信息集合
}
type WaybillApplySubscriptionInfo struct {
	Usingquantity  int                                        `json:"usingquantity"`  //必填	通用	32	已用面单数量	40
	Sitecode       string                                     `json:"sitecode"`       //必填	通用	32	网点编码	1232
	Sitename       string                                     `json:"sitename"`       //必填	通用	32	网点名称	1232
	Sitestatus     string                                     `json:"sitestatus"`     //必填	通用	32	网点状态	1232
	Cancelquantity int                                        `json:"cancelquantity"` //必填	通用	32	取消的面单数量	40
	Printquantity  int                                        `json:"printquantity"`  //必填	通用	32	已经打印的面单总数	30
	Quantity       int                                        `json:"quantity"`       //必填	通用	32	面单数量(可用数量)	30
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

func (client *Client) GetWayBillCondition(postData *WayBillConditionPost) (*WayBillConditionReturn, error) {
	method := "Differ.JH.Logistics.GetWayBillCondition" //定义菠萝派退款检测批量接口
	bizcontent, err := json.Marshal(postData)
	req := make(map[string]interface{})
	req["bizcontent"] = string(bizcontent)
	params, err := common.InterfaceToParameter(req)
	_, body, err := client.Execute(method, params)
	fmt.Println(string(body))
	//logs.Debug(string(body))
	OutData := new(WayBillConditionReturn)
	if err != nil {
		return OutData, err
	}
	err = json.Unmarshal(body, &OutData)
	return OutData, err
}
