package module

import "moekoe-go/util"

func init() { Register("/artist/honour", ArtistHonour) }

func ArtistHonour(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		BaseURL:"http://h5activity.kugou.com",URL:"/v1/query_singer_honour_detail",Method:"POST",
		Params:map[string]interface{}{"singer_id":params["id"],"pagesize":toIntDefault(params,"pagesize",30),"page":toIntDefault(params,"page",1)},
		EncryptType:"android",Cookie:cookies,
	})
}
