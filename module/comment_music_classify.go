package module

import "moekoe-go/util"

func init() { Register("/comment/music/classify", CommentMusicClassify) }

func CommentMusicClassify(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	s := 1; if toInt(params["sort"]) == 2 { s = 2 }
	return requestFn(util.RequestConfig{
		URL: "/mcomment/v1/cmt_classify_list", Method: "POST",
		Params: map[string]interface{}{"mixsongid": params["mixsongid"], "need_show_image": 1, "page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30), "type_id": params["type_id"], "extdata": "0", "code": "fc4be23b4e972707f36b8a828a93ba8a", "sort_method": s},
		EncryptType: "android", Cookie: cookies,
	})
}