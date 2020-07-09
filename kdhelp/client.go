package kdhelp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

var (
	Timeout time.Duration
)

type Client struct {
	Params *common.ClientParams
}

func InitClient() *Client {
	client := new(Client)
	return client
}

// execute 执行API接口
func execute(client *Client, url, method string, param common.Parameter) (bytes []byte, err error) {

	var req *http.Request
	data := param.GetRequestData()
	req, err = http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	httpClient := &http.Client{}
	httpClient.Timeout = Timeout
	var response *http.Response
	response, err = httpClient.Do(req)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}

	bytes, err = ioutil.ReadAll(response.Body)
	return
}

// Execute 执行API接口
func (client *Client) Execute(url, method string, param common.Parameter) (res []byte, err error) {

	res, err = execute(client, url, method, param)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func bytesToResult(bytes []byte) (res *simplejson.Json, err error) {
	res, err = simplejson.NewJson(bytes)
	return
}
