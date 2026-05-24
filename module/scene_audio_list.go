package module

import "moekoe-go/util"

func init() { Register("/scene/audio/list", SceneAudioList) }
func SceneAudioList(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/scene/v1/scene/audio_list", Method: "POST",
		Data: map[string]interface{}{"appid": util.GetAppID(), "clientver": util.GetClientVer(), "token": getVal(params, cookies, "token"), "userid": getValInt(params, cookies, "userid")},
		Params: map[string]interface{}{"scene_id": params["id"], "module_id": params["module_id"], "tag": params["tag"], "page": toIntDefault(params, "page", 1), "page_size": toIntDefault(params, "pagesize", 30)},
		EncryptType: "android", Cookie: cookies,
	})
}