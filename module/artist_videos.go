package module

import "moekoe-go/util"

var avTags = map[string]string{"official":"18","live":"20","fan":"23","artist":"42419"}

func init() { Register("/artist/videos", ArtistVideos) }

func ArtistVideos(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	tag := avTags[toString(params["tag"])]
	return requestFn(util.RequestConfig{
		BaseURL:"https://openapicdn.kugou.com",URL:"/kmr/v1/author/videos",Method:"GET",
		Params:map[string]interface{}{"author_id":params["id"],"is_fanmade":"","tag_idx":tag,"pagesize":toIntDefault(params,"pagesize",30),"page":toIntDefault(params,"page",1)},
		EncryptType:"android",Cookie:cookies,
	})
}
