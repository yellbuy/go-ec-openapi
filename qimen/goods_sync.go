/*
文档地址：https://qimen.taobao.com/qimen/index.htm#/api-doc?fromPage=api-list&&apiServiceId=&officialApiId=27262&_k=oulc05
*/

package qimen

import (
	"encoding/xml"

	"github.com/yellbuy/go-ec-openapi/common"
)

// 产品异步下载同步
func GoodsSyncParse(body []byte) (res []*common.Product, warehouseCode, ownerCode string, err error) {
	res = make([]*common.Product, 0)
	dto := new(GoodsSyncDto)
	err = xml.Unmarshal(body, dto)
	if err != nil {
		return
	}
	//fmt.Printf("%+v", dto)
	warehouseCode = dto.WarehouseCode
	ownerCode = dto.OwnerCode

	for _, goods := range dto.Items {
		product := new(common.Product)
		product.ProductId = goods.ItemId
		product.ProductCode = goods.ItemCode
		product.ProductName = goods.ItemName

		product.SkuList = make([]*common.Sku, 1)
		sku := new(common.Sku)
		if goods.SkuProperty == "" {
			sku.SkuId = product.ProductId
			sku.SkuCode = product.ProductCode
			sku.SkuName = product.ProductName
		} else {
			sku.SkuId = goods.SkuProperty
			sku.SkuCode = goods.SkuProperty
			sku.SkuName = goods.SkuProperty
		}

		product.SkuList[0] = sku
		res = append(res, product)
	}
	return
}

type GoodsSyncDto struct {
	// 操作类型(两种类型：add|update)
	ActionType string `xml:"actionType"`
	//仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	WarehouseCode string `xml:"warehouseCode"`
	//货主编码
	OwnerCode string           `xml:"ownerCode"`
	Items     []*GoodsSyncItem `xml:"item"`
}
type GoodsSyncItem struct {
	// 仓储系统商品编码(该字段是WMS分配的商品编号;WMS如果分配了商品编码;则后续的商品操作都需要传该字段;如果WMS不使用 ;WMS可 以返回itemId=itemCode的值)
	ItemId string `xml:"itemId"`
	// 商品编码
	ItemCode string `xml:"itemCode"`
	// 商品名称
	ItemName string `xml:"itemName"`
	// 商品名称
	ShortName string `xml:"shortName"`
	//条形码
	BarCode string `xml:"barCode"`
	//商品属性
	SkuProperty string `xml:"skuProperty"`
	// 商品属性
	StockUnit string `xml:"stockUnit"`
}

type GoodsSyncResponseDto struct {
	XMLName xml.Name `xml:"response"`
	Response
	Items []*GoodsSyncResponseItem `xml:"items,omitempty"`
}
type GoodsSyncResponseItem struct {
	ItemCode string `xml:"itemCode"`
	Message  string `xml:"message"`
}

func NewGoodsSyncSuccessResponse() *GoodsSyncResponseDto {
	dto := new(GoodsSyncResponseDto)
	dto.Flag = "success"
	dto.Code = "0"
	return dto
}

func (self *GoodsSyncResponseDto) AddItem(itemCode, message string) {
	item := new(GoodsSyncResponseItem)
	item.ItemCode = itemCode
	item.Message = message
	if len(self.Items) == 0 {
		self.Items = make([]*GoodsSyncResponseItem, 1)
		self.Items[0] = item
	} else {
		self.Items = append(self.Items, item)
	}
}
