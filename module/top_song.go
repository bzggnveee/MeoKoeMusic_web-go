package module

import (
	"moekoe-go/util"
)

func init() {
	Register("/top/song", TopSong)
}

// TopSong 新歌榜
func TopSong(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:         "/v1/top/song",
		Method:      "GET",
		Params:      map[string]interface{}{},
		EncryptType: "android",
		Headers: map[string]string{
			"x-router": "topsong.kugou.com",
		},
		Cookie: cookies,
	})
}
