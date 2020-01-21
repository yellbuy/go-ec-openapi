package common

func NewSuccessResDto(isSuccess bool, code int, message, itemId string) *SuccessResDto {
	dto := new(SuccessResDto)
	dto.Response = new(successRes)
	if isSuccess {
		dto.Response.Flag = "success"
	} else {
		dto.Response.Flag = "failure"
	}
	dto.Response.Code = code
	dto.Response.Message = message
	dto.Response.ItemId = itemId
	return dto
}

func NewErrorResDto(code int, message string, subCode int, subMsg string) *ErrorResDto {
	dto := new(ErrorResDto)
	dto.Response = new(errorRes)
	dto.Response.Code = code
	dto.Response.Message = message
	dto.Response.SubCode = subCode
	dto.Response.SubMsg = subMsg
	return dto
}

type SuccessResDto struct {
	Response *successRes `json:"response"`
}

type ErrorResDto struct {
	Response *errorRes `json:"error_response"`
}

// 奇门下载成功响应内容
type successRes struct {
	//区名称（三级地址）
	Flag    string `json:"flag"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	ItemId  string `json:"itemId"`
}
type errorRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	SubCode int    `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

type Product struct {
	ProductId          string      `json:"platproductid"`
	ProductName        string      `json:"name"`
	ProductCode        string      `json:"outerid"`
	Price              string      `json:"price"`
	Num                string      `json:"num"`
	PictureUrl         string      `json:"pictureurl"`
	WhseCode           string      `json:"whsecode"`
	Attrbutes          interface{} `json:"attrbutes"`
	CategoryId         string      `json:"categoryid"`
	Status             string      `json:"status"`
	Statusdesc         string      `json:"statusdesc"`
	SkuList            []*Sku      `json:"skus"`
	Sendtype           string      `json:"sendtype"`
	Skutype            string      `json:"skutype"`
	PropertyAlias      string      `json:"propertyalias"`
	Isplatstorageorder string      `json:"isplatstorageorder"`
	Cooperationno      string      `json:"cooperationno"`
}
type Sku struct {
	SkuId         string `json:"skuid"`
	SkuCode       string `json:"skuouterid"`
	SkuPrice      string `json:"skuprice"`
	SkuQuantity   string `json:"skuquantity"`
	SkuName       string `json:"skuname"`
	SkuProperty   string `json:"skuproperty"`
	SkuType       string `json:"skutype"`
	SkuPictureUrl string `json:"skupictureurl"`
	SkuName2      string `json:"skuname2"`
}
