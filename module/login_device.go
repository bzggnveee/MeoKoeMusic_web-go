package module

import ("moekoe-go/util"; "time")

func init() { Register("/login/device", LoginDevice) }

func LoginDevice(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	ct := time.Now().UnixMilli()
	enc := util.CryptoAesEncrypt(`{"token":"`+getVal(params,cookies,"token")+`"}`, "", "")
	return requestFn(util.RequestConfig{
		BaseURL: "https://userinfoservice.kugou.com", URL: "/v2/get_dev", Method: "POST",
		Data: map[string]interface{}{"plat": 1, "userid": getValInt(params,cookies,"userid"), "clienttime_ms": ct, "pk": util.CryptoRSAEncrypt(map[string]interface{}{"clienttime_ms": ct, "key": enc.Key}), "params": enc.Str},
		EncryptType: "android", Cookie: cookies,
	})
}