package module

import ("moekoe-go/util"; "time"; "strconv")

func init() { Register("/playlist/tracks/add", PlaylistTracksAdd) }
func PlaylistTracksAdd(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	uid := getValInt(params, cookies, "userid"); tk := getVal(params, cookies, "token"); ct := time.Now().Unix()
	var res []map[string]interface{}
	for _, s := range splitCSV(toString(params["data"])) {
		parts := splitByPipe(s)
		m := map[string]interface{}{"number": 1, "name": "", "hash": "", "size": 0, "sort": 0, "timelen": 0, "bitrate": 0, "album_id": 0, "mixsongid": 0}
		if len(parts) > 0 { m["name"] = parts[0] }
		if len(parts) > 1 { m["hash"] = parts[1] }
		if len(parts) > 2 { if n, _ := strconv.Atoi(parts[2]); n != 0 { m["album_id"] = n } }
		if len(parts) > 3 { if n, _ := strconv.Atoi(parts[3]); n != 0 { m["mixsongid"] = n } }
		res = append(res, m)
	}
	return requestFn(util.RequestConfig{
		URL: "/cloudlist.service/v6/add_song", Method: "POST",
		Data: map[string]interface{}{"userid": uid, "token": tk, "listid": params["listid"], "list_ver": 0, "type": 0, "slow_upload": 1, "scene": "false;null", "data": res},
		Params: map[string]interface{}{"last_time": ct, "last_area": "gztx", "userid": uid, "token": tk},
		EncryptType: "android", Cookie: cookies,
	})
}