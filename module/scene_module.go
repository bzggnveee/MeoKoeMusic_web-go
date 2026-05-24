package module

import "moekoe-go/util"

func init() { Register("/scene/module", SceneModule) }
func SceneModule(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{URL: "/scene/v1/scene/module", Method: "POST", Params: map[string]interface{}{"scene_id": params["id"]}, EncryptType: "android", Cookie: cookies})
}