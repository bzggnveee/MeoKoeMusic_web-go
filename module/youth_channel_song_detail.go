package module

import "moekoe-go/util"

func init() { Register("/youth/channel/song/detail", YouthChannelSongDetail) }
func YouthChannelSongDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/youth/v2/post/get_song_detail",Method:"GET",
		Params: map[string]interface{}{"global_collection_id":params["global_collection_id"],"fileid":params["fileid"]},
		EncryptType:"android",Cookie:cookies,
	})
}