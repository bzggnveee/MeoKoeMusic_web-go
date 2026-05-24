package module

import "moekoe-go/util"

func init() { Register("/song/ranking/filter", SongRankingFilter) }
func SongRankingFilter(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/grow/v1/song_ranking/unlock/v2/ranking_filter", Method: "GET",
		Params: map[string]interface{}{"album_audio_id": params["album_audio_id"], "page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30)},
		EncryptType: "android", Cookie: cookies,
	})
}