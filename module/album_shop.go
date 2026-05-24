package module

import "moekoe-go/util"

func init() { Register("/album/shop", AlbumShop) }

func AlbumShop(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/zhuanjidata/v3/album_shop_v2/get_classify_data",Method:"GET",
		EncryptType:"android",Cookie:cookies,
	})
}
