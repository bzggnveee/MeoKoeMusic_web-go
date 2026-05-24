package module

import "moekoe-go/util"

func init() { Register("/comment/floor", CommentFloor) }

func CommentFloor(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	rt := toString(params["resource_type"]); if rt == "" { rt = toString(params["resourceType"]) }
	rt = toLowerCase(rt)
	code := toString(params["code"])
	if code == "" {
		switch rt {
		case "playlist": code = "ca53b96fe5a1d9c22d71c8f522ef7c4f"
		case "album": code = "94f1792ced1df89aa68a7939eaf2efca"
		default: code = "fc4be23b4e972707f36b8a828a93ba8a"
		}
	}
	useSvc := rt == "playlist" || rt == "album" || code == "ca53b96fe5a1d9c22d71c8f522ef7c4f" || code == "94f1792ced1df89aa68a7939eaf2efca"
	u := "/mcomment/v1/hot_replylist"
	if useSvc { u = "/m.comment.service/v1/hot_replylist" }
	pm := map[string]interface{}{"childrenid": params["special_id"], "need_show_image": 1, "p": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30), "show_classify": toIntDefault(params, "show_classify", 1), "show_hotword_list": toIntDefault(params, "show_hotword_list", 1), "code": code}
	if v := toString(params["tid"]); v != "" { pm["tid"] = v }
	if v := toString(params["mixsongid"]); v != "" { pm["mixsongid"] = v }
	return requestFn(util.RequestConfig{URL: u, Method: "POST", Params: pm, EncryptType: "android", Cookie: cookies})
}