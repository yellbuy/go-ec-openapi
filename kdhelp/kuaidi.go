package kdhelp

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// http://www.kuaidi.com/

var (
	kuaidiSelectInfoRouter        = "http://www.kuaidi.com/index-ajaxselectinfo-%s.html"
	kuaidiSelectCourierInfoRouter = "http://www.kuaidi.com/index-ajaxselectcourierinfo-%s-%s.html"
)

func KuaidiLogisticsQuery(logisticsCode string) (resDto *KuaidiLogisticsQueryResDto, err error) {
	logisticsCode = strings.TrimSpace(logisticsCode)
	if logisticsCode == "" {
		return nil, errors.New("运单号不能为空")
	}
	err = nil
	resDto = new(KuaidiLogisticsQueryResDto)

	client := InitClient()
	req := make(map[string]interface{})

	var bodyBytes []byte
	// 识别单号对应的快递公司类型
	// bodyBytes, err = client.Execute(fmt.Sprintf(routerAutoNumber, postId), "POST", req)
	bodyBytes, err = client.Execute(fmt.Sprintf(kuaidiSelectInfoRouter, logisticsCode), "GET", req)
	if err != nil {
		fmt.Println(err)
		return
	}
	comList := make(map[string]KuaidiCompanyQueryResDto)
	err = json.Unmarshal(bodyBytes, &comList)
	if err != nil {
		fmt.Println(err, string(bodyBytes))
		err = errors.New("抱歉，暂无查询记录")
	}
	if len(comList) == 0 {
		err = errors.New("单号不正确")
	}
	var index int
	for index = 0; index < len(comList); index++ {
		if val, ok := comList[fmt.Sprintf("%d", index)]; ok {
			url := fmt.Sprintf(kuaidiSelectCourierInfoRouter, logisticsCode, val.ExName)
			//url = "http://www.kuaidi.com/index-ajaxselectcourierinfo-3103616296777-yunda.html"
			bodyBytes, err = client.Execute(url, "POST", req)
			if err != nil {
				fmt.Println(err)
			} else {
				err = json.Unmarshal(bodyBytes, resDto)
				if err != nil {
					fmt.Println(err, string(bodyBytes))
				} else if resDto.Success {
					if len(resDto.Data) > 1 {
						fmt.Println("KuaidiLogisticsQuery:", resDto)
						return
					} else if len(resDto.Data) == 1 && resDto.Data[0].Context != "查无结果" {
						fmt.Println("KuaidiLogisticsQuery:", resDto)
						return
					} else {
						fmt.Println("KuaidiLogisticsQuery url:", url, resDto)
					}
				} else {
					fmt.Println("KuaidiLogisticsQuery url:", url, resDto)
				}
			}
		}
	}
	fmt.Println("KuaidiLogisticsQuery:", logisticsCode, comList)
	err = errors.New("抱歉，暂无查询记录")
	return
	// {"comCode":"","num":"3223","auto":[]}
	// {"comCode":"","num":"773027491214552","auto":[{"comCode":"shentong","lengthPre":15,"noCount":24218044,"noPre":"773027"},{"comCode":"huitongkuaidi","lengthPre":15,"noCount":10,"noPre":"773027"}]}
	// {"message":"快递公司参数异常：单号不存在或者已经过期","nu":"","ischeck":"0","condition":"","com":"","status":"201","state":"0","data":[]}
	// {"message":"ok","nu":"773027491214552","ischeck":"1","condition":"F00","com":"shentong","status":"200","state":"3","data":[{"time":"2020-03-17 13:32:03","ftime":"2020-03-17 13:32:03","context":"已签收，签收人凭取货码签收。","location":null},{"time":"2020-03-16 16:10:59","ftime":"2020-03-16 16:10:59","context":"快件已暂存至岳阳上塔市镇汽车站店菜鸟驿站，如有疑问请联系15387404677","location":null},{"time":"2020-03-16 15:27:44","ftime":"2020-03-16 15:27:44","context":"湖南平江公司-平江申通(13907407722,0730-6266331)-派件中","location":null},{"time":"2020-03-16 14:42:44","ftime":"2020-03-16 14:42:44","context":"已到达-湖南平江公司","location":null},{"time":"2020-03-16 07:04:00","ftime":"2020-03-16 07:04:00","context":"湖南平江公司-已发往-湖南平江冬塔乡服务点","location":null},{"time":"2020-03-16 06:52:22","ftime":"2020-03-16 06:52:22","context":"已到达-湖南平江公司","location":null},{"time":"2020-03-15 18:22:35","ftime":"2020-03-15 18:22:35","context":"湖南长沙转运中心-已发往-湖南平江公司","location":null},{"time":"2020-03-15 18:10:46","ftime":"2020-03-15 18:10:46","context":"快件已在【湖南长沙转运中心】进行卸车，扫描员【进港五面扫4号】","location":null},{"time":"2020-03-15 18:10:46","ftime":"2020-03-15 18:10:46","context":"已到达-湖南长沙转运中心","location":null},{"time":"2020-03-15 04:47:34","ftime":"2020-03-15 04:47:34","context":"广东广州转运中心-已装袋发往-湖南长沙转运中心","location":null},{"time":"2020-03-15 04:47:34","ftime":"2020-03-15 04:47:34","context":"广东广州转运中心-已进行装车扫描","location":null},{"time":"2020-03-15 04:15:53","ftime":"2020-03-15 04:15:53","context":"广东广州转运中心-已装袋发往-湖南长沙转运中心","location":null},{"time":"2020-03-15 04:15:53","ftime":"2020-03-15 04:15:53","context":"广东广州转运中心-已进行装袋扫描","location":null},{"time":"2020-03-15 04:15:09","ftime":"2020-03-15 04:15:09","context":"广东广州转运中心-已装袋发往-湖南长沙转运中心","location":null},{"time":"2020-03-15 03:44:43","ftime":"2020-03-15 03:44:43","context":"已到达-广东广州转运中心","location":null},{"time":"2020-03-15 00:44:08","ftime":"2020-03-15 00:44:08","context":"广东广州火车站点-已发往-广东广州转运中心","location":null},{"time":"2020-03-15 00:40:42","ftime":"2020-03-15 00:40:42","context":"广东广州火车站点-直营揽收(18138701110,020-37183133)-已收件","location":null},{"time":"2020-03-14 22:45:07","ftime":"2020-03-14 22:45:07","context":"广东广州火车站点-瑶台18(16602092703,020-37183133)-已收件","location":null}]}

}

type KuaidiCompanyQueryResDto struct {
	ExName string `json:"exname"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Url    string `json:"url"`
}

type KuaidiLogisticsQueryResDto struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	ExName  string `json:"exname"`
	Nu      string `json:"nu"`
	Reason  string `json:"reason"`
	Data    []struct {
		Context string `json:"context"`
		Time    string `json:"time"`
	} `json:"data"`
}
