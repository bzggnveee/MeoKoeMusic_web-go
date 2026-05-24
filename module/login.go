package module

import ("fmt"; "moekoe-go/util"; "time")

func init() { Register("/login", Login) }

func Login(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dateNow := time.Now().UnixMilli()
	enc := util.CryptoAesEncrypt(fmt.Sprintf(`{"pwd":"%s","code":"","clienttime_ms":%d}`, toString(params["password"]), dateNow), "", "")
	resp, err := requestFn(util.RequestConfig{
		URL: "/v9/login_by_pwd", Method: "POST",
		Data: map[string]interface{}{
			"plat": 1, "support_multi": 1, "clienttime_ms": dateNow,
			"t1": "562a6f12a6e803453647d16a08f5f0c2ff7eee692cba2ab74cc4c8ab47fc467561a7c6b586ce7dc46a63613b246737c03a1dc8f8d162d8ce1d2c71893d19f1d4b797685a4c6d3d81341cbde65e488c4829a9b4d42ef2df470eb102979fa5adcdd9b4eecfea8b909ff7599abeb49867640f10c3c70fc444effca9d15db44a9a6c907731e2bb0f22cd9b3536380169995693e5f0e2424e3378097d3813186e3fe96bbe7023808a0981b4e2b6135a76faac",
			"t2": "31c4daf4cf480169ccea1cb7d4a209295865a9d2b788510301694db229b87807469ea0d41b4d4b9173c2151da7294aeebfc9738df154bbdf11a4e117bb5dff6a3af8ce5ce333e681c1f29a44038f27567d58992eb81283e080778ac77db1400fdf49b7cf7e26be2e5af4da7830cc3be4",
			"t3": "MCwwLDAsMCwwLDAsMCwwLDA=", "username": params["username"], "params": enc.Str,
			"pk": util.CryptoRSAEncrypt(map[string]interface{}{"clienttime_ms": dateNow, "key": enc.Key}),
		},
		EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "login.user.kugou.com"},
	})
	if err == nil && resp != nil && resp.Body != nil {
		if body, ok := resp.Body.(map[string]interface{}); ok {
			if s, ok := body["status"].(float64); ok && s == 1 {
				if data, ok := body["data"].(map[string]interface{}); ok {
					if sp, ok := data["secu_params"]; ok {
						dec := util.CryptoAesDecrypt(toString(sp), enc.Key, "")
						if obj, ok := dec.(map[string]interface{}); ok {
							for k, v := range obj { data[k] = v; resp.Cookie = append(resp.Cookie, fmt.Sprintf("%s=%v", k, v)) }
						} else { data["token"] = dec; resp.Cookie = append(resp.Cookie, fmt.Sprintf("token=%v", dec)) }
						saveLoginCookies(resp, data)
					}
				}
			}
		}
	}
	return resp, err
}