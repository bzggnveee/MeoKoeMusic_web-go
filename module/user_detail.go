package module

import ("fmt"; "moekoe-go/util"; "time")

func init() { Register("/user/detail", UserDetail) }
func UserDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	ct := time.Now().Unix()
	tk := getVal(params,cookies,"token")
	// 必须保持 token 在前、clienttime 在后的 key 顺序（与 Node.js JSON.stringify 一致）
	pk := util.CryptoRSAEncryptString(fmt.Sprintf(`{"token":"%s","clienttime":%d}`, tk, ct))
	return requestFn(util.RequestConfig{
		URL:"/v3/get_my_info",Method:"POST",
		Data: map[string]interface{}{"visit_time":ct,"usertype":1,"p":pk,"userid":getValInt(params,cookies,"userid")},
		Params: map[string]interface{}{"plat":1},EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"usercenter.kugou.com"},
	})
}