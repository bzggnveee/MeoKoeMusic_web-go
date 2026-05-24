package module

import "moekoe-go/util"

func init() { Register("/search/suggest", SearchSuggest) }
func SearchSuggest(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/v2/getSearchTip", Method: "GET",
		Params: map[string]interface{}{"keyword": params["keywords"], "AlbumTipCount": toIntDefault(params, "albumTipCount", 10), "CorrectTipCount": toIntDefault(params, "correctTipCount", 10), "MVTipCount": toIntDefault(params, "mvTipCount", 10), "MusicTipCount": toIntDefault(params, "musicTipCount", 10), "radiotip": 1},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "searchtip.kugou.com"},
	})
}