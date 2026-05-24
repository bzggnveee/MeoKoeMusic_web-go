package module

import ("moekoe-go/util"; "time")

func init() { Register("/theme/playlist/track", ThemePlaylistTrack) }
func ThemePlaylistTrack(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/v2/gettheme_songidlist", Method: "POST",
		Data: map[string]interface{}{"platform": "android", "clientver": util.GetClientVer(), "clienttime": time.Now().UnixMilli(), "area_code": 1, "module_id": 1, "userid": getValInt(params, cookies, "userid"), "theme_id": params["theme_id"]},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "everydayrec.service.kugou.com"},
	})
}