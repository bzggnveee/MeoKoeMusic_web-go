package module

import "moekoe-go/util"

func init() { Register("/rank/info", RankInfo) }
func RankInfo(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/ocean/v6/rank/info", Method: "GET",
		Params: map[string]interface{}{"rank_cid": toIntDefault(params, "rank_cid", 0), "rankid": params["rankid"], "with_album_img": toIntDefault(params, "album_img", 1), "zone": toString(params["zone"])},
		EncryptType: "android", Cookie: cookies,
	})
}