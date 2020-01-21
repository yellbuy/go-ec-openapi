package pdd

import (
	"errors"

	"github.com/yellbuy/go-ec-openapi/common"
)

func (client *Client) DownloadProductList(pageIndex, pageSize int, status string, extData ...string) (res []*common.Product, hasNextPage bool, body []byte, err error) {
	err = errors.New("未实现")
	return
}

func (client *Client) DownloadOrderList(pageIndex, pageSize int, startTime, endTime, timeType, orderStatus string, orderToken string, extData ...string) (res []*common.OrderInfo, hasNextPage bool, nextToken string, body []byte, err error) {
	err = errors.New("未实现")
	return
}
