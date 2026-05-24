package module

import ("moekoe-go/util"; "time")

func init() { Register("/user/listen", UserListen) }
func UserListen(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	ct := time.Now().Unix()
	return requestFn(util.RequestConfig{
		BaseURL:"https://listenservice.kugou.com",URL:"/v2/get_list",Method:"POST",
		Data: map[string]interface{}{"t_userid":getVal(params,cookies,"userid"),"userid":getVal(params,cookies,"userid"),"list_type":toIntDefault(params,"type",0),"area_code":1,"cover":2,"p":util.CryptoRSAEncrypt(map[string]interface{}{"clienttime":ct,"token":getVal(params,cookies,"token")})},
		Params: map[string]interface{}{"clienttime":ct,"plat":0},EncryptType:"android",Cookie:cookies,
	})
}