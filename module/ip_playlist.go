package module

import "moekoe-go/util"

func init() { Register("/ip/playlist", IPPlaylist) }

func IPPlaylist(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/ocean/v6/pubsongs/list_info_for_ip", Method: "POST",
		Params: map[string]interface{}{"ip": params["id"], "page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30)},
		EncryptType: "android", Cookie: cookies,
	})
}