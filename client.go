package ecopenapi

import (
	"errors"

	"github.com/yellbuy/go-ec-openapi/common"
	"github.com/yellbuy/go-ec-openapi/pdd"
	"github.com/yellbuy/go-ec-openapi/polyapi"
	"github.com/yellbuy/go-ec-openapi/taobao"
)

const (
	TB = iota
	JD
	PDD
	POLYAPI
)

func NewClient(platformType int, params *common.ClientParams) (Client, error) {
	switch platformType {
	case TB:
		return &taobao.Client{params}, nil
	case PDD:
		return &pdd.Client{params}, nil
	case POLYAPI:
		return &polyapi.Client{params}, nil
	}
	return nil, errors.New("不支持的平台类型")
}

type Client interface {
	// 获取OAuth地址，菠萝派API需要通过extData传AppKey和AppSecret
	GetAuthorizeUrl(redirectUri, state string, extData ...string) (string, error)
	// 获取token
	GetAccessToken(code, redirectUri, state string, extData ...string) (res *common.AccessToken, body []byte, err error)
	// 电子面单
	GetWaybill(request *common.WaybillApplyNewRequest, extData ...string) (res *common.WaybillApplyNewCols, body []byte, err error)
	// 电子面单模板查询
	GetWaybillTemplates(request *common.WaybillTemplateRequest, extData ...string) (res *common.WaybillTemplateDto, body []byte, err error)
	// 商品下载
	DownloadProductList(pageIndex, pageSize int, status string, extData ...string) (res []*common.Product, hasNextPage bool, body []byte, err error)
	// 订单下载
	DownloadOrderList(pageIndex, pageSize int, startTime, endTime, timeType, orderStatus string, orderToken string, extData ...string) (res []*common.OrderInfo, hasNextPage bool, nextToken string, body []byte, err error)
	// 退货退款单下载
	DownloadRefundList(pageIndex, pageSize int, startTime, endTime, timeType, status string, orderToken string, extData ...string) (res []*common.OrderInfo, hasNextPage bool, nextToken string, body []byte, err error)
	// 订单发货
	LogisticsSend(dto *common.LogisticsSendReqDto, extData ...string) ([]byte, error)
}
