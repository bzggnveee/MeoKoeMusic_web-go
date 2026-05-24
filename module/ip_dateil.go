package module

import "moekoe-go/util"

func init() { Register("/ip/dateil", IPDetail) }

func IPDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var data []map[string]interface{}
	if s := toString(params["id"]); s != "" {
		for _, id := range splitCSV(s) { data = append(data, map[string]interface{}{"ip_id": id}) }
	}
	return requestFn(util.RequestConfig{
		URL: "/openapi/v1/ip", Method: "POST", Data: map[string]interface{}{"data": data, "is_publish": 1},
		EncryptType: "android", Cookie: cookies,
	})
}