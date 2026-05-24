package module

import ("moekoe-go/util"; "time")

func init() { Register("/audio", Audio) }

func Audio(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dt := time.Now().UnixMilli()
	var data []map[string]interface{}
	if s := toString(params["hash"]); s != "" {
		for _, h := range splitCSV(s) { data = append(data, map[string]interface{}{"hash":h,"audio_id":0}) }
	}
	dataMap := map[string]interface{}{"appid":util.GetAppID(),"clienttime":dt,"clientver":util.GetClientVer(),"data":data,"dfid":dfidVal(params,cookies),"key":util.SignParamsKey(dt),"mid":cookies["KUGOU_API_MID"]}
	if t := getVal(params,cookies,"token"); t != "" { dataMap["token"] = t }
	if u := getValInt(params,cookies,"userid"); u != 0 { dataMap["userid"] = u }
	return requestFn(util.RequestConfig{
		BaseURL:"http://kmr.service.kugou.com",URL:"/v1/audio/audio",Method:"POST",
		Data:dataMap,EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"kmr.service.kugou.com","Content-Type":"application/json"},
	})
}
