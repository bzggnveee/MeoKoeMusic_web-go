package module

import "moekoe-go/util"

func init() { Register("/user/video/collect", UserVideoCollect) }
func UserVideoCollect(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/collectservice/v2/collect_list_mixvideo",Method:"POST",
		Data: map[string]interface{}{"userid":getVal(params,cookies,"userid"),"token":getVal(params,cookies,"token"),"page":toIntDefault(params,"page",1),"pagesize":toIntDefault(params,"pagesize",30)},
		Params: map[string]interface{}{"plat":1},EncryptType:"android",Cookie:cookies,
	})
}