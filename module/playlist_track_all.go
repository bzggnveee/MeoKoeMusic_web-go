package module

import "moekoe-go/util"

func init() { Register("/playlist/track/all", PlaylistTrackAll) }
func PlaylistTrackAll(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	ps := toIntDefault(params, "pagesize", 30); p := toIntDefault(params, "page", 1)
	return requestFn(util.RequestConfig{
		URL: "/pubsongs/v2/get_other_list_file_nofilt", Method: "GET",
		Params: map[string]interface{}{"area_code": 1, "begin_idx": (p-1)*ps, "plat": 1, "type": 1, "mode": 1, "personal_switch": 1, "extend_fields": "abtags,hot_cmt,popularization", "pagesize": ps, "global_collection_id": params["id"]},
		EncryptType: "android", Cookie: cookies,
	})
}