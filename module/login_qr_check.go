package module

import ("fmt"; "moekoe-go/util")

func init() { Register("/login/qr/check", LoginQRCheck) }

func LoginQRCheck(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	resp, _ := requestFn(util.RequestConfig{
		BaseURL: "https://login-user.kugou.com", URL: "/v2/get_userinfo_qrcode", Method: "GET",
		Params: map[string]interface{}{"plat": 4, "appid": util.GetAppID(), "srcappid": util.SrcAppID, "qrcode": params["key"]},
		EncryptType: "web", Cookie: cookies,
	})
	if resp != nil && resp.Body != nil {
		if body, ok := resp.Body.(map[string]interface{}); ok {
			if data, ok := body["data"].(map[string]interface{}); ok {
				if s, _ := data["status"].(float64); s == 4 {
					// 保存所有登录态 cookie
					saveLoginCookies(resp, data)
				}
			}
		}
	}
	return resp, nil
}

// saveLoginCookies 从 data 中提取所有登录态 cookie
func saveLoginCookies(resp *util.Response, data map[string]interface{}) {
	keys := []string{"token","userid","t1","vip_type","vip_token","dfid","nickname","pic"}
	for _, k := range keys {
		if v := data[k]; v != nil {
			val := toString(v)
			if val != "" {
				if k == "userid" || k == "vip_type" {
					resp.Cookie = append(resp.Cookie, fmt.Sprintf("%s=%.0f", k, toFloat(v)))
				} else {
					resp.Cookie = append(resp.Cookie, fmt.Sprintf("%s=%v", k, v))
				}
			}
		}
	}
}