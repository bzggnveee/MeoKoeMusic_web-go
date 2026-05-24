package module

import "moekoe-go/util"

func init() { Register("/longaudio/daily/recommend", LongaudioDailyRecommend) }
func LongaudioDailyRecommend(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL: "/longaudio/v1/home_new/daily_recommend", Method: "POST", Params: map[string]interface{}{"module_id": 1, "size": toIntDefault(params, "pagesize", 30), "page": toIntDefault(params, "page", 1)}, EncryptType: "android", Cookie: cookies})
}