package module

import ("moekoe-go/util"; "time")

func init() { Register("/artist/follow", ArtistFollow) }

func ArtistFollow(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	sid := toInt(params["id"]); tk := getVal(params,cookies,"token"); uid := getValInt(params,cookies,"userid"); ct := time.Now().Unix()
	enc := util.CryptoAesEncrypt(`{"singerid":`+toString(sid)+`,"token":"`+tk+`"}`, "", "")
	return requestFn(util.RequestConfig{
		URL:"/followservice/v3/follow_singer",Method:"POST",
		Data:map[string]interface{}{"plat":0,"userid":uid,"singerid":sid,"source":7,"p":rsaEncrypt2(map[string]interface{}{"clienttime":ct,"key":enc.Key}),"params":enc.Str},
		Params:map[string]interface{}{"clienttime":ct},
		EncryptType:"android",Cookie:cookies,
	})
}
