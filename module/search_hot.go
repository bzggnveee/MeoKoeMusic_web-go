package module

import "moekoe-go/util"

func init() { Register("/search/hot", SearchHot) }
func SearchHot(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/api/v3/search/hot_tab", Method: "GET",
		Params: map[string]interface{}{"navid": 1, "plat": 2}, EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "msearch.kugou.com"},
	})
}