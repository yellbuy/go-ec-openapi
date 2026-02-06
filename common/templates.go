package common

import "encoding/json"

type GetTemplates struct {
	Templatestype string `json:"templatestype"` //必填	通用	32	订单信息(所有模版=ALL，客户拥有的模版=OW...	ALL
	Logistictype  string `json:"logistictype"`  //必填	拼多多金虹桥	-	快递公司类别	YTO
	BizVersion    string
}

type TemplatesReturn struct {
	Code             string                    `json:"code"`             //必填	通用	64	返回码	10000
	Msg              string                    `json:"msg"`              //必填	通用	64	返回消息	Success
	Subcode          string                    `json:"subcode"`          //必填	通用	200	子集编码	LXO.JD.REQUEST_FAIL
	Submessage       string                    `json:"submessage"`       //必填	通用	200	子级消息	订单已出库
	Polyapitotalms   json.Number               `json:"polyapitotalms"`   //必填	通用	64	菠萝派总耗时	102
	Polyapirequestid string                    `json:"polyapirequestid"` //必填	通用	64	请求菠萝派编号	20161222154212742
	Results          []*TemplatesReturnResults `json:"results"`
}
type TemplatesReturnResults struct {
	Cpcode  string          `json:"cpcode"`  //必填	通用	32	承运公司编码	POSTB
	Results []*TemplateInfo `json:"results"` //必填	通用	-	CP模版信息集合	-
}
type TemplateInfo struct {
	Id                      string `json:"id"`           //必填	通用	32	模板id	1024
	Name                    string `json:"name"`         //必填	通用	32	模板名称	YTO模板
	Url                     string `json:"url"`          //必填	通用	32	模板url	http://www.wdgj.com
	Templatetype            string `json:"templatetype"` //必填	通用	32	模板类别
	Brandcode               string `json:"brandcode"`    //必填	通用	32	品牌编码
	Templatetypedescription string `json:"templatetypedescription"`
}
