package qimen

import (
	"encoding/xml"
)

type Response struct {
	XMLName xml.Name `xml:"response"`
	//响应结果:success|failure
	Flag string `xml:"flag"`
	//响应码
	Code string `xml:"code"`
	//响应码
	Message string `xml:"message"`
}

func NewSuccessResponse(message string) *Response {
	dto := new(Response)
	dto.Flag = "success"
	dto.Code = "0"
	dto.Message = message
	return dto
}

func NewFailResponse(code, message string) *Response {
	dto := new(Response)
	dto.Flag = "failure"
	dto.Code = code
	dto.Message = message
	return dto
}

type JsonResponse struct {
	Code            string `json:"code"`
	DeliveryOrderId string `json:"deliveryOrderId"`
	Flag            string `json:"flag"`    //响应结果:success|failure
	Message         string `json:"message"` //响应信息
}
type PddResponse struct {
	Response *JsonResponse `json:"response"`
}

func NewPDDSuccessResponse(message string) *PddResponse {
	dto := new(JsonResponse)
	dto.Flag = "success"
	dto.Code = "0"
	dto.Message = message
	return &PddResponse{Response: dto}
}
func NewPDDFailResponse(code, message string) *PddResponse {
	dto := new(JsonResponse)
	dto.Flag = "failure"
	dto.Code = code
	dto.Message = message
	return &PddResponse{Response: dto}
}
