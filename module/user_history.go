package module

import "moekoe-go/util"

func init() { Register("/user/history", UserHistory) }
func UserHistory(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dm := map[string]interface{}{"token":getVal(params,cookies,"token"),"userid":getValInt(params,cookies,"userid"),"source_classify":"app","to_subdivide_sr":1}
	if v := toString(params["bp"]); v != "" { dm["bp"] = v }
	return requestFn(util.RequestConfig{URL:"/playhistory/v1/get_songs",Method:"POST",Data:dm,EncryptType:"android",Cookie:cookies})
}