package module

import ("moekoe-go/util"; "time")

func init() { Register("/fm/image", FMImage) }

func FMImage(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dt := time.Now().UnixMilli()
	var fd []map[string]interface{}
	if s := toString(params["fmid"]); s != "" {
		for _, id := range splitCSV(s) { fd = append(fd, map[string]interface{}{"fields": "imgUrl100,imgUrl50", "fmid": id, "fmtype": 2}) }
	}
	dm := map[string]interface{}{"appid": util.GetAppID(), "clienttime": dt, "clientver": util.GetClientVer(), "data": fd, "dfid": dfidVal(params, cookies), "key": util.SignParamsKey(dt), "mid": cookies["KUGOU_API_MID"]}
	if u := getVal(params, cookies, "userid"); u != "" { dm["userid"] = u }
	if t := getVal(params, cookies, "token"); t != "" { dm["token"] = t }
	return requestFn(util.RequestConfig{
		URL: "/v1/fm_info", Method: "POST", Data: dm, EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "fm.service.kugou.com", "Content-Type": "application/json"},
	})
}