package module

import ("moekoe-go/util"; "time")

func init() { Register("/top/playlist", TopPlaylist) }
func TopPlaylist(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dt := time.Now().Unix()
	return requestFn(util.RequestConfig{
		URL: "/v2/special_recommend", Method: "POST",
		Data: map[string]interface{}{
			"appid": util.GetAppID(), "mid": cookies["KUGOU_API_MID"], "clientver": util.GetClientVer(), "platform": "android", "clienttime": dt,
			"userid": getValInt(params, cookies, "userid"), "module_id": toIntDefault(params, "module_id", 1), "page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30),
			"key": util.SignParamsKey(dt), "special_recommend": map[string]interface{}{"withtag": toIntDefault(params, "withtag", 1), "withsong": toIntDefault(params, "withsong", 1), "sort": toIntDefault(params, "sort", 1), "ugc": 1, "is_selected": 0, "withrecommend": 1, "area_code": 1, "categoryid": toIntDefault(params, "category_id", 0)},
			"req_multi": 1, "retrun_min": 5, "return_special_falg": 1,
		},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "specialrec.service.kugou.com"},
	})
}