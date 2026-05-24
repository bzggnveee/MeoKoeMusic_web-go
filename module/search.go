package module

import (
	"moekoe-go/util"
)

func init() {
	Register("/search", Search)
}

// Search 搜索
func Search(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	keyword := ""
	if v, ok := params["keywords"]; ok {
		keyword = toString(v)
	}

	page := 1
	if v, ok := params["page"]; ok {
		page = toInt(v)
	}

	pagesize := 30
	if v, ok := params["pagesize"]; ok {
		pagesize = toInt(v)
	}

	searchType := "song"
	if v, ok := params["type"]; ok {
		t := toString(v)
		switch t {
		case "special", "lyric", "song", "album", "author", "mv":
			searchType = t
		}
	}

	dataMap := map[string]interface{}{
		"albumhide":    0,
		"iscorrection": 1,
		"keyword":      keyword,
		"nocollect":    0,
		"page":         page,
		"pagesize":     pagesize,
		"platform":     "AndroidFilter",
	}

	apiVersion := "v3"
	if searchType != "song" {
		apiVersion = "v1"
	}

	return requestFn(util.RequestConfig{
		URL:         "/" + apiVersion + "/search/" + searchType,
		Method:      "GET",
		Params:      dataMap,
		EncryptType: "android",
		Headers:     map[string]string{"x-router": "complexsearch.kugou.com"},
		Cookie:      cookies,
	})
}
