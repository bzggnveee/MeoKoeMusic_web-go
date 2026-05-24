package module

import "moekoe-go/util"

func init() { Register("/sheet/rank", SheetRank) }
func SheetRank(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/opern/v1/home/get_rank_opern", Method: "POST",
		Params: map[string]interface{}{"pagesize": toIntDefault(params, "pagesize", 30), "page": toIntDefault(params, "page", 1), "opern_level": toIntDefault(params, "level", 0), "instruments": toIntDefault(params, "instruments", 1), "tagid": toIntDefault(params, "tagid", 0)},
		EncryptType: "android", Cookie: cookies,
	})
}