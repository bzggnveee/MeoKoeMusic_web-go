package module

import "moekoe-go/util"

func init() { Register("/youth/channel/amway", YouthChannelAmway) }
func YouthChannelAmway(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL:"/youth/api/amway/v2/index",Method:"GET",Params:map[string]interface{}{"global_collection_id":params["global_collection_id"]},EncryptType:"android",Cookie:cookies})
}