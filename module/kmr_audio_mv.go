package module

import "moekoe-go/util"

func init() { Register("/kmr/audio/mv", KMRAudioMV) }

func KMRAudioMV(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var data []map[string]interface{}
	if s := toString(params["album_audio_id"]); s != "" {
		for _, id := range splitCSV(s) { data = append(data, map[string]interface{}{"album_audio_id": id}) }
	}
	return requestFn(util.RequestConfig{
		URL: "/kmr/v1/audio/mv", Method: "POST",
		Data: map[string]interface{}{"data": data, "fields": toString(params["fields"])},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "openapi.kugou.com", "KG-TID": "38"},
	})
}