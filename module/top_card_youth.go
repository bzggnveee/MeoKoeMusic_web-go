package module

import "moekoe-go/util"

func init() { Register("/top/card/youth", TopCardYouth) }
func TopCardYouth(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "youth/v1/song/single_card_recommend", Method: "POST",
		Data: map[string]interface{}{"tagid": toString(params["tagid"]), "u_info": "", "source_mixsong": ""},
		Params: map[string]interface{}{"card_id": toIntDefault(params, "card_id", 3005), "area_code": 1, "platform": "ops", "module_id": 1, "ver": "v2", "pagesize": toIntDefault(params, "pagesize", 30)},
		EncryptType: "android", Cookie: cookies,
	})
}