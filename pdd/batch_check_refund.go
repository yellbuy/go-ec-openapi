package pdd

import (
	"errors"

	"github.com/yellbuy/go-ec-openapi/common"
)

//批量检测退货
func (client *Client) BatchCheckRefundStatus(platorderno []string, extData ...string) (res []*common.BatchCheckRefundStatusRes, body []byte, err error) {
	err = errors.New("未实现")
	return
}
