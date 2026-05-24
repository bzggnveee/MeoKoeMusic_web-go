package module

import "moekoe-go/util"

func init() { Register("/sheet/song", SheetSong) }
func SheetSong(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/opern/v1/detail/song_info", Method: "GET",
		Params: map[string]interface{}{"mixsongid": params["album_audio_id"], "instruments": toIntDefault(params, "instruments", 1), "opern_level": toIntDefault(params, "level", 0)},
		EncryptType: "android", Cookie: cookies,
	})
}