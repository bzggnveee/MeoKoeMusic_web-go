package module

import "moekoe-go/util"

func init() { Register("/scene/lists/v2", SceneListsV2) }
func SceneListsV2(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	st := map[string]int{"rec": 1, "hot": 2, "new": 3}
	s := 1; if v, ok := st[toString(params["sort"])]; ok { s = v }
	return requestFn(util.RequestConfig{
		URL: "/scene/v1/scene/list_v2", Method: "POST",
		Params: map[string]interface{}{"scene_id": params["id"], "page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30), "sort_type": s, "kugouid": getVal(params, cookies, "userid")},
		Data: map[string]interface{}{"exposure": []string{}},
		EncryptType: "android", Cookie: cookies,
	})
}