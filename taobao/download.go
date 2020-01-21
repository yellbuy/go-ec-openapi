package taobao

import (
	"errors"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

func (client *Client) DownloadProductList(pageIndex, pageSize int, status string, extData ...string) (res []*common.Product, hasNextPage bool, body []byte, err error) {
	err = errors.New("未实现")
	return
}

func (client *Client) DownloadOrderList(pageIndex, pageSize int, startTime, endTime, timeType, orderStatus string, extData ...string) (res *simplejson.Json, body []byte, err error) {
	err = errors.New("未实现")
	return
}
