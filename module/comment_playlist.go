package module

import "moekoe-go/util"

func init() { Register("/comment/playlist", CommentPlaylist) }

func CommentPlaylist(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/m.comment.service/v1/cmtlist", Method: "POST",
		Params: map[string]interface{}{"childrenid": params["id"], "need_show_image": 1, "p": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30), "show_classify": toIntDefault(params, "show_classify", 1), "show_hotword_list": toIntDefault(params, "show_hotword_list", 1), "code": "ca53b96fe5a1d9c22d71c8f522ef7c4f", "content_type": 0, "tag": 5},
		EncryptType: "android", Cookie: cookies,
	})
}