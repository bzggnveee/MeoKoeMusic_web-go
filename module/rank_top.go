package module

import "moekoe-go/util"

func init() { Register("/rank/top", RankTop) }
func RankTop(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL: "/mobileservice/api/v5/rank/rec_rank_list", Method: "GET", EncryptType: "android", Cookie: cookies})
}