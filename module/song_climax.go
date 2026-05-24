package module

import ("encoding/json"; "moekoe-go/util")

func init() { Register("/song/climax", SongClimax) }
func SongClimax(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var data []map[string]interface{}
	for _, h := range splitCSV(toString(params["hash"])) { data = append(data, map[string]interface{}{"hash": h}) }
	b, _ := json.Marshal(data)
	return requestFn(util.RequestConfig{
		BaseURL: "https://expendablekmrcdn.kugou.com", URL: "/v1/audio_climax/audio", Method: "GET",
		Params: map[string]interface{}{"data": string(b)}, EncryptType: "android", Cookie: cookies,
	})
}