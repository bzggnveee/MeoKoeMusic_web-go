package module

import "moekoe-go/util"

func init() { Register("/ip", IP) }

func IP(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	t := toString(params["type"])
	if t != "audios" && t != "albums" && t != "videos" && t != "author_list" { t = "audios" }
	return requestFn(util.RequestConfig{
		URL: "/openapi/v1/ip/" + t, Method: "POST",
		Data: map[string]interface{}{"is_publish": 1, "ip_id": params["id"], "sort": 3, "page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30), "query": 1},
		EncryptType: "android", Cookie: cookies,
	})
}