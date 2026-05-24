package module

import "moekoe-go/util"

func init() { Register("/privilege/lite", PrivilegeLite) }
func PrivilegeLite(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	hashes := splitCSV(toString(params["hash"])); albumIDs := splitCSV(toString(params["album_id"]))
	var res []map[string]interface{}
	for i, h := range hashes {
		m := map[string]interface{}{"type": "audio", "page_id": 0, "hash": h, "album_id": "0"}
		if i < len(albumIDs) && albumIDs[i] != "" { m["album_id"] = albumIDs[i] }
		res = append(res, m)
	}
	return requestFn(util.RequestConfig{
		URL: "/v2/get_res_privilege/lite", Method: "POST",
		Data: map[string]interface{}{"appid": util.GetAppID(), "area_code": 1, "behavior": "play", "clientver": util.GetClientVer(), "need_hash_offset": 1, "relate": 1, "support_verify": 1, "resource": res, "qualities": []string{"128", "320", "flac", "high", "viper_atmos", "viper_tape", "viper_clear", "super", "multitrack"}},
		EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "media.store.kugou.com", "Content-Type": "application/json"},
	})
}