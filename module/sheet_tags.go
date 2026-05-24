package module

import "moekoe-go/util"

func init() { Register("/sheet/tags", SheetTags) }
func SheetTags(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL: "/opern/v1/home/get_tags", Method: "GET", EncryptType: "android", Cookie: cookies})
}