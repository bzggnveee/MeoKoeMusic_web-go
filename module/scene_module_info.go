package module

import "moekoe-go/util"

func init() { Register("/scene/module/info", SceneModuleInfo) }
func SceneModuleInfo(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL: "/scene/v1/scene/module_info", Method: "GET", Params: map[string]interface{}{"scene_id": params["id"], "module_id": params["module_id"]}, EncryptType: "android", Cookie: cookies})
}