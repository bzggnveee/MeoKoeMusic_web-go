package module

import "moekoe-go/util"

func init() { Register("/album/detail", AlbumDetail) }

func AlbumDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/kmr/v2/albums",Method:"POST",
		Data:map[string]interface{}{"data":[]map[string]interface{}{{"album_id":params["id"]}},"is_buy":toInt(params["is_buy"]),"fields":"album_id,album_name,publish_date,sizable_cover,intro,language,is_publish,heat,type,quality,authors,exclusive,author_name,trans_param"},
		EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"x-router":"openapi.kugou.com","kg-tid":"255"},
	})
}
