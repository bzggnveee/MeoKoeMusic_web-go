package module

import ("moekoe-go/util"; "time")

func init() { Register("/theme/playlist", ThemePlaylist) }
func ThemePlaylist(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/v2/getthemelist", Method: "POST",
		Data: map[string]interface{}{"platform": "android", "clientver": util.GetClientVer(), "clienttime": time.Now().UnixMilli(), "area_code": 1, "module_id": 1, "userid": getValInt(params, cookies, "userid")},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "everydayrec.service.kugou.com"},
	})
}