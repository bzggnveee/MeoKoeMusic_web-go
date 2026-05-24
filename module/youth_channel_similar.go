package module

import "moekoe-go/util"

func init() { Register("/youth/channel/similar", YouthChannelSimilar) }
func YouthChannelSimilar(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/youth/v1/channel/get_friendly_channel",Method:"POST",
		Data: map[string]interface{}{"area_code":1,"playlist_ver":2,"vip_type":getValInt(params,cookies,"vip_type"),"platform":"ios"},
		Params: map[string]interface{}{"channel_id":params["channel_id"]},EncryptType:"android",Cookie:cookies,
	})
}