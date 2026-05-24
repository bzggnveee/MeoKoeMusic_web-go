package module

import "moekoe-go/util"

func init() { Register("/user/cloud/url", UserCloudURL) }
func UserCloudURL(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	h := toLowerCase(toString(params["hash"]))
	return requestFn(util.RequestConfig{
		URL: "/bsstrackercdngz/v2/query_musicclound_url", Method: "GET",
		Params: map[string]interface{}{"hash": h, "ssa_flag": "is_fromtrack", "version": "20102", "ssl": 0, "album_audio_id": toIntDefault(params, "album_audio_id", 0), "pid": 20026, "audio_id": toIntDefault(params, "audio_id", 0), "kv_id": 2, "key": util.SignCloudKey(h, "20026"), "bucket": "musicclound", "name": toString(params["name"]), "with_res_tag": 0},
		EncryptType: "android", Cookie: cookies,
	})
}