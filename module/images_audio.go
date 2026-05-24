package module

import ("encoding/json"; "fmt"; "moekoe-go/util"; "net/url"; "sort"; "strings")

func init() { Register("/images/audio", ImagesAudio) }

func ImagesAudio(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var data []map[string]interface{}
	hashes := splitCSV(toString(params["hash"]))
	audioIDs := splitCSV(toString(params["audio_id"]))
	albumAudioIDs := splitCSV(toString(params["album_audio_id"]))
	filenames := splitCSV(toString(params["filename"]))
	for i, h := range hashes {
		d := map[string]interface{}{"audio_id": "0", "hash": h, "album_audio_id": "0", "filename": ""}
		if i < len(audioIDs) && audioIDs[i] != "" { d["audio_id"] = audioIDs[i] }
		if i < len(albumAudioIDs) && albumAudioIDs[i] != "" { d["album_audio_id"] = albumAudioIDs[i] }
		if i < len(filenames) { d["filename"] = filenames[i] }
		data = append(data, d)
	}
	pm := map[string]interface{}{"appid": util.GetAppID(), "clientver": util.GetClientVer(), "count": toIntDefault(params, "count", 5), "data": data, "isCdn": 1, "publish_time": 1, "show_authors": 1}
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
		URL: fmt.Sprintf("/v2/author_image/audio?%s", strings.Join(parts, "&")),
		Method: "GET", EncryptType: "android", Params: map[string]interface{}{"signature": sig},
		Cookie: cookies, NotSign: true, ClearDefaultParams: true,
	})
}