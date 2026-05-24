package module

import "moekoe-go/util"

func init() { Register("/album/songs", AlbumSongs) }

func AlbumSongs(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/v1/album_audio/lite",Method:"POST",
		Data:map[string]interface{}{"album_id":params["id"],"is_buy":toString(params["is_buy"]),"page":toIntDefault(params,"page",1),"pagesize":toIntDefault(params,"pagesize",30)},
		EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"openapi.kugou.com","kg-tid":"255"},
	})
}
