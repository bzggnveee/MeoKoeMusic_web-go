package module

import ("moekoe-go/util"; "time")

func init() { Register("/theme/music/detail", ThemeMusicDetail) }
func ThemeMusicDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/everydayrec.service/v1/theme_category_recommend", Method: "POST",
		Data: map[string]interface{}{"platform": "android", "clienttime": time.Now().Unix(), "theme_category_id": params["id"], "show_theme_category_id": 0, "userid": getValInt(params, cookies, "userid"), "module_id": 508},
		EncryptType: "android", Cookie: cookies,
	})
}