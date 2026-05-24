package module

import "moekoe-go/util"

func init() { Register("/user/playlist", UserPlaylist) }
func UserPlaylist(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	userid := getValInt(params, cookies, "userid")
	token := getVal(params, cookies, "token")
	return requestFn(util.RequestConfig{
		URL: "/v7/get_all_list", Method: "POST",
		Data: map[string]interface{}{
			"userid": userid, "token": token, "total_ver": 979, "type": 2,
			"page": toIntDefault(params, "page", 1),
			"pagesize": toIntDefault(params, "pagesize", 30),
		},
		Params: map[string]interface{}{
			"plat": 1, "userid": userid, "token": token,
			"t": toString(params["t"]),
		},
		EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "cloudlist.service.kugou.com"},
	})
}