package module

import "moekoe-go/util"

func init() { Register("/artist/detail", ArtistDetail) }

func ArtistDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/kmr/v3/author",Method:"POST",
		Data:map[string]interface{}{"author_id":params["id"]},
		EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"openapi.kugou.com","kg-tid":"36"},
	})
}
