package module

import "moekoe-go/util"

func init() { Register("/playlist/track/all/new", PlaylistTrackAllNew) }
func PlaylistTrackAllNew(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/v4/get_list_all_file", Method: "POST",
		Data: map[string]interface{}{"listid": params["listid"], "userid": getVal(params, cookies, "userid"), "area_code": 1, "show_relate_goods": 0, "pagesize": toIntDefault(params, "pagesize", 30), "allplatform": 1, "show_cover": 1, "type": 0, "token": getVal(params, cookies, "token"), "page": toIntDefault(params, "page", 1)},
		EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "cloudlist.service.kugou.com"},
	})
}