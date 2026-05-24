package module

import "moekoe-go/util"

func init() { Register("/youth/user/song", YouthUserSong) }
func YouthUserSong(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/youth/v1/get_user_song_public",Method:"GET",
		Params: map[string]interface{}{"filter_video":0,"type":toIntDefault(params,"type",0),"userid":params["userid"],"pagesize":toIntDefault(params,"pagesize",30),"page":toIntDefault(params,"page",1),"is_filter":0},
		EncryptType:"android",Cookie:cookies,
	})
}