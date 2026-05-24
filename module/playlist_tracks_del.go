package module

import "moekoe-go/util"

func init() { Register("/playlist/tracks/del", PlaylistTracksDel) }
func PlaylistTracksDel(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var data []map[string]interface{}
	for _, id := range splitCSV(toString(params["fileids"])) { data = append(data, map[string]interface{}{"fileid": toInt(id)}) }
	return requestFn(util.RequestConfig{
		URL: "/v4/delete_songs", Method: "POST",
		Data: map[string]interface{}{"listid": params["listid"], "userid": getValInt(params, cookies, "userid"), "data": data, "type": 0, "token": getVal(params, cookies, "token"), "list_ver": 0},
		EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "cloudlist.service.kugou.com"},
	})
}