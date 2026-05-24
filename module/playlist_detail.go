package module

import "moekoe-go/util"

func init() { Register("/playlist/detail", PlaylistDetail) }

func PlaylistDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var data []map[string]interface{}
	if s := toString(params["ids"]); s != "" {
		for _, id := range splitCSV(s) { data = append(data, map[string]interface{}{"global_collection_id": id}) }
	}
	return requestFn(util.RequestConfig{
		URL: "/v3/get_list_info", Method: "POST",
		Data: map[string]interface{}{"data": data, "userid": getValInt(params, cookies, "userid"), "token": getVal(params, cookies, "token")},
		EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "pubsongs.kugou.com"},
	})
}