package module

import "moekoe-go/util"

func init() { Register("/favorite/count", FavoriteCount) }

func FavoriteCount(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/count/v1/audio/mget_collect", Method: "GET",
		Params: map[string]interface{}{"mixsongids": params["mixsongids"]}, EncryptType: "android", Cookie: cookies,
	})
}