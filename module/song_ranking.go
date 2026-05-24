package module

import "moekoe-go/util"

func init() { Register("/song/ranking", SongRanking) }
func SongRanking(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/grow/v1/song_ranking/play_page/ranking_info", Method: "GET",
		Params: map[string]interface{}{"album_audio_id": params["album_audio_id"]},
		EncryptType: "android", Cookie: cookies,
	})
}