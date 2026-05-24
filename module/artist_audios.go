package module

import ("moekoe-go/util"; "time")

func init() { Register("/artist/audios", ArtistAudios) }

func ArtistAudios(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	ct := time.Now().Unix()
	s := 2; if toString(params["sort"])=="hot" { s=1 }
	return requestFn(util.RequestConfig{
		BaseURL:"https://openapi.kugou.com",URL:"/kmr/v1/audio_group/author",Method:"POST",
		Data:map[string]interface{}{"appid":util.GetAppID(),"clientver":util.GetClientVer(),"mid":cookies["KUGOU_API_MID"],"clienttime":ct,"key":util.SignParamsKey(ct),"author_id":params["id"],"pagesize":toIntDefault(params,"pagesize",30),"page":toIntDefault(params,"page",1),"sort":s,"area_code":"all"},
		EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"openapi.kugou.com","kg-tid":"220"},
	})
}
