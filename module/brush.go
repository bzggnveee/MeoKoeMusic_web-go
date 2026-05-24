package module

import ("moekoe-go/util"; "time")

func init() { Register("/brush", Brush) }

func Brush(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	uid := getValInt(params,cookies,"userid"); vt := getValInt(params,cookies,"vip_type"); dt := time.Now().UnixMilli()
	return requestFn(util.RequestConfig{
		URL:"/genesisapi/v1/newepoch_song_rec/feed",Method:"POST",
		Data:map[string]interface{}{
			"behaviors":[]string{},"abtest":map[string]interface{}{"abtest":map[string]interface{}{"shuashua":map[string]interface{}{"commentcard":2}}},
			"personal_recommend_params":map[string]interface{}{"userid":uid,"appid":util.GetAppID(),"playlist_ver":2,"clienttime":dt,"mid":cookies["KUGOU_API_MID"],"new_sync_point":dt,"module_id":1,"action":"login","vip_type":vt,"vip_flags":3,"recommend_source_locked":0,"song_pool_id":toInt(params["song_pool_id"]),"callerid":0,"m_type":1,"kguid":uid,"platform":"ios","area_code":1,"fakem":"ca981cfc583a4c37f28d2d49000013c16a0a","clientver":11850,"mode":toStringDefault(params,"mode","normal"),"active_swtich":"on","key":util.SignParamsKey(dt)},
		},
		Params:map[string]interface{}{"sort_type":1,"platform":"ios","page":1,"content_ver":4,"clientver":11850},
		EncryptType:"android",Cookie:cookies,
	})
}
