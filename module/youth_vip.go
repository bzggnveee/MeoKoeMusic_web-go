package module

import ("moekoe-go/util"; "time")

func init() { Register("/youth/vip", YouthVip) }
func YouthVip(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	t := time.Now().UnixMilli()
	return requestFn(util.RequestConfig{
		URL:"/youth/v1/ad/play_report",Method:"POST",
		Data: map[string]interface{}{"ad_id":12307537187,"play_end":t,"play_start":t-30000},
		EncryptType:"android",Cookie:cookies,
	})
}