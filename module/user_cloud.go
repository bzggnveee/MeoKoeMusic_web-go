package module

import ("moekoe-go/util"; "strings"; "time")

func init() { Register("/user/cloud", UserCloud) }
func UserCloud(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	uid := getValInt(params, cookies, "userid"); tk := getVal(params, cookies, "token"); ct := time.Now().Unix()
	aesEnc := playlistAesEncrypt(map[string]interface{}{"page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30), "getkmr": 1})
	p := strings.ToUpper(util.RsaEncrypt2(map[string]interface{}{"aes": aesEnc.Key, "uid": uid, "token": tk}))
	resp, _ := requestFn(util.RequestConfig{
		BaseURL: "https://mcloudservice.kugou.com", URL: "/v1/get_list", Method: "POST",
		Params: map[string]interface{}{"clienttime": ct, "mid": cookies["KUGOU_API_MID"], "key": util.SignParamsKey(ct), "clientver": util.GetClientVer(), "appid": util.GetAppID(), "p": p},
		Data: aesEnc.Str, EncryptType: "android", Cookie: cookies, ClearDefaultParams: true, NotSign: true,
	})
	if resp != nil {
		resp.Body = playlistAesDecrypt(toString(resp.Body), aesEnc.Key)
	}
	return resp, nil
}