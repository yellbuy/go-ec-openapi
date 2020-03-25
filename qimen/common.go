package qimen

import "encoding/xml"

type Response struct {
	XMLName xml.Name `xml:"response"`
	//响应结果:success|failure
	Flag string `xml:"flag"`
	//响应码
	Code string `xml:"code"`
	//响应码
	Message string `xml:"message"`
}

func NewSuccessResponse() *Response {
	dto := new(Response)
	dto.Flag = "success"
	dto.Code = "0"
	return dto
}

func NewFailResponse(code, message string) *Response {
	dto := new(Response)
	dto.Flag = "failure"
	dto.Code = code
	dto.Message = message
	return dto
}
