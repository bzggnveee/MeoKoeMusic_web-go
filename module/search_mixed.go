package module

import ("moekoe-go/util"; "time")

func init() { Register("/search/mixed", SearchMixed) }
func SearchMixed(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	t := time.Now().UnixMilli()
	rid := util.CryptoMD5("bdaa53d04e7475feb9024164a47032f9" + formatInt(t)) + "_0"
	return requestFn(util.RequestConfig{
		URL: "/v3/search/mixed", Method: "GET",
		Params: map[string]interface{}{"ab_tag": 0, "ability": 511, "albumhide": 0, "apiver": 22, "area_code": 1, "clientver": 20125, "cursor": 0, "is_gpay": 0, "iscorrection": 1, "keyword": params["keyword"], "nocollect": 0, "osversion": 16.5, "platform": "IOSFilter", "recver": 2, "req_ai": 1, "requestid": rid, "search_ability": 3, "sec_aggre": 1, "sec_aggre_bitmap": 0, "style_type": 3, "tag": "em"},
		EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "complexsearch.kugou.com", "kg-clienttimems": formatInt(t)},
	})
}
func formatInt(v int64) string { return toString(v) }