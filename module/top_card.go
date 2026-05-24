package module

import ("moekoe-go/util"; "time")

func init() { Register("/top/card", TopCard) }
func TopCard(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dt := time.Now().UnixMilli()
	return requestFn(util.RequestConfig{
		URL: "/singlecardrec.service/v1/single_card_recommend", Method: "POST",
		Data: map[string]interface{}{"appid": util.GetAppID(), "clientver": util.GetClientVer(), "platform": "android", "clienttime": dt, "userid": getValInt(params, cookies, "userid"), "key": util.SignParamsKey(dt), "fakem": "60f7ebf1f812edbac3c63a7310001701760f", "area_code": 1, "mid": cookies["KUGOU_API_MID"], "uuid": "-", "client_playlist": []string{}, "u_info": "a0c35cd40af564444b5584c2754dedec"},
		Params: map[string]interface{}{"card_id": toIntDefault(params, "card_id", 1), "fakem": "60f7ebf1f812edbac3c63a7310001701760f", "area_code": 1, "platform": "ios"},
		EncryptType: "android", Cookie: cookies,
	})
}