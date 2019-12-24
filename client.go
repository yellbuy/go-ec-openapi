package ecopenapi

import (
	simplejson "github.com/bitly/go-simplejson"
	"github.com/yellbuy/go-ec-openapi/common"
)

type Client interface {
	Execute(method string, param common.Parameter) (res *simplejson.Json, err error)
}
