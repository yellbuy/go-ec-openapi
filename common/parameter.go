package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/nilorg/sdk/convert"
	"github.com/yellbuy/go-ec-openapi/cache"
)

var (
	// Timeout ...
	Timeout time.Duration
	// CacheExpiration 缓存过期时间
	CacheExpiration = time.Hour
	// GetCache 获取缓存
	GetCache cache.GetCacheFunc
	// SetCache 设置缓存
	SetCache cache.SetCacheFunc
)

type ClientParams struct {
	AppKey string
	// AppSecret 秘密
	AppSecret string
	// Session 用户登录授权成功后，TOP颁发给应用的授权信息。当此API的标签上注明：“需要授权”，则此参数必传；“不需要授权”，则此参数不需要传；“可选授权”，则此参数为可选
	Session string
}

// Parameter 参数
type Parameter map[string]interface{}

func InterfaceToParameter(request interface{}) (Parameter, error) {
	if request == nil {
		return nil, errors.New("请求参数不能为空")
	}
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	reqJson, err := simplejson.NewJson(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	params, err := reqJson.Map()
	if err != nil {
		fmt.Println(err)
	}
	return params, err
}

// copyParameter 复制参数
func CopyParameter(srcParams Parameter) Parameter {
	newParams := make(Parameter)
	for key, value := range srcParams {
		newParams[key] = value
	}
	return newParams
}

func (p Parameter) SetRequestData(params *ClientParams) {
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	p["timestamp"] = strconv.FormatInt(loc.Unix(), 10)
	p["format"] = "json"
	p["app_key"] = params.AppKey
	p["v"] = "2.0"
	p["sign_method"] = "md5"
	p["partner_id"] = "Nilorg"
	if params.Session != "" {
		p["session"] = params.Session
	}
	// 设置签名
	p["sign"] = GetSign(params.AppSecret, p)
}

// 获取请求数据
func (p Parameter) GetRequestData() string {
	// 公共参数
	args := url.Values{}
	// 请求参数
	for key, val := range p {
		args.Set(key, InterfaceToString(val))
	}
	return args.Encode()
}

// 获取签名
func GetSign(appSecret string, params Parameter) string {
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
		query.WriteString(InterfaceToString(params[k]))
	}
	query.WriteString(appSecret)
	// 使用MD5加密
	h := md5.New()
	io.Copy(h, query)
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func InterfaceToString(src interface{}) string {
	if src == nil {
		panic(ErrTypeIsNil)
	}
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return convert.ToString(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// newCacheKey 创建缓存Key
func NewCacheKey(params Parameter) string {
	cpParams := Parameter(params)
	delete(cpParams, "session")
	delete(cpParams, "timestamp")
	delete(cpParams, "sign")

	keys := []string{}
	for k := range cpParams {
		keys = append(keys, k)
	}
	// 排序asc
	sort.Strings(keys)
	// 把所有参数名和参数值串在一起
	cacheKeyBuf := new(bytes.Buffer)
	for _, k := range keys {
		cacheKeyBuf.WriteString(k)
		cacheKeyBuf.WriteString("=")
		cacheKeyBuf.WriteString(InterfaceToString(cpParams[k]))
	}
	h := md5.New()
	io.Copy(h, cacheKeyBuf)
	return hex.EncodeToString(h.Sum(nil))
}
