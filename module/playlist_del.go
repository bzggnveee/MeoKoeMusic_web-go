package module

import ("moekoe-go/util"; "strings"; "time")

func init() { Register("/playlist/del", PlaylistDel) }

func PlaylistDel(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	userid := getValInt(params, cookies, "userid"); token := getVal(params, cookies, "token"); ct := time.Now().Unix()
	aesEnc := playlistAesEncrypt(map[string]interface{}{"listid": toInt(params["listid"]), "total_ver": 0, "type": 1})
	p := util.RsaEncrypt2(map[string]interface{}{"aes": aesEnc.Key, "uid": userid, "token": token})
	p = strings.ToUpper(p)
	resp, _ := requestFn(util.RequestConfig{
		URL: "/v2/delete_list", Method: "POST",
		Params: map[string]interface{}{"clienttime": ct, "key": util.SignParamsKey(ct), "last_area": "gztx", "clientver": util.GetClientVer(), "appid": util.GetAppID(), "last_time": ct, "p": p},
		Data: aesEnc.Str, EncryptType: "android", Cookie: cookies, Headers: map[string]string{"x-router": "cloudlist.service.kugou.com"},
	})
	if resp != nil {
		resp.Body = playlistAesDecrypt(toString(resp.Body), aesEnc.Key)
	}
	return resp, nil
}