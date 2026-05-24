package module

import (
	"moekoe-go/util"
)

func init() {
	Register("/recommend/songs", EverydayRecommend)
}

// EverydayRecommend 每日推荐歌曲
func EverydayRecommend(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:         "/v2/daily_recommend",
		Method:      "GET",
		Params: map[string]interface{}{
			"with_rcmd_extra": 1,
		},
		EncryptType: "android",
		Headers: map[string]string{
			"x-router": "everydayrec.kugou.com",
		},
		Cookie: cookies,
	})
}
