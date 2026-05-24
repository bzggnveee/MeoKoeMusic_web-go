package module

import "moekoe-go/util"

func init() { Register("/recommend/songs", RecommendSongs) }
func RecommendSongs(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/everyday_song_recommend", Method: "POST",
		Data: map[string]interface{}{"platform": toStringDefault(params, "platform", "android"), "userid": toString(getVal(params, cookies, "userid"))},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "everydayrec.service.kugou.com"},
	})
}