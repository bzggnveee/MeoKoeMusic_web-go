package module

import "moekoe-go/util"

func init() { Register("/user/vip/detail", UserVipDetail) }
func UserVipDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		BaseURL:"https://kugouvip.kugou.com",URL:"/v1/get_union_vip",Method:"GET",
		Params: map[string]interface{}{"busi_type":"concept"},EncryptType:"android",Cookie:cookies,
	})
}