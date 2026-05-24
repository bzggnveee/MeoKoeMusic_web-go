package module

import ("moekoe-go/util"; "time")

func init() { Register("/song/url/new", SongURLNew) }
func SongURLNew(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	ct := time.Now().UnixMilli()
	uid := getValInt(params, cookies, "userid")
	vt := getValInt(params, cookies, "vip_type")
	vipToken := getVal(params, cookies, "vip_token")
	token := getVal(params, cookies, "token")
	dfid := dfidVal(params, cookies)
	hash := toString(params["hash"])
	fp := 0; if toBool(params["free_part"]) { fp = 1 }
	mc := make(map[string]string)
	for k, v := range cookies { mc[k] = v }
	mc["dfid"] = dfid
	return requestFn(util.RequestConfig{
		BaseURL: "http://tracker.kugou.com", URL: "/v6/priv_url", Method: "POST",
		Data: map[string]interface{}{
			"area_code": "1", "behavior": "play",
			"qualities": []string{"128", "320", "flac", "high", "multitrack", "viper_atmos", "viper_tape", "viper_clear", "super"},
			"resource": map[string]interface{}{"album_audio_id": params["album_audio_id"], "collect_list_id": "3", "collect_time": ct, "hash": hash, "id": 0, "page_id": 1, "type": "audio"},
			"token": token,
			"tracker_param": map[string]interface{}{"all_m": 1, "auth": "", "is_free_part": fp, "key": util.SignKey(hash, cookies["KUGOU_API_MID"], uid, util.GetAppID()), "module_id": 0, "need_climax": 1, "need_xcdn": 1, "open_time": "", "pid": "411", "pidversion": "3001", "priv_vip_type": "6", "viptoken": vipToken},
			"userid": toString(uid), "vip": vt,
		},
		EncryptType: "android", Cookie: mc,
	})
}