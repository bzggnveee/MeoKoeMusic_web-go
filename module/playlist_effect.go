package module

import "moekoe-go/util"

func init() { Register("/playlist/effect", PlaylistEffect) }

func PlaylistEffect(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/pubsongs/v1/get_sound_effect_list", Method: "POST",
		Data: map[string]interface{}{"page": toIntDefault(params, "page", 1), "pagesize": toIntDefault(params, "pagesize", 30)},
		EncryptType: "android", Cookie: cookies,
	})
}