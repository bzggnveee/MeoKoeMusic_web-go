package module

import "moekoe-go/util"

func init() { Register("/login/device/kick", LoginDeviceKick) }

func LoginDeviceKick(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/loginservice/v1/dev_logout", Method: "GET",
		Params: map[string]interface{}{"mid": cookies["KUGOU_API_MID"], "userid": getValInt(params,cookies,"userid"), "token": getVal(params,cookies,"token")},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"Host": "gateway.kugou.com"},
	})
}