package module

import "moekoe-go/util"

func init() { Register("/youth/day/vip/upgrade", YouthDayVipUpgrade) }
func YouthDayVipUpgrade(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/youth/v1/listen_song/upgrade_vip_reward",Method:"POST",
		Params: map[string]interface{}{"kugouid":getValInt(params,cookies,"userid"),"ad_type":1},
		EncryptType:"android",Cookie:cookies,
	})
}