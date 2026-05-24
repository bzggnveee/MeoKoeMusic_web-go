package module

import "moekoe-go/util"

func init() { Register("/youth/channel/song", YouthChannelSong) }
func YouthChannelSong(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/youth/api/channel/v1/channel_get_song_audit_passed",Method:"GET",
		Params: map[string]interface{}{"global_collection_id":params["global_collection_id"],"pagesize":toIntDefault(params,"pagesize",30),"page":toIntDefault(params,"page",1),"is_filter":0},
		EncryptType:"android",Cookie:cookies,
	})
}