package module

import "moekoe-go/util"

func init() { Register("/ip/zone/home", IPZoneHome) }

func IPZoneHome(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/v1/zone/home", Method: "GET",
		Params: map[string]interface{}{"id": params["id"], "share": 0},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "yuekucategory.kugou.com"},
	})
}