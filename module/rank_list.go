package module

import (
	"moekoe-go/util"
)

func init() {
	Register("/rank/list", RankList)
}

// RankList 排行榜列表
func RankList(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:         "/v1/rank/list",
		Method:      "GET",
		Params: map[string]interface{}{
			"withsong": 1,
		},
		EncryptType: "android",
		Headers: map[string]string{
			"x-router": "rank.kugou.com",
		},
		Cookie: cookies,
	})
}
