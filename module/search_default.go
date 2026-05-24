package module

import "moekoe-go/util"

func init() { Register("/search/default", SearchDefault) }
func SearchDefault(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	uid := getValInt(params, cookies, "userid")
	vt := toInt(getVal(params, cookies, "vip_type")); if vt == 0 { vt = 65530 }
	return requestFn(util.RequestConfig{
		URL: "/searchnofocus/v1/search_no_focus_word", Method: "POST",
		Data: map[string]interface{}{"plat": 0, "userid": uid, "tags": "{}", "vip_type": vt, "m_type": 0, "own_ads": map[string]interface{}{}, "ability": "3", "sources": []string{}, "bitmap": 2, "mode": "normal"},
		Params: map[string]interface{}{"clientver": 12329},
		EncryptType: "android", Cookie: cookies,
	})
}