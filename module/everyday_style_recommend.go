package module

import "moekoe-go/util"

func init() { Register("/everyday/style/recommend", EverydayStyleRecommend) }

func EverydayStyleRecommend(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/everydayrec.service/everyday_style_recommend", Method: "POST", Data: map[string]interface{}{},
		Params: map[string]interface{}{"tagids": toString(params["tagids"])}, EncryptType: "android", Cookie: cookies,
	})
}