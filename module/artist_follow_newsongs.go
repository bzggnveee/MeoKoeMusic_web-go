package module

import "moekoe-go/util"

func init() { Register("/artist/follow/newsongs", ArtistFollowNewsongs) }

func ArtistFollowNewsongs(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	s := 1; if toInt(params["opt_sort"])==2 { s=2 }
	return requestFn(util.RequestConfig{
		URL:"/feed/v1/follow/newsong_album_list",Method:"POST",
		Data:map[string]interface{}{"last_album_id":toInt(params["last_album_id"])},
		Params:map[string]interface{}{"last_album_id":toInt(params["last_album_id"]),"page_size":toIntDefault(params,"pagesize",30),"opt_sort":s},
		EncryptType:"android",Cookie:cookies,
	})
}
