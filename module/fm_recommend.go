package module

import ("moekoe-go/util"; "time")

func init() { Register("/fm/recommend", FMRecommend) }

func FMRecommend(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dt := time.Now().UnixMilli()
	return requestFn(util.RequestConfig{
		URL: "/v1/rcmd_list", Method: "POST",
		Data: map[string]interface{}{"appid": util.GetAppID(), "clientver": util.GetClientVer(), "clienttime": dt, "mid": cookies["KUGOU_API_MID"], "key": util.SignParamsKey(dt), "rcmdsongcount": 1, "level": 0, "area_code": 1, "get_tracker": 1, "uid": 0},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "fm.service.kugou.com"},
	})
}