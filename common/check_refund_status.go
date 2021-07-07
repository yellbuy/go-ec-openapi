package common

import (
	"fmt"
	"net/url"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/server/web"
	"yellbuy.com/YbCloudDataApi/models/common"
	"yellbuy.com/YbCloudDataApi/models/exp"
)

var PostUrl = make(map[string]string)

type BatchCheckRefundStatusReq struct {
	Orders []*BatchCheckRefundStatusOrder `json:"orders"`
}

type BatchCheckRefundStatusOrder struct {
	PlatOrderNo string `json:"platorderno"`
	ShopType    string `json:"shoptype"`
	CountryCode string `json:"countrycode"`
}

type BatchCheckRefundStatusRes struct {
	Issuccess                string                  `json:"issuccess"`
	Code                     string                  `json:"code"`
	Message                  string                  `json:"message"`
	PlatOrderNo              string                  `json:"platorderno"`
	RefundStatus             string                  `json:"refundstatus"`
	RefundStatusDescription  string                  `json:"refundstatusdescription"`
	TradeStatus              string                  `json:"tradestatus"`
	ChildrenRefundStatusList []*ChildrenRefundStatus `json:"childrenrefundstatus"`
}

type ChildrenRefundStatus struct {
	RefundNo                string `json:"refundno"`
	SubOrderNo              string `json:"suborderno"`
	ProductName             string `json:"productname"`
	PlatProductId           string `json:"platproductid"`
	TradeGoodsNo            string `json:"tradegoodsno"`
	RefundStatus            string `json:"refundstatus"`
	RefundStatusDescription string `json:"refundstatusdescription"`
}

func InitializationPostUrl() {
	var dsn string
	dbType := web.AppConfig.DefaultString("db.type", "postgres")
	if dbType == "" {
		dbType = "mysql"
	}
	dbhost, _ := web.AppConfig.String("db.host")
	dbport, _ := web.AppConfig.String("db.port")
	dbuser, _ := web.AppConfig.String("db.user")
	dbpassword, _ := web.AppConfig.String("db.password")
	dbname, _ := web.AppConfig.String("db.name")
	timezone, _ := web.AppConfig.String("db.timezone")

	if dbport == "" {
		if dbType == "mysql" {
			dbport = "3306"
		} else if dbType == "postgres" {
			dbport = "5432"
		}
	}
	if dbType == "mysql" {
		dsn = dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
		// fmt.Println(dsn)

		if timezone != "" {
			dsn = dsn + "&loc=" + url.QueryEscape(timezone)
		}
	} else if dbType == "postgres" {
		dsn = "user=" + dbuser + " password=" + dbpassword + " dbname=" + dbname + " host=" + dbhost + " port=" + dbport + " sslmode=disable"
	}

	maxidleconns := web.AppConfig.DefaultInt("db.maxidleconns", 0)
	maxopenconns := web.AppConfig.DefaultInt("db.maxopenconns", 0)
	stmtcachesize := web.AppConfig.DefaultInt("db.stmtcachesize", 100)

	runmode := web.AppConfig.DefaultString("runmode", "prod")
	if runmode == "dev" {
		fmt.Println("dsn:", dsn)
		orm.Debug = true
	}
	var aliasName string
	if aliasName == "" {
		aliasName = "default"
	}
	orm.RegisterDataBase(aliasName, dbType, dsn, maxidleconns, maxopenconns, stmtcachesize)

	//约单号需要提交的URL
	expList := make([]*exp.Company, 0)
	sql := "select * from exp_company where is_del=0 and appid=9"
	_, err := orm.NewOrm().Raw(sql).QueryRows(&expList)
	if err != nil {
		logs.Debug("初始化菠萝派提交地址错误【请联系管理员[约单号接口]】")
	}
	for _, val := range expList {
		if len(val.Remark) > 0 {
			PostUrl[val.PlatLogisticsId] = val.Remark
		}
	}
	//其他接口需要的URL
	list := make([]*common.CommonData, 0)
	sql = "select * from common_data where type='polyapi_shop_kind' and is_del=0"
	_, err = orm.NewOrm().Raw(sql).QueryRows(&list)
	if err != nil {
		logs.Debug("初始化菠萝派提交地址错误【请联系管理员[其他接口]】")
	}
	for _, val := range list {
		if len(val.Desc) > 0 {
			PostUrl[val.Code] = val.Desc
		}
	}
	logs.Debug("菠萝派提交URL初始化完成")
}
