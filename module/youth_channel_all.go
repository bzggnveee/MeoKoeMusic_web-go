package module

import "moekoe-go/util"

func init() { Register("/youth/channel/all", YouthChannelAll) }
func YouthChannelAll(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL:"/youth/v2/channel/channel_all_list",Method:"GET",Params:map[string]interface{}{"page":toIntDefault(params,"page",1),"pagesize":toIntDefault(params,"pagesize",30),"type":1},EncryptType:"android",Cookie:cookies})
}