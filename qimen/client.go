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
	"strings"
	"time"

	"github.com/thinkoner/openssl"

	"github.com/yellbuy/go-ec-openapi/cache"
	"github.com/yellbuy/go-ec-openapi/common"
)

var (
	// 正式地址,https
	router = "https://eco.taobao.com/router/rest"
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

func InitClient(appKey, appSecret, session string) *Client {
	client := new(Client)
	client.Params = &common.ClientParams{appKey, appSecret, session, ""}
	return client
}
func setRequestData(p common.Parameter, params *common.ClientParams) common.Parameter {
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	p["timestamp"] = loc.Format("2006-01-02 15:04:05")
	p["appkey"] = params.AppKey
	p["session"] = params.Session
	p["format"] = "xml"
	p["v"] = "2.0"
	p["sign_method"] = "md5"
	// 设置签名
	p["sign"] = common.GetSign(params.AppSecret, p, true)
	return p
}

// execute 执行API接口
func execute(client *Client, param common.Parameter, body []byte) (bytes []byte, err error) {
	err = checkConfig(client)
	if err != nil {
		return
	}

	var req *http.Request
	req, err = http.NewRequest("POST", fmt.Sprintf("%s?app_key=%s&method=%s&sign=%s&timestamp=%s", router, param["appkey"], param["method"], param["sign"], param["timestamp"]), strings.NewReader(string(body)))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	httpClient := &http.Client{}
	httpClient.Timeout = Timeout
	var response *http.Response
	response, err = httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("data:", body)
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}
	defer response.Body.Close()
	bytes, err = ioutil.ReadAll(response.Body)
	//fmt.Println(string(bytes))
	return
}

// Execute 执行API接口
func (client *Client) Execute(method string, param common.Parameter, data []byte) (body []byte, err error) {
	param["method"] = method
	param = setRequestData(param, client.Params)
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
	if client.Params.Session == "" {
		return errors.New("Session 不能为空")
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
	byteArr := []byte(appSecret)
	fmt.Println(byteArr)
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
