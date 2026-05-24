package module

import "moekoe-go/util"

func init() { Register("/comment/count", CommentCount) }

func CommentCount(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	pm := map[string]interface{}{"r": "comments/getcommentsnum", "code": "fc4be23b4e972707f36b8a828a93ba8a"}
	if v := toString(params["hash"]); v != "" { pm["hash"] = v }
	if v := toString(params["special_id"]); v != "" { pm["childrenid"] = v }
	return requestFn(util.RequestConfig{
		URL: "/index.php", Method: "GET", Params: pm, EncryptType: "web", Cookie: cookies,
		Headers: map[string]string{"x-router": "sum.comment.service.kugou.com"},
	})
}