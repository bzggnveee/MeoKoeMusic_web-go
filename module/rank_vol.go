package module

import "moekoe-go/util"

func init() { Register("/rank/vol", RankVol) }
func RankVol(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/ocean/v6/rank/vol", Method: "GET",
		Params: map[string]interface{}{"rank_cid": toIntDefault(params, "rank_cid", 0), "rankid": params["rankid"], "ranktype": 1, "type": 0, "plat": 2},
		EncryptType: "android", Cookie: cookies,
	})
}