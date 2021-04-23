package pdd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/yellbuy/go-ec-openapi/common"
)

func (client *Client) GetWaybill(request *common.WaybillApplyNewRequest, extData ...string) (*common.WaybillApplyNewCols, []byte, error) {
	req := make(map[string]interface{})
	req["waybill_apply_new_request"] = request
	params, err := common.InterfaceToParameter(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	res, err := client.Execute("pdd.waybill.get", params)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	data, err := res.Encode()
	if err != nil {
		fmt.Println(err)
		return nil, data, err
	}
	// fmt.Println("waybill_apply_new_cols:", string(data))
	dto := new(PddWaybillGetResponse)
	err = json.Unmarshal(data, dto)
	if err != nil {
		fmt.Println(err)
	}
	if len(dto.Modules) == 0 {
		return nil, data, errors.New("响应内容不完整")
	}
	result := new(common.WaybillApplyNewCols)
	result.WaybillApplyNewInfo = make([]*common.WaybillApplyNewInfo, len(dto.Modules))
	for index, val := range dto.Modules {
		info := new(common.WaybillApplyNewInfo)
		info.WaybillCode = val.WaybillCode
		info.PrintData = val.PrintData
		result.WaybillApplyNewInfo[index] = info
		// 根据收货地址返回大头笔信息
		// ShortAddress string `json:"short_address"`
		// // 返回的面单号
		// WaybillCode string `json:"waybill_code"`
		// // 集包地代码
		// PackageCenterCode string `json:"package_center_code"`
		// // 集包地名称
		// PackageCenterName string `json:"package_center_name"`
		// // 打印配置项，传给ali-print组件
		// PrintConfig string `json:"print_config"`
		// // 面单号对应的物流服务商网点（分支机构）代码
		// ShippingBranchCode string `json:"shipping_branch_code"`
		// // 包裹对应的派件（收件）物流服务商网点（分支机构）名称:余杭一部
		// ConsigneeBranchName string `json:"consignee_branch_name"`
		// // 面单号对于的物流服务商网点（分支机构）名称:西湖二部
		// ShippingBranchName string `json:"shipping_branch_name"`
		// // 包裹对应的派件（收件）物流服务商网点（分支机构）代码
		// ConsigneeBranchCode string `json:"consignee_branch_code"`
	}

	return result, data, err
}

type ParamWaybillCloudPrintApplyNewRequest struct {
	//设定取号返回的云打印报文是否加密
	NeedEncrypt bool `json:"need_encrypt"`
	// 物流公司Code
	WpCode string `json:"wp_code"`
	// 发货人信息
	Sender *Sender `json:"sender"`
	// 请求面单信息，数量限制为10
	TradeOrderInfoDtos []*TradeOrderInfoDto `json:"trade_order_info_dtos"`
}
type Sender struct {
	Address *Address `json:"address"`
	Mobile  string   `json:"mobile"`
	// 姓名
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
type Address struct {
	Country string `json:"country"`
	// 必填,省
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Town     string `json:"town"`
	// 必填.详细地址
	Detail string `json:"detail"`
}
type TradeOrderInfoDto struct {
	LogisticsServices string `json:"logistics_services"`
	// 必填	请求id
	ObjectId string `json:"object_id"`
	// 必填 订单信息
	OrderInfo *OrderInfo `json:"order_info"`
	// 必填 包裹信息
	PackageInfo *PackageInfo `json:"package_info"`
	// 非必填 收件人信息
	Recipient *Recipient `json:"recipient"`
	// 必填 标准模板模板URL
	TemplateUrl string `json:"template_url"`
	// 必填 使用者ID
	UserId int64 `json:"user_id"`
}
type OrderInfo struct {
	// 必填 订单渠道平台编码
	OrderChannelsType string `json:"order_channels_type"`
	// 必填 订单号,数量限制100
	TradeOrderList []string `json:"trade_order_list"`
}
type PackageInfo struct {
	// 快运货品描述
	GoodsDescription string `json:"goods_description"`
	// 包裹id,拆合单使用
	Id string `json:"id"`
	// 必填 商品信息,数量限制为100
	Items []*Item `json:"items"`
	// 非必填 快运包装方式描述
	PackagingDescription string `json:"packaging_description"`
	// 非必填 子母件总包裹数
	TotalPackagesCount int `json:"total_packages_count"`
	// 非必填 体积, 单位 ml
	Volume int64 `json:"volume"`
	// 非必填 重量,单位 g
	Weight int `json:"weight"`
}
type Item struct {
	// 必填 数量
	Count int `json:"count"`
	// 必填 名称
	Name string `json:"name"`
}
type Recipient struct {
	Address *Address `json:"address"`
	Mobile  string   `json:"mobile"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
}

// 拼多多面单响应内容
type PddWaybillGetResponse struct {
	Modules []*Module `json:"modules"`
}
type Module struct {
	// 请求id
	ObjectId string `json:"object_id"`
	//快运母单号
	ParentWaybillCode string `json:"parent_waybill_code"`
	// 面单信息
	PrintData string `json:"print_data"`
	// 面单号
	WaybillCode string `json:"waybill_code"`
}

func (self *ParamWaybillCloudPrintApplyNewRequest) getParamWaybillCloudPrintApplyNewRequest(req *common.WaybillApplyNewRequest) {
	self.NeedEncrypt = false
	self.WpCode = req.CpCode
	sender := new(Sender)
	sender.Address = new(Address)
	//sender.Address.Country=req.ShippingAddress.Country
	sender.Address.Province = req.ShippingAddress.Province
	sender.Address.City = req.ShippingAddress.City
	sender.Address.District = req.ShippingAddress.Area
	sender.Address.Town = req.ShippingAddress.Town
	sender.Address.Detail = req.ShippingAddress.AddressDetail
	// sender.Mobile=orderInfo.SendMobile
	if len(req.TradeOrderInfoCols) > 0 {
		orderInfo := req.TradeOrderInfoCols[0]
		sender.Phone = orderInfo.SendPhone
		sender.Name = orderInfo.SendName
	}

	self.Sender = sender

	tradeOrderInfoDtos := make([]*TradeOrderInfoDto, len(req.TradeOrderInfoCols))
	for index, val := range req.TradeOrderInfoCols {
		dto := new(TradeOrderInfoDto)
		if len(val.LogisticsServiceList) > 0 {
			for k, svc := range val.LogisticsServiceList {
				if k > 0 {
					dto.LogisticsServices = dto.LogisticsServices + ","
				}
				dto.LogisticsServices = dto.LogisticsServices + svc.ServiceCode

			}
		}
		dto.ObjectId = val.ObjectId
		// 订单信息
		orderInfo := new(OrderInfo)
		orderInfo.OrderChannelsType = val.OrderType
		if val.PackageId != "" {
			orderInfo.TradeOrderList = strings.Split(val.PackageId, ",")
		}
		dto.OrderInfo = orderInfo
		// 包裹信息
		packageInfo := new(PackageInfo)
		//packageInfo.GoodsDescription=val.ItemName
		packageInfo.Items = make([]*Item, len(val.PackageItems))
		for j, pkg := range val.PackageItems {
			item := new(Item)
			item.Count = pkg.Count
			item.Name = pkg.ItemName
			packageInfo.Items[j] = item
		}
		packageInfo.Volume = val.Volumn
		packageInfo.Weight = val.Weight

		dto.PackageInfo = packageInfo

		recipient := new(Recipient)
		recipient.Address = new(Address)
		//recipient.Address.Country = val.ConsigneeAddress.Country
		recipient.Address.Province = val.ConsigneeAddress.Province
		recipient.Address.City = val.ConsigneeAddress.City
		recipient.Address.District = val.ConsigneeAddress.Area
		recipient.Address.Town = val.ConsigneeAddress.Town
		recipient.Address.Detail = val.ConsigneeAddress.AddressDetail
		// sender.Mobile=orderInfo.SendMobile
		recipient.Phone = val.ConsigneePhone
		recipient.Name = val.ConsigneeName
		dto.Recipient = recipient
		// 必填 标准模板模板URL
		dto.TemplateUrl = ""
		dto.UserId = val.RealUserId

		tradeOrderInfoDtos[index] = dto
	}

	self.TradeOrderInfoDtos = tradeOrderInfoDtos
}
func (client *Client) GetWaybillTemplates(request *common.WaybillTemplateRequest, extData ...string) (res *common.WaybillTemplateDto, body []byte, err error) {
	err = errors.New("未实现")
	return
}
func (client *Client) CancelWaybill(request []common.WaybillCancel, extData ...string) (*common.WaybillCancelReturn, error) {
	err := errors.New("未实现")
	return nil, err
}

func (client *Client) DownloadOrderListV2(request common.DownLoadOrderListPostBizcontent, extData ...string) (common.DownloadOrderListReturn, error) {
	err := errors.New("未实现")
	var OutData common.DownloadOrderListReturn
	return OutData, err
}

func (client *Client) CheckRefundV2(request common.BatchCheckRefundStatusBizcontent, extData ...string) (common.CheckRefundReturn, error) {
	err := errors.New("未实现")
	var OutData common.CheckRefundReturn
	return OutData, err
}
