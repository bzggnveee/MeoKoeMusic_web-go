package module

import "moekoe-go/util"

func init() { Register("/scene/music", SceneMusic) }
func SceneMusic(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/genesisapi/v1/scene_music/rec_music", Method: "POST",
		Params: map[string]interface{}{"scene_id": params["id"], "page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30)},
		Data: map[string]interface{}{"exposure": []string{}},
		EncryptType: "android", Cookie: cookies,
	})
}