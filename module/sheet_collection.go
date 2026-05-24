package module

import "moekoe-go/util"

func init() { Register("/sheet/collection", SheetCollection) }
func SheetCollection(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	p := 2; if v := toInt(params["position"]); v != 0 { p = v }
	return requestFn(util.RequestConfig{URL: "/miniyueku/v1/opern_square/get_home_module_config", Method: "GET", Params: map[string]interface{}{"srcappid": util.SrcAppID, "position": p}, EncryptType: "web", Cookie: cookies})
}