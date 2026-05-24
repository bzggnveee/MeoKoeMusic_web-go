package module

import ("moekoe-go/util"; "time")

func init() { Register("/fm/class", FMClass) }

func FMClass(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dt := time.Now().UnixMilli()
	uid := getValInt(params, cookies, "userid")
	return requestFn(util.RequestConfig{
		URL: "/v1/class_fm_song", Method: "POST",
		Data: map[string]interface{}{"kguid": uid, "clienttime": dt, "mid": cookies["KUGOU_API_MID"], "platform": "android", "clientver": util.GetClientVer(), "uid": uid, "get_tracker": 1, "key": util.SignParamsKey(dt), "appid": util.GetAppID()},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "fm.service.kugou.com"},
	})
}