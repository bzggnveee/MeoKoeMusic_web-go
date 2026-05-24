package module

import ("moekoe-go/util"; "time")

func init() { Register("/playlist/similar", PlaylistSimilar) }
func PlaylistSimilar(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	ct := time.Now().UnixMilli()
	var data []map[string]interface{}
	for _, id := range splitCSV(toString(params["ids"])) { data = append(data, map[string]interface{}{"global_collection_id": id}) }
	return requestFn(util.RequestConfig{
		URL: "/pubsongs/v1/kmr_get_similar_lists", Method: "POST",
		Data: map[string]interface{}{"appid": util.GetAppID(), "clientver": util.GetClientVer(), "clienttime": ct, "key": util.SignParamsKey(ct), "userid": getValInt(params, cookies, "userid"), "ugc": 1, "show_list": 1, "need_songs": 1, "data": data},
		EncryptType: "android", Cookie: cookies,
	})
}