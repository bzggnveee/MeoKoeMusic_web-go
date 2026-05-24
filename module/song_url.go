package module

import (
	"moekoe-go/util"
)

func init() {
	Register("/song/url", SongURL)
}

// SongURL 获取歌曲播放URL
func SongURL(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	isLite := util.IsLite()
	pageID := 151369488
	ppageID := "463467626,350369493,788954147"
	if isLite {
		pageID = 967177915
		ppageID = "356753938,823673182,967485191"
	}

	if v, ok := params["ppage_id"]; ok && isLite {
		ppageID = toString(v)
	}

	quality := "128"
	if v, ok := params["quality"]; ok {
		q := toString(v)
		magicTypes := []string{"piano", "acappella", "subwoofer", "ancient", "dj", "surnay"}
		isMagic := false
		for _, t := range magicTypes {
			if q == t {
				isMagic = true
				break
			}
		}
		if isMagic {
			quality = "magic_" + q
		} else {
			quality = q
		}
	}

	hash := ""
	if v, ok := params["hash"]; ok {
		hash = toString(v)
	}
	hash = toLowerCase(hash)

	albumID := 0
	if v, ok := params["album_id"]; ok {
		albumID = toInt(v)
	}
	albumAudioID := 0
	if v, ok := params["album_audio_id"]; ok {
		albumAudioID = toInt(v)
	}
	freePart := 0
	if v, ok := params["free_part"]; ok {
		if toBool(v) {
			freePart = 1
		}
	}

	dataMap := map[string]interface{}{
		"album_id":       albumID,
		"area_code":      1,
		"hash":           hash,
		"ssa_flag":       "is_fromtrack",
		"version":        11430,
		"page_id":        pageID,
		"quality":        quality,
		"album_audio_id": albumAudioID,
		"behavior":       "play",
		"pid":            411,
		"cmd":            26,
		"pidversion":     3001,
		"IsFreePart":     freePart,
		"ppage_id":       ppageID,
		"cdnBackup":      1,
		"module":         "",
		"clientver":      11430,
	}

	if isLite {
		dataMap["pid"] = 411
	} else {
		dataMap["pid"] = 2
	}

	mergedCookies := make(map[string]string)
	for k, v := range cookies {
		mergedCookies[k] = v
	}
	mergedCookies["dfid"] = util.RandomString(24)

	return requestFn(util.RequestConfig{
		URL:         "/v5/url",
		Method:      "GET",
		Params:      dataMap,
		EncryptType: "android",
		EncryptKey:  true,
		Headers:     map[string]string{"x-router": "trackercdn.kugou.com"},
		Cookie:      mergedCookies,
	})
}
