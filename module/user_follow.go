package module

import ("moekoe-go/util"; "time")

func init() { Register("/user/follow", UserFollow) }
func UserFollow(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	ct := time.Now().Unix()
	return requestFn(util.RequestConfig{
		URL:"/v4/follow_list",Method:"POST",
		Data: map[string]interface{}{"merge":2,"need_iden_type":1,"ext_params":"k_pic,jumptype,singerid,score","userid":getVal(params,cookies,"userid"),"type":0,"id_type":0,"p":util.CryptoRSAEncrypt(map[string]interface{}{"clienttime":ct,"token":getVal(params,cookies,"token")})},
		Params: map[string]interface{}{"plat":1},EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"relationuser.kugou.com"},
	})
}