package module

import "moekoe-go/util"

func init() { Register("/longaudio/rank/recommend", LongaudioRankRecommend) }
func LongaudioRankRecommend(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL: "/longaudio/v1/home_new/rank_card_recommend", Method: "GET", Params: map[string]interface{}{"platform": "ios"}, EncryptType: "android", Cookie: cookies})
}