package module

import ("encoding/json"; "fmt"; "io"; "moekoe-go/util"; "net/http"; "time")

func init() { Register("/login/openplat", LoginOpenplat) }

func LoginOpenplat(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	wxAppID := util.WxAppID; wxSecret := util.WxSecret
	if util.IsLite() { wxAppID = util.WxLiteAppID; wxSecret = util.WxLiteSecret }
	resp, _ := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?secret=%s&appid=%s&code=%s&grant_type=authorization_code", wxSecret, wxAppID, toString(params["code"])))
	if resp != nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var wxResp map[string]interface{}
		if json.Unmarshal(body, &wxResp) == nil {
			if at, _ := wxResp["access_token"].(string); at != "" {
				if oid, _ := wxResp["openid"].(string); oid != "" {
					dateNow := time.Now().UnixMilli()
					enc := util.CryptoAesEncrypt(`{"access_token":"`+at+`"}`, "", "")
					pk := util.CryptoRSAEncrypt(map[string]interface{}{"clienttime_ms": dateNow, "key": enc.Key})
					liteT2Key := "fd14b35e3f81af3817a20ae7adae7020"; liteT2Iv := "17a20ae7adae7020"
					liteT1Key := "5e4ef500e9597fe004bd09a46d8add98"; liteT1Iv := "04bd09a46d8add98"
					t2 := util.CryptoAesEncrypt(fmt.Sprintf("%s|0f607264fc6318a92b9e13c65db7cd3c|%s|%s|%d", cookies["KUGOU_API_GUID"], cookies["KUGOU_API_MAC"], cookies["KUGOU_API_DEV"], dateNow), liteT2Key, liteT2Iv)
					t1 := util.CryptoAesEncrypt(fmt.Sprintf("|%d", dateNow), liteT1Key, liteT1Iv)
					dm := map[string]interface{}{"dev": cookies["KUGOU_API_DEV"], "force_login": 1, "partnerid": 36, "clienttime_ms": dateNow, "t1": 0, "t2": 0, "t3": "MCwwLDAsMCwwLDAsMCwwLDA=", "openid": oid, "params": enc.Str, "pk": pk}
					if util.IsLite() { dm["t1"] = t1.Str; dm["t2"] = t2.Str }
					resp2, _ := requestFn(util.RequestConfig{
						URL: "/v6/login_by_openplat", Method: "POST", Data: dm, Cookie: cookies, EncryptType: "android",
						Headers: map[string]string{"x-router": "login.user.kugou.com"},
					})
					if resp2 != nil && resp2.Body != nil {
						if body2, ok := resp2.Body.(map[string]interface{}); ok {
							if s, _ := body2["status"].(float64); s == 1 {
								if data, ok := body2["data"].(map[string]interface{}); ok {
									if sp, ok := data["secu_params"]; ok {
										dec := util.CryptoAesDecrypt(toString(sp), enc.Key, "")
										if obj, ok := dec.(map[string]interface{}); ok {
											for k, v := range obj { data[k] = v; resp2.Cookie = append(resp2.Cookie, fmt.Sprintf("%s=%v", k, v)) }
										} else { data["token"] = dec; resp2.Cookie = append(resp2.Cookie, fmt.Sprintf("token=%v", dec)) }
										saveLoginCookies(resp2, data)
									}
								}
							}
						}
					}
					return resp2, nil
				}
			}
		}
	}
	return &util.Response{Status: 502, Body: map[string]interface{}{"status": 0, "msg": "wx auth failed"}}, nil
}