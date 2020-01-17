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
	GetAccessToken(code, redirectUri, state string, extData ...string) (res *common.AccessToken, body []byte, err error)
	GetWaybill(request *common.WaybillApplyNewRequest) (res *common.WaybillApplyNewCols, body []byte, err error)
}
