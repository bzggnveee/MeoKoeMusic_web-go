package module

import "moekoe-go/util"

func init() { Register("/everyday/history", EverydayHistory) }

func EverydayHistory(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	pm := map[string]interface{}{"mode": toStringDefault(params, "mode", "list"), "platform": toStringDefault(params, "platform", "ios")}
	if v := toString(params["history_name"]); v != "" { pm["history_name"] = v }
	if v := toString(params["date"]); v != "" { pm["date"] = v }
	return requestFn(util.RequestConfig{
		URL: "/everyday/api/v1/get_history", Method: "POST", Params: pm, EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "everydayrec.service.kugou.com"},
	})
}