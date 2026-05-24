package module

import "moekoe-go/util"

func init() { Register("/sheet/detail", SheetDetail) }
func SheetDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL: "/opern/v1/detail/info", Method: "GET", Params: map[string]interface{}{"opern_id": params["id"]}, EncryptType: "android", Cookie: cookies})
}