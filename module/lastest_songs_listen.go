package module

import "moekoe-go/util"

func init() { Register("/lastest/songs/listen", LastestSongsListen) }

func LastestSongsListen(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/playque/devque/v1/get_latest_songs", Method: "POST",
		Data: map[string]interface{}{"area_code": "1", "sources": []string{"pc", "mobile", "tv", "car"}, "userid": getValInt(params, cookies, "userid"), "ret_info": 1, "token": getVal(params, cookies, "token"), "pagesize": toIntDefault(params, "pagesize", 30)},
		EncryptType: "android", Cookie: cookies,
	})
}