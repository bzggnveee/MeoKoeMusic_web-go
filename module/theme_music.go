package module

import ("moekoe-go/util"; "time")

func init() { Register("/theme/music", ThemeMusic) }
func ThemeMusic(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/everydayrec.service/v1/mul_theme_category_recommend", Method: "POST",
		Data: map[string]interface{}{"platform": "android", "clienttime": time.Now().Unix(), "show_theme_category_ids": params["ids"], "userid": getValInt(params, cookies, "userid"), "module_id": 508},
		EncryptType: "android", Cookie: cookies,
	})
}