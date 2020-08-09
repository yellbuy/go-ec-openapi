package qimen

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/thinkoner/openssl"

	"github.com/yellbuy/go-ec-openapi/cache"
	"github.com/yellbuy/go-ec-openapi/common"
)

var (
	// 正式地址,https
	router = "http://qimen.api.taobao.com/router/qimen/service"
	// 正式地址,http
	// router = "http://gw.api.taobao.com/router/rest"

	// 测试地址
	// router = "http://aliyuntest.polyapi.com/OpenAPI/do"
	// Timeout ...
	Timeout time.Duration
	// CacheExpiration 缓存过期时间
	CacheExpiration = time.Hour
	// GetCache 获取缓存
	GetCache cache.GetCacheFunc
	// SetCache 设置缓存
	SetCache cache.SetCacheFunc
)

type Client struct {
	Params *common.ClientParams
}

func InitClient(appKey, appSecret, session, customerId string) *Client {
	client := new(Client)
	client.Params = &common.ClientParams{appKey, appSecret, session, customerId}
	return client
}

// 获取签名
func getSign(appSecret string, params common.Parameter, body string) string {
	//12104E748A8F2DAD3EBA9F735A8D9C4F
	// 获取Key
	keys := []string{}
	for k := range params {
		keys = append(keys, k)
	}
	// 排序asc
	sort.Strings(keys)
	// 把所有参数名和参数值串在一起
	query := bytes.NewBufferString(appSecret)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(common.InterfaceToString(params[k]))
	}
	query.WriteString(body)
	query.WriteString(appSecret)
	//fmt.Println(query.String())
	//query = bytes.NewBufferString("b17cc059ffba3f6cf0d9131359d0be2aapp_key23268761formatxmlmethodtaobao.qimen.entryorder.createsign_methodmd5timestamp2020-04-08 10:56:00v2.0<request></request>b17cc059ffba3f6cf0d9131359d0be2a")
	// 使用MD5加密
	h := md5.New()
	io.Copy(h, query)
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func setRequestData(p common.Parameter, params *common.ClientParams, body []byte) common.Parameter {
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	p["timestamp"] = loc.Format("2006-01-02 15:04:05")
	p["app_key"] = params.AppKey
	p["format"] = "xml"
	p["v"] = "2.0"
	p["sign_method"] = "md5"
	p["customerId"] = params.PlatId
	// 设置签名
	p["sign"] = getSign(params.AppSecret, p, string(body))
	//fmt.Println(p["sign"])
	return p
}

// execute 执行API接口
func execute(client *Client, param common.Parameter, body []byte) (bytes []byte, err error) {
	err = checkConfig(client)
	if err != nil {
		return
	}

	var req *http.Request
	fullUrl := param.GetRequestData()
	fullUrl = fmt.Sprintf("%s?%s", router, fullUrl)
	//fmt.Println(fullUrl,string(body))
	beego.Debug(fullUrl, string(body))
	req, err = http.NewRequest("POST", fullUrl, strings.NewReader(string(body)))
	if err != nil {
		return
	}
	//fmt.Println(fullUrl, string(body))
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	httpClient := &http.Client{}
	httpClient.Timeout = Timeout
	var response *http.Response
	response, err = httpClient.Do(req)
	if err != nil {
		fmt.Println(err, fullUrl)
		return
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误，错误代码:%d，地址：%s", response.StatusCode, fullUrl)
		return
	}

	bytes, err = ioutil.ReadAll(response.Body)
	//fmt.Println("res:", string(bytes))
	return
}

// Execute 执行API接口
func (client *Client) Execute(method string, param common.Parameter, data []byte) (body []byte, err error) {
	param["method"] = method
	param = setRequestData(param, client.Params, data)
	body, err = execute(client, param, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = bytesToResult(body)
	return
}

func bytesToResult(bytes []byte) (err error) {
	res := new(Response)
	err = xml.Unmarshal(bytes, res)
	if err != nil {
		fmt.Println(string(bytes))
		return
	}
	if res.Flag != "success" {
		fmt.Println(string(bytes))
		err = fmt.Errorf("%s，错误代码：%s", res.Message, res.Code)
	}
	return err
}

// 检查配置
func checkConfig(client *Client) error {
	if client.Params.AppKey == "" {
		return errors.New("AppKey 不能为空")
	}
	if client.Params.AppSecret == "" {
		return errors.New("AppSecret 不能为空")
	}
	if router == "" {
		return errors.New("router 不能为空")
	}
	return nil
}

// 账号同步所需的加密方法
func aesEncrypt(appSecret string, origData []byte) ([]byte, error) {
	key := genPassword(appSecret)
	//fmt.Println("key:", key)
	keyBytes := []byte(key)
	crypted, _ := openssl.AesECBEncrypt(origData, keyBytes, openssl.PKCS7_PADDING)
	return crypted, nil
}

// 生成AES加密所需密码
func genPassword(appSecret string) string {
	//byteArr := []byte(appSecret)
	//fmt.Println(byteArr)
	query := bytes.NewBufferString(appSecret)
	// 使用MD5加密
	h := md5.New()
	io.Copy(h, query)
	content := h.Sum(nil)
	// 把二进制转化为大写的十六进制
	subContent := content[4:12]
	res := hex.EncodeToString(subContent)
	return res
}

// 字节数组转16进制字符串
func byteArrToHexString(DecimalSlice []byte) string {
	var sa = make([]string, 0)
	for _, v := range DecimalSlice {
		sa = append(sa, fmt.Sprintf("%02X", v))
	}
	ss := strings.Join(sa, "")
	return ss
}
