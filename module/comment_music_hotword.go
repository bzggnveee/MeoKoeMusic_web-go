package module

import "moekoe-go/util"

func init() { Register("/comment/music/hotword", CommentMusicHotword) }

func CommentMusicHotword(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/mcomment/v1/get_hot_word", Method: "POST",
		Params: map[string]interface{}{"mixsongid": params["mixsongid"], "need_show_image": 1, "p": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30), "hot_word": params["hot_word"], "extdata": "0", "code": "fc4be23b4e972707f36b8a828a93ba8a"},
		EncryptType: "android", Cookie: cookies,
	})
}