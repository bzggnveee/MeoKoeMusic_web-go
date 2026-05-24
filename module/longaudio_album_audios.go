package module

import "moekoe-go/util"

func init() { Register("/longaudio/album/audios", LongaudioAlbumAudios) }

func LongaudioAlbumAudios(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/longaudio/v2/album_audios", Method: "POST",
		Data: map[string]interface{}{"album_id": params["album_id"], "area_code": 1, "tagid": 0, "page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30)},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "openapi.kugou.com", "KG-TID": "78"},
	})
}