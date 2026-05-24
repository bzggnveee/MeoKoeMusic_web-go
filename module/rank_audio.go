package module

import "moekoe-go/util"

func init() { Register("/rank/audio", RankAudio) }
func RankAudio(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/openapi/kmr/v2/rank/audio", Method: "POST",
		Data: map[string]interface{}{"show_portrait_mv": 1, "show_type_total": 1, "filter_original_remarks": 1, "area_code": 1, "pagesize": toIntDefault(params, "pagesize", 30), "rank_cid": toIntDefault(params, "rank_cid", 0), "type": 1, "page": toIntDefault(params, "page", 1), "rank_id": params["rankid"]},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"kg-tid": "369"},
	})
}