package ecopenapi

import (
	"errors"

	"github.com/yellbuy/go-ec-openapi/common"
	"github.com/yellbuy/go-ec-openapi/taobao"
)

const (
	TAOBAO = iota
	JD
	PDD
)

func NewClient(platformType int, params *common.ClientParams) (Client, error) {
	switch platformType {
	case TAOBAO:
		return &taobao.Client{params}, nil
	}
	return nil, errors.New("不支持的平台类型")
}

type Client interface {
	GetWaybill(request *common.WaybillApplyNewRequest) (res *common.WaybillApplyNewCols, err error)
}
