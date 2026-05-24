package module

import (
	"moekoe-go/util"
)

func init() {
	Register("/personal/fm", PersonalFM)
}

// PersonalFM 私人FM
func PersonalFM(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:         "/v1/personal/fm",
		Method:      "GET",
		Params:      map[string]interface{}{},
		EncryptType: "android",
		Headers: map[string]string{
			"x-router": "frequencymachineservice.kugou.com",
		},
		Cookie: cookies,
	})
}
