package pdd

import (
	"errors"

	"github.com/yellbuy/go-ec-openapi/common"
)

func (client *Client) DownloadProductList(pageIndex, pageSize int, status, productToken string, extData ...string) (res []*common.Product, hasNextPage bool, nextToken string, body []byte, err error) {
	err = errors.New("未实现")
	return
}

func (client *Client) DownloadOrderList(pageIndex, pageSize int, startTime, endTime, timeType, orderStatus string, orderToken string, extData ...string) (res []*common.OrderInfo, hasNextPage bool, nextToken string, body []byte, err error) {
	err = errors.New("未实现")
	return
}
func (client *Client) DownloadOrderListByQimen(pageIndex, pageSize int, startTime, endTime, timeType, orderStatus string, extData ...string) (hasNextPage bool, body []byte, err error) {
	err = errors.New("未实现")
	return
}

// 退货退款单下载
func (client *Client) DownloadRefundList(pageIndex, pageSize int, startTime, endTime, timeType, status string, orderToken string, extData ...string) (res []*common.OrderInfo, hasNextPage bool, nextToken string, body []byte, err error) {
	err = errors.New("未实现")
	return
}
