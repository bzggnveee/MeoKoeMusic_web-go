package module

import "moekoe-go/util"

func init() { Register("/video/privilege", VideoPrivilege) }
func VideoPrivilege(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var res []map[string]interface{}
	for _, h := range splitCSV(toString(params["hash"])) { res = append(res, map[string]interface{}{"hash":h,"id":0,"name":""}) }
	return requestFn(util.RequestConfig{
		URL:"/v1/get_video_privilege",Method:"POST",
		Data: map[string]interface{}{"appid":util.GetAppID(),"area_code":1,"behavior":"play","clientver":util.GetClientVer(),"dfid":dfidVal(params,cookies),"mid":cookies["KUGOU_API_MID"],"resource":res,"token":getVal(params,cookies,"token"),"userid":getValInt(params,cookies,"userid"),"vip":getValInt(params,cookies,"vip_type")},
		EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"media.store.kugou.com"},
	})
}