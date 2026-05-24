package module

import (
	"moekoe-go/util"
)

func init() {
	Register("/lyric", Lyric)
}

// Lyric 获取歌词
func Lyric(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	hash := ""
	if v, ok := params["hash"]; ok {
		hash = toLowerCase(toString(v))
	}

	dataMap := map[string]interface{}{
		"hash":         hash,
		"album_id":     0,
		"alt_media":    "QxQ",
		"platformid":   1,
		"tempkey":      "0",
		"lrcfmt":       "lrc",
		"ver":          1,
		"rpt":          "",
		"srctrackhash": "",
	}

	if v, ok := params["album_id"]; ok {
		dataMap["album_id"] = toInt(v)
	}

	return requestFn(util.RequestConfig{
		URL:         "/v1/song/klyric",
		Method:      "GET",
		Params:      dataMap,
		EncryptType: "android",
		Headers:     map[string]string{"x-router": "krcpublish.kugou.com"},
		Cookie:      cookies,
	})
}
