package module

import ("moekoe-go/util"; "time"; "strconv")

func init() { Register("/fm/songs", FMSongs) }

func FMSongs(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dt := time.Now().UnixMilli(); uid := getVal(params, cookies, "userid")
	var fd []map[string]interface{}
	fmids := splitCSV(toString(params["fmid"]))
	fmtypes := splitCSV(toString(params["fmtype"]))
	fmoffsets := splitCSV(toString(params["fmoffset"]))
	fmsizes := splitCSV(toString(params["fmsize"]))
	for i, id := range fmids {
		ft := 2; if i < len(fmtypes) { if n, _ := strconv.Atoi(fmtypes[i]); n != 0 { ft = n } }
		fo := -1; if i < len(fmoffsets) { if n, _ := strconv.Atoi(fmoffsets[i]); n != 0 { fo = n } }
		fs := 20; if i < len(fmsizes) { if n, _ := strconv.Atoi(fmsizes[i]); n != 0 { fs = n } }
		fd = append(fd, map[string]interface{}{"fmid": id, "fmtype": ft, "offset": fo, "size": fs, "singername": ""})
	}
	return requestFn(util.RequestConfig{
		URL: "/v1/app_song_list_offset", Method: "POST",
		Data: map[string]interface{}{"appid": util.GetAppID(), "area_code": 1, "clienttime": dt, "clientver": util.GetClientVer(), "data": fd, "get_tracker": 1, "key": util.SignParamsKey(dt), "mid": cookies["KUGOU_API_MID"], "uid": uid},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "fm.service.kugou.com", "Content-Type": "application/json"},
	})
}