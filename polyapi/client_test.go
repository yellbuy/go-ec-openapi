package polyapi

import (
	"fmt"
	"testing"

	simplejson "github.com/bitly/go-simplejson"
)

func init() {

}
func TestExecute(t *testing.T) {
	// http://open.taobao.com/docs/api.htm?apiId=24515
	resJson := simplejson.New()

	resJson.Set("a", "a1")
	resJson.Set("b", "b1")

	bizcontent, err := resJson.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("bizcontent：", string(bizcontent))
	bizcontent, err = aesEncrypt("5ee2084de90043be989d4d99d0dd0eaa", bizcontent)
	if err != nil {
		fmt.Println(err)
		return
	}
	// bizcontent, err = aesEncrypt("5ee2084de90043be989d4d99d0dd0eaa", bizcontent)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// 5E309FC10461C44CF62554E0BE5DFF16447A35830F3084F40ABB042D23F381EF
	str := byteArrToHexString(bizcontent)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("bizcontent:", str)
}
