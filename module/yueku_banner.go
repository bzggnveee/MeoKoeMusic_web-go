package module

import "moekoe-go/util"

func init() { Register("/yueku/banner", YuekuBanner) }
func YuekuBanner(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/ads.gateway/v3/listen_banner",Method:"POST",
		Data: map[string]interface{}{"plat":0,"channel":201,"operator":7,"networktype":2,"userid":getValInt(params,cookies,"userid"),"vip_type":0,"m_type":0,"tags":[]string{},"apiver":5,"ability":2,"mode":"normal"},
		EncryptType:"android",Cookie:cookies,
	})
}