package module

import ("moekoe-go/util"; "time")

func init() { Register("/ai/recommend", AIRecommend) }

func AIRecommend(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	clienttime := time.Now().UnixMilli()
	var rs []map[string]interface{}
	if s := toString(params["album_audio_id"]); s != "" {
		for _, id := range splitCSV(s) { rs = append(rs, map[string]interface{}{"ID": toInt(id)}) }
	}
	return requestFn(util.RequestConfig{
		URL: "/recommend", Method: "POST",
		Data: map[string]interface{}{
			"platform":"ios","clientver":util.GetClientVer(),"clienttime":clienttime,
			"userid":getValInt(params,cookies,"userid"),"client_playlist":[]string{},
			"source_type":2,"playlist_ver":2,"area_code":1,"appid":util.GetAppID(),
			"key":util.SignParamsKey(clienttime),"mid":cookies["KUGOU_API_MID"],"recommend_source":rs,
		},
		EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"songlistairec.kugou.com"},
		ClearDefaultParams:true,
	})
}
