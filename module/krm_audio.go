package module

import "moekoe-go/util"

func init() { Register("/krm/audio", KRMAudio) }

func KRMAudio(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var data []map[string]interface{}
	if s := toString(params["album_audio_id"]); s != "" {
		for _, id := range splitCSV(s) { data = append(data, map[string]interface{}{"entity_id": toInt(id)}) }
	}
	return requestFn(util.RequestConfig{
		URL: "/kmr/v2/audio", Method: "POST",
		Data: map[string]interface{}{"data": data, "fields": toStringDefault(params, "fields", "base")},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "openapi.kugou.com", "KG-TID": "238"},
	})
}