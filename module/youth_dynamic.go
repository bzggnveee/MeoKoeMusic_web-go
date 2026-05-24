package module

import "moekoe-go/util"

func init() { Register("/youth/dynamic", YouthDynamic) }
func YouthDynamic(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL:"/youth/v3/user/get_dynamic",Method:"GET",EncryptType:"android",Cookie:cookies})
}