package module

import (
	"moekoe-go/util"
)

func init() {
	Register("/audio/related", AudioRelated)
}

// AudioRelated 相关歌曲
func AudioRelated(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	hash := ""
	if v, ok := params["hash"]; ok {
		hash = toLowerCase(toString(v))
	}

	return requestFn(util.RequestConfig{
		URL:         "/v1/same_song/recommend",
		Method:      "GET",
		Params: map[string]interface{}{
			"hash": hash,
		},
		EncryptType: "android",
		Headers: map[string]string{
			"x-router": "audiorelated.kugou.com",
		},
		Cookie: cookies,
	})
}
