package module

import "moekoe-go/util"

func init() { Register("/scene/lists", SceneLists) }
func SceneLists(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL: "/scene/v1/scene/list", Method: "GET", EncryptType: "android", Cookie: cookies})
}