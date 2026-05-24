package module

import "moekoe-go/util"

func init() { Register("/server/now", ServerNow) }
func ServerNow(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/v1/server_now", Method: "POST", Data: map[string]interface{}{"token": getVal(params, cookies, "token"), "userid": getValInt(params, cookies, "userid")},
		Params: map[string]interface{}{"plat": 3}, EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "usercenter.kugou.com"},
	})
}