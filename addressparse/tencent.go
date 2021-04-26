package addressparse

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	gourl "net/url"
	"strings"
	"time"
)

type AddressInfo struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Result  Address `json:"result"`
}

type Address struct {
	// 街道
	Street string `json:"street"`
	// 完整名称
	FullName string `json:"full_name"`
	// 电话
	Phone string `json:"telephone"`
	// 手机
	Mobile string `json:"mobile"`
	// 省
	Province string `json:"province"`
	// 市
	City string `json:"city"`
	// 区
	County string `json:"area"`
}

func calcAuthorization(source string, secretId string, secretKey string) (auth string, datetime string, err error) {
	//timeLocation, _ := time.LoadLocation("Local")
	datetime = time.Now().In(time.UTC).Format("Mon, 02 Jan 2006 15:04:05 GMT")
	signStr := fmt.Sprintf("x-date: %s\nx-source: %s", datetime, source)

	// hmac-sha1
	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(signStr))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	auth = fmt.Sprintf("hmac id=\"%s\", algorithm=\"hmac-sha1\", headers=\"x-date x-source\", signature=\"%s\"",
		secretId, sign)

	return auth, datetime, nil
}

func urlencode(params map[string]string) string {
	var p = gourl.Values{}
	for k, v := range params {
		p.Add(k, v)
	}
	return p.Encode()
}

func TenantAddressParse(secretId, secretKey, content string) (res *AddressInfo, err error) {
	res = new(AddressInfo)
	// // 云市场分配的密钥Id
	// secretId := "xxxx"
	// // 云市场分配的密钥Key
	// secretKey := "xxxx"
	source := "market"

	// 签名
	auth, datetime, _ := calcAuthorization(source, secretId, secretKey)

	// 请求方法
	method := "GET"
	// 请求头
	headers := map[string]string{"X-Source": source, "X-Date": datetime, "Authorization": auth}

	// 查询参数
	queryParams := make(map[string]string)
	queryParams["content"] = content
	// body参数
	bodyParams := make(map[string]string)

	// url参数拼接
	url := "https://service-6qbnpmw2-1300683954.gz.apigw.tencentcs.com/release/dizhi"
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s?%s", url, urlencode(queryParams))
	}

	bodyMethods := map[string]bool{"POST": true, "PUT": true, "PATCH": true}
	var body io.Reader = nil
	if bodyMethods[method] {
		body = strings.NewReader(urlencode(bodyParams))
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, e := http.NewRequest(method, url, body)
	if e != nil {
		fmt.Println(err)
		err = e
		return
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	response, e := client.Do(request)
	if e != nil {
		fmt.Println(err)
		err = e
		return
	}
	defer response.Body.Close()

	bodyBytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		fmt.Println(err)
		err = e
		return
	}
	e = json.Unmarshal(bodyBytes, res)
	if e != nil {
		fmt.Println(err)
		err = e
	}
	return
}
