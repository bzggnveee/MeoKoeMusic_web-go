package module

import "moekoe-go/util"

func init() { Register("/search/lyric", SearchLyric) }
func SearchLyric(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		BaseURL: "https://lyrics.kugou.com", URL: "/v1/search", Method: "GET",
		Params: map[string]interface{}{"album_audio_id": toIntDefault(params, "album_audio_id", 0), "appid": util.GetAppID(), "clientver": util.GetClientVer(), "duration": toIntDefault(params, "duration", 0), "hash": toString(params["hash"]), "keyword": toString(params["keywords"]), "lrctxt": 1, "man": toStringDefault(params, "man", "no")},
		Cookie: cookies, EncryptType: "android", ClearDefaultParams: true, NotSign: true,
	})
}