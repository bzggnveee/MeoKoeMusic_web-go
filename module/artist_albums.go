package module

import "moekoe-go/util"

func init() { Register("/artist/albums", ArtistAlbums) }

func ArtistAlbums(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	s := 1
	if toString(params["sort"])=="hot" { s=3 }
	return requestFn(util.RequestConfig{
		URL:"/kmr/v1/author/albums",Method:"POST",
		Data:map[string]interface{}{"author_id":params["id"],"pagesize":toIntDefault(params,"pagesize",30),"page":toIntDefault(params,"page",1),"sort":s,"category":1,"area_code":"all"},
		EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"openapi.kugou.com","kg-tid":"36"},
	})
}
