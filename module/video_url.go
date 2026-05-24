package module

import "moekoe-go/util"

func init() { Register("/video/url", VideoURL) }
func VideoURL(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/v2/interface/index",Method:"GET",
		Params: map[string]interface{}{"backupdomain":1,"cmd":123,"ext":"mp4","ismp3":0,"hash":params["hash"],"pid":1,"type":1},
		EncryptType:"android",EncryptKey:true,Cookie:cookies,
		Headers:map[string]string{"x-router":"trackermv.kugou.com"},
	})
}