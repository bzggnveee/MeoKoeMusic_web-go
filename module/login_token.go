package module

import ("fmt"; "moekoe-go/util"; "time")

func init() { Register("/login/token", LoginToken) }

var tokKey = "90b8382a1bb4ccdcf063102053fd75b8"
var tokIv = "f063102053fd75b8"
var tokLiteKey = "c24f74ca2820225badc01946dba4fdf7"
var tokLiteIv = "adc01946dba4fdf7"

func LoginToken(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dateNow := time.Now().UnixMilli()
	token := getVal(params, cookies, "token")
	userid := toString(getVal(params, cookies, "userid"))
	k, iv := tokKey, tokIv
	if util.IsLite() { k, iv = tokLiteKey, tokLiteIv }
	enc := util.CryptoAesEncrypt(fmt.Sprintf(`{"clienttime":%d,"token":"%s"}`, time.Now().Unix(), token), k, iv)
	encParams := util.CryptoAesEncrypt(`{}`, "", "")
	pk := util.CryptoRSAEncrypt(map[string]interface{}{"clienttime_ms": dateNow, "key": encParams.Key})
	liteT2Key := "fd14b35e3f81af3817a20ae7adae7020"; liteT2Iv := "17a20ae7adae7020"
	liteT1Key := "5e4ef500e9597fe004bd09a46d8add98"; liteT1Iv := "04bd09a46d8add98"
	t2 := util.CryptoAesEncrypt(fmt.Sprintf("%s|0f607264fc6318a92b9e13c65db7cd3c|%s|%s|%d", cookies["KUGOU_API_GUID"], cookies["KUGOU_API_MAC"], cookies["KUGOU_API_DEV"], dateNow), liteT2Key, liteT2Iv)
	t1Data := fmt.Sprintf("|%d", dateNow)
	if v := cookies["t1"]; v != "" { t1Data = fmt.Sprintf("%s|%d", v, dateNow) }
	t1 := util.CryptoAesEncrypt(t1Data, liteT1Key, liteT1Iv)
	dm := map[string]interface{}{"dfid": dfidVal(params, cookies), "p3": enc.Str, "plat": 1, "t1": 0, "t2": 0, "t3": "MCwwLDAsMCwwLDAsMCwwLDA=", "pk": pk, "params": encParams.Str, "userid": userid, "clienttime_ms": dateNow}
	if util.IsLite() { dm["t1"] = t1.Str; dm["t2"] = t2.Str; dm["dev"] = cookies["KUGOU_API_DEV"] }
	resp, _ := requestFn(util.RequestConfig{
		BaseURL: "http://login.user.kugou.com", URL: "/v5/login_by_token", Method: "POST",
		Data: dm, Cookie: cookies, EncryptType: "android",
	})
	if resp != nil && resp.Body != nil {
		if body, ok := resp.Body.(map[string]interface{}); ok {
			if s, _ := body["status"].(float64); s == 1 {
				if data, ok := body["data"].(map[string]interface{}); ok {
					if sp, ok := data["secu_params"]; ok {
						dec := util.CryptoAesDecrypt(toString(sp), encParams.Key, "")
						if obj, ok := dec.(map[string]interface{}); ok {
							for k, v := range obj { data[k] = v; resp.Cookie = append(resp.Cookie, fmt.Sprintf("%s=%v", k, v)) }
						} else { data["token"] = dec }
					}
					saveLoginCookies(resp, data)
				}
			}
		}
	}
	return resp, nil
}