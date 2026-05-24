package module

import "moekoe-go/util"

func init() { Register("/captcha/sent", CaptchaSent) }

func CaptchaSent(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		BaseURL:"http://login.user.kugou.com",URL:"/v7/send_mobile_code",Method:"POST",
		Data:map[string]interface{}{"businessid":5,"mobile":toString(params["mobile"]),"plat":3},
		EncryptType:"android",Cookie:map[string]string{"mid":cookies["KUGOU_API_MID"]},
	})
}
