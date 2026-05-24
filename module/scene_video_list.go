package module

import "moekoe-go/util"

func init() { Register("/scene/video/list", SceneVideoList) }
func SceneVideoList(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/scene/v1/distribution/video_list", Method: "POST",
		Data: map[string]interface{}{"appid": util.GetAppID(), "clientver": util.GetClientVer(), "token": getVal(params, cookies, "token"), "userid": getValInt(params, cookies, "userid"), "tag_id": params["tag_id"], "page": toIntDefault(params, "page", 1), "page_size": toIntDefault(params, "pagesize", 30), "exposed_data": []string{}},
		EncryptType: "android", Cookie: cookies,
	})
}