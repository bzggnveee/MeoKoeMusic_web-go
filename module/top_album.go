package module

import "moekoe-go/util"

func init() { Register("/top/album", TopAlbum) }
func TopAlbum(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/musicadservice/v1/mobile_newalbum_sp", Method: "POST",
		Data: map[string]interface{}{"apiver": util.ApiVer, "token": getVal(params, cookies, "token"), "page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30), "withpriv": 1},
		EncryptType: "android", Cookie: cookies,
	})
}