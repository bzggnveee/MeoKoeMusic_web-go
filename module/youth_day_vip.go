package module

import "moekoe-go/util"

func init() { Register("/youth/day/vip", YouthDayVip) }
func YouthDayVip(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/youth/v1/recharge/receive_vip_listen_song",Method:"POST",
		Params: map[string]interface{}{"source_id":90139,"receive_day":params["receive_day"]},
		EncryptType:"android",Cookie:cookies,
	})
}