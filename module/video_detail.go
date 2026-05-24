package module

import ("moekoe-go/util"; "time")

func init() { Register("/video/detail", VideoDetail) }
func VideoDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	ct := time.Now().Unix()
	var data []map[string]interface{}
	for _, id := range splitCSV(toString(params["id"])) { data = append(data, map[string]interface{}{"video_id":id}) }
	return requestFn(util.RequestConfig{
		URL:"/v1/video",Method:"POST",
		Data: map[string]interface{}{"appid":util.GetAppID(),"clientver":util.GetClientVer(),"clienttime":ct,"mid":cookies["KUGOU_API_MID"],"uuid":util.CryptoMD5(dfidVal(params,cookies)+cookies["KUGOU_API_MID"]),"dfid":dfidVal(params,cookies),"token":getVal(params,cookies,"token"),"key":util.SignParamsKey(ct),"show_resolution":1,"data":data},
		EncryptType:"android",Cookie:cookies,ClearDefaultParams:true,
		Headers:map[string]string{"x-router":"kmr.service.kugou.com"},
	})
}