package kd100

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

var (
	routerAutoNumber = "https://www.kuaidi100.com/autonumber/autoComNum?resultv2=1&text=%s"
	routerQuery      = "https://www.kuaidi100.com/query?type=%s&postid=%s&temp=0.%s&phone="
	// router = "http://gw.api.tbsandbox.com/router/rest"
	// router = "https://baidu.com"
	// Timeout ...
	Timeout time.Duration
)

type Client struct {
	Params *common.ClientParams
}

func InitClient() *Client {
	client := new(Client)
	return client
}
func LogisticsQuery(postId string) (resDto *LogisticsQueryResDto, err error) {
	postId = strings.TrimSpace(postId)
	if postId == "" {
		return nil, errors.New("运单号不能为空")
	}
	err = nil
	resDto = new(LogisticsQueryResDto)
	client := InitClient()
	req := make(map[string]interface{})
	var bodyBytes []byte
	// 识别单号对应的快递公司类型
	bodyBytes, err = client.Execute(fmt.Sprintf(routerAutoNumber, postId), "POST", req)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := bytesToResult(bodyBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	arr, err := res.Get("auto").Array()
	if len(arr) == 0 {
		err = errors.New("无法识别的快递单号")
		return
	}
	for index := range arr {
		auto := res.Get("auto").GetIndex(index)
		var comCode string
		comCode, err = auto.Get("comCode").String()
		if err != nil {
			return resDto, err
		}
		bodyBytes, err = client.Execute(fmt.Sprintf(routerQuery, comCode, postId, time.Now().Unix()), "GET", req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = json.Unmarshal(bodyBytes, resDto)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if resDto.Message == "ok" {
			err = nil
			return
		}

	}
	err = errors.New("抱歉，暂无查询记录")
	return
	// {"comCode":"","num":"3223","auto":[]}
	// {"comCode":"","num":"773027491214552","auto":[{"comCode":"shentong","lengthPre":15,"noCount":24218044,"noPre":"773027"},{"comCode":"huitongkuaidi","lengthPre":15,"noCount":10,"noPre":"773027"}]}
	// {"message":"快递公司参数异常：单号不存在或者已经过期","nu":"","ischeck":"0","condition":"","com":"","status":"201","state":"0","data":[]}
	// {"message":"ok","nu":"773027491214552","ischeck":"1","condition":"F00","com":"shentong","status":"200","state":"3","data":[{"time":"2020-03-17 13:32:03","ftime":"2020-03-17 13:32:03","context":"已签收，签收人凭取货码签收。","location":null},{"time":"2020-03-16 16:10:59","ftime":"2020-03-16 16:10:59","context":"快件已暂存至岳阳上塔市镇汽车站店菜鸟驿站，如有疑问请联系15387404677","location":null},{"time":"2020-03-16 15:27:44","ftime":"2020-03-16 15:27:44","context":"湖南平江公司-平江申通(13907407722,0730-6266331)-派件中","location":null},{"time":"2020-03-16 14:42:44","ftime":"2020-03-16 14:42:44","context":"已到达-湖南平江公司","location":null},{"time":"2020-03-16 07:04:00","ftime":"2020-03-16 07:04:00","context":"湖南平江公司-已发往-湖南平江冬塔乡服务点","location":null},{"time":"2020-03-16 06:52:22","ftime":"2020-03-16 06:52:22","context":"已到达-湖南平江公司","location":null},{"time":"2020-03-15 18:22:35","ftime":"2020-03-15 18:22:35","context":"湖南长沙转运中心-已发往-湖南平江公司","location":null},{"time":"2020-03-15 18:10:46","ftime":"2020-03-15 18:10:46","context":"快件已在【湖南长沙转运中心】进行卸车，扫描员【进港五面扫4号】","location":null},{"time":"2020-03-15 18:10:46","ftime":"2020-03-15 18:10:46","context":"已到达-湖南长沙转运中心","location":null},{"time":"2020-03-15 04:47:34","ftime":"2020-03-15 04:47:34","context":"广东广州转运中心-已装袋发往-湖南长沙转运中心","location":null},{"time":"2020-03-15 04:47:34","ftime":"2020-03-15 04:47:34","context":"广东广州转运中心-已进行装车扫描","location":null},{"time":"2020-03-15 04:15:53","ftime":"2020-03-15 04:15:53","context":"广东广州转运中心-已装袋发往-湖南长沙转运中心","location":null},{"time":"2020-03-15 04:15:53","ftime":"2020-03-15 04:15:53","context":"广东广州转运中心-已进行装袋扫描","location":null},{"time":"2020-03-15 04:15:09","ftime":"2020-03-15 04:15:09","context":"广东广州转运中心-已装袋发往-湖南长沙转运中心","location":null},{"time":"2020-03-15 03:44:43","ftime":"2020-03-15 03:44:43","context":"已到达-广东广州转运中心","location":null},{"time":"2020-03-15 00:44:08","ftime":"2020-03-15 00:44:08","context":"广东广州火车站点-已发往-广东广州转运中心","location":null},{"time":"2020-03-15 00:40:42","ftime":"2020-03-15 00:40:42","context":"广东广州火车站点-直营揽收(18138701110,020-37183133)-已收件","location":null},{"time":"2020-03-14 22:45:07","ftime":"2020-03-14 22:45:07","context":"广东广州火车站点-瑶台18(16602092703,020-37183133)-已收件","location":null}]}

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

type LogisticsQueryResDto struct {
	Message       string `json:"message"`
	Nu            string `json:"nu"`
	Ischeck       string `json:"ischeck"`
	Condition     string `json:"condition"`
	Com           string `json:"com"`
	Status        string `json:"status"`
	State         string `json:"state"`
	LogisticsNo   string `json:"logisticsNo"`
	LogisticsName string `json:"logisticsName"`
	Data          []struct {
		Time     string      `json:"time"`
		Ftime    string      `json:"ftime"`
		Context  string      `json:"context"`
		Location interface{} `json:"location"`
	} `json:"data"`
}
