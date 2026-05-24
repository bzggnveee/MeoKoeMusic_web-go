package module

import "moekoe-go/util"

func init() { Register("/longaudio/album/detail", LongaudioAlbumDetail) }

func LongaudioAlbumDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var data []map[string]interface{}
	if s := toString(params["album_id"]); s != "" {
		for _, id := range splitCSV(s) { data = append(data, map[string]interface{}{"album_id": id}) }
	}
	return requestFn(util.RequestConfig{
		URL: "/openapi/v2/broadcast", Method: "POST",
		Data: map[string]interface{}{"data": data, "show_album_tag": 1, "fields": "album_name,album_id,category,authors,sizable_cover,intro,author_name,trans_param,album_tag,mix_intro,full_intro,is_publish"},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"KG-TID": "78"},
	})
}