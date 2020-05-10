package common

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
