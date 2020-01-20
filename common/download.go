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
