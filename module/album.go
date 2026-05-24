package module

import ("moekoe-go/util"; "time")

func init() { Register("/album", Album) }

func Album(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dateTime := time.Now().UnixMilli()
	var data []map[string]interface{}
	if s := toString(params["album_id"]); s != "" {
		for _, id := range splitCSV(s) { data = append(data, map[string]interface{}{"album_id":id,"album_name":"","author_name":""}) }
	}
	dataMap := map[string]interface{}{
		"appid":util.GetAppID(),"clienttime":dateTime,"clientver":util.GetClientVer(),"data":data,
		"dfid":dfidVal(params,cookies),"fields":toString(params["fields"]),
		"key":util.SignParamsKey(dateTime),"mid":cookies["KUGOU_API_MID"],
	}
	if t := getVal(params,cookies,"token"); t != "" { dataMap["token"] = t }
	if u := getValInt(params,cookies,"userid"); u != 0 { dataMap["userid"] = u }
	return requestFn(util.RequestConfig{
		BaseURL:"http://kmr.service.kugou.com",URL:"/v1/album",Method:"POST",
		Data:dataMap,EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"kmr.service.kugou.com","Content-Type":"application/json"},
	})
}
