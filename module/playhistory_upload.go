package module

import ("moekoe-go/util"; "time")

func init() { Register("/playhistory/upload", PlayhistoryUpload) }

func PlayhistoryUpload(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	uid := getValInt(params, cookies, "userid"); tk := getVal(params, cookies, "token")
	ot := toInt(params["time"]); if ot == 0 { ot = int(time.Now().Unix()) }
	pc := toIntDefault(params, "pc", 1)
	return requestFn(util.RequestConfig{
		URL: "/playhistory/v1/upload_songs", Method: "POST",
		Data: map[string]interface{}{"songs": []map[string]interface{}{{"mxid": toInt(params["mxid"]), "op": 1, "ot": ot, "pc": pc}}, "token": tk, "userid": uid},
		Params: map[string]interface{}{"plat": 3}, EncryptType: "android", Cookie: cookies,
	})
}