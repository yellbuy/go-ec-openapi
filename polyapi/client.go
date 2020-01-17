package polyapi

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/cache"
	"github.com/yellbuy/go-ec-openapi/common"
)

var (
	router = "http://39.98.7.126/OpenAPI/do"
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
	p["token"] = params.Session
	p["platid"] = params.PlatId
	p["version"] = "1.5"
	p["contenttype"] = "json"
	// 设置签名
	p["sign"] = common.GetSign(params.AppSecret, p, true)
	return p
}

// execute 执行API接口
func execute(client *Client, param common.Parameter) (bytes []byte, err error) {
	err = checkConfig(client)
	if err != nil {
		return
	}

	var req *http.Request
	data := param.GetRequestData()
	// fmt.Println("data:", data)
	req, err = http.NewRequest("POST", router, strings.NewReader(data))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	httpClient := &http.Client{}
	httpClient.Timeout = Timeout
	var response *http.Response
	response, err = httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}
	defer response.Body.Close()
	bytes, err = ioutil.ReadAll(response.Body)
	return
}

// Execute 执行API接口
func (client *Client) Execute(method string, param common.Parameter) (res *simplejson.Json, body []byte, err error) {
	param["method"] = method
	param = setRequestData(param, client.Params)
	body, err = execute(client, param)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err = bytesToResult(body)
	return
}

func bytesToResult(bytes []byte) (res *simplejson.Json, err error) {
	res, err = simplejson.NewJson(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	code := res.Get("code").MustString()
	if code != "10000" {
		fmt.Println("code:", code)
		fmt.Println("subcode:", res.Get("subcode").MustString())
		msg := res.Get("submessage").MustString()
		if msg == "" {
			msg = res.Get("msg").MustString()
		}
		fmt.Println("msg:", msg)
		err = errors.New(msg)
	}
	return
}

// ExecuteCache 执行API接口，缓存
func (client *Client) ExecuteCache(method string, param common.Parameter) (res *simplejson.Json, err error) {
	param["method"] = method
	param = setRequestData(param, client.Params)

	cacheKey := common.NewCacheKey(param)
	// 获取缓存
	if GetCache != nil {
		cacheBytes := GetCache(cacheKey)
		if len(cacheBytes) > 0 {
			res, err = simplejson.NewJson(cacheBytes)
			if err == nil && res != nil {
				return
			}
		}
	}

	var bodyBytes []byte
	bodyBytes, err = execute(client, param)
	if err != nil {
		return
	}
	res, err = bytesToResult(bodyBytes)
	if err != nil {
		return
	}
	ejsonBody, _ := res.MarshalJSON()
	// 设置缓存
	if SetCache != nil {
		go SetCache(cacheKey, ejsonBody, CacheExpiration)
	}
	return
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
	// { 0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF, 0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF }
	iv := []byte{18, 52, 86, 120, 144, 171, 205, 239, 18, 52, 86, 120, 144, 171, 205, 239}
	key := genPassword(appSecret)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	//fmt.Println("blockSize:", blockSize)
	origData = paddingPKCS7(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func paddingPKCS7(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func paddingPKCS5(ciphertext []byte) []byte {
	return paddingPKCS7(ciphertext, 8)
}

// 生成AES加密所需密码
func genPassword(appSecret string) string {
	query := bytes.NewBufferString(appSecret)
	// 使用MD5加密
	h := md5.New()
	io.Copy(h, query)
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
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
