package module

import "moekoe-go/util"

func init() { Register("/comment/album", CommentAlbum) }

func CommentAlbum(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/m.comment.service/v1/cmtlist",Method:"POST",
		Params:map[string]interface{}{"childrenid":params["id"],"need_show_image":1,"p":toIntDefault(params,"page",1),"pagesize":toIntDefault(params,"pagesize",30),"show_classify":toIntDefault(params,"show_classify",1),"show_hotword_list":toIntDefault(params,"show_hotword_list",1),"code":"94f1792ced1df89aa68a7939eaf2efca"},
		EncryptType:"android",Cookie:cookies,
	})
}
