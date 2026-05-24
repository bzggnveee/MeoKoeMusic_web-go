package module

import "moekoe-go/util"

func init() { Register("/longaudio/week/recommend", LongaudioWeekRecommend) }
func LongaudioWeekRecommend(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL: "/longaudio/v1/home_new/week_new_albums_recommend", Method: "POST", Data: map[string]interface{}{"album_playlist": []string{}}, Params: map[string]interface{}{"clientver": 12329}, EncryptType: "android", Cookie: cookies})
}