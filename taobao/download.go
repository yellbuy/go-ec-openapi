package taobao

import (
	"errors"

	simplejson "github.com/bitly/go-simplejson"
)

func (client *Client) DownloadProductList(pageIndex, pageSize int, extData ...string) (res *simplejson.Json, body []byte, err error) {
	return nil, nil, errors.New("未实现")
}

func (client *Client) DownloadOrderList(pageIndex, pageSize int, startTime, endTime, timeType, orderStatus string, extData ...string) (res *simplejson.Json, body []byte, err error) {
	return nil, nil, errors.New("未实现")
}
