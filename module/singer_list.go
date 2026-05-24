package module

import "moekoe-go/util"

func init() { Register("/singer/list", SingerList) }
func SingerList(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/ocean/v6/singer/list", Method: "GET",
		Params: map[string]interface{}{"hotsize": toIntDefault(params, "hotsize", 200), "musician": 0, "sextype": toIntDefault(params, "sextype", 0), "showtype": 2, "type": toIntDefault(params, "type", 0)},
		EncryptType: "android", Cookie: cookies,
	})
}