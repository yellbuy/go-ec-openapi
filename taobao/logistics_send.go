package taobao

import (
	"fmt"

	"github.com/yellbuy/go-ec-openapi/common"
)

// 订单发货
func (client *Client) LogisticsSend(dto *common.LogisticsSendReqDto, extData ...string) ([]byte, error) {
	return nil, fmt.Errorf("%s", "暂不支持的平台类型")
}
