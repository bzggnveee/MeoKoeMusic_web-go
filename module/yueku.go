package module

import "moekoe-go/util"

func init() { Register("/yueku", Yueku) }
func Yueku(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/v1/yueku/recommend_v2",Method:"GET",
		Params: map[string]interface{}{"operator":7,"plat":0,"type":11,"area_code":1,"req_multi":1},
		EncryptType:"android",Cookie:cookies,Headers:map[string]string{"x-router":"service.mobile.kugou.com"},
	})
}