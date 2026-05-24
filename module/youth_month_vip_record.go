package module

import "moekoe-go/util"

func init() { Register("/youth/month/vip/record", YouthMonthVipRecord) }
func YouthMonthVipRecord(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL:"/youth/v1/activity/get_month_vip_record",Method:"GET",Params:map[string]interface{}{"latest_limit":100},EncryptType:"android",Cookie:cookies})
}