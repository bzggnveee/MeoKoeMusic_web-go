package module

import ("encoding/json"; "fmt"; "moekoe-go/util"; "net/url"; "sort"; "strings")

func init() { Register("/images", Images) }

func Images(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var data []map[string]interface{}
	hashes := splitCSV(toString(params["hash"]))
	albumIDs := splitCSV(toString(params["album_id"]))
	audioIDs := splitCSV(toString(params["album_audio_id"]))
	for i, h := range hashes {
		d := map[string]interface{}{"album_id": "0", "hash": h, "album_audio_id": "0"}
		if i < len(albumIDs) && albumIDs[i] != "" { d["album_id"] = albumIDs[i] }
		if i < len(audioIDs) && audioIDs[i] != "" { d["album_audio_id"] = audioIDs[i] }
		data = append(data, d)
	}
	pm := map[string]interface{}{"album_image_type": "-3", "appid": util.GetAppID(), "clientver": util.GetClientVer(), "author_image_type": "3,4,5", "count": toIntDefault(params, "count", 5), "data": data, "isCdn": 1, "publish_time": 1}
	sig := util.SignatureAndroidParams(pm, "")

	keys := make([]string, 0, len(pm)); for k := range pm { keys = append(keys, k) }
	sort.Strings(keys)
	var parts []string
	for _, k := range keys {
		v := pm[k]; var vs string
		if b, err := json.Marshal(v); err == nil { vs = string(b) } else { vs = fmt.Sprintf("%v", v) }
		parts = append(parts, fmt.Sprintf("%s=%s", k, url.QueryEscape(vs)))
	}
	return requestFn(util.RequestConfig{
		BaseURL: "https://expendablekmr.kugou.com",
		URL: fmt.Sprintf("/container/v2/image?%s", strings.Join(parts, "&")),
		Method: "GET", EncryptType: "android", Params: map[string]interface{}{"signature": sig},
		Cookie: cookies, NotSign: true, ClearDefaultParams: true,
	})
}