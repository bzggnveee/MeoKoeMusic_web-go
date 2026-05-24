package module

import "moekoe-go/util"

func init() { Register("/pc/diantai", PCDiantai) }

func PCDiantai(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		BaseURL: "https://adservice.kugou.com", URL: "/v3/pc_diantai", Method: "POST",
		Data: map[string]interface{}{"isvip": 0, "userid": getValInt(params, cookies, "userid"), "vipType": 0},
		EncryptType: "android", Cookie: cookies,
	})
}