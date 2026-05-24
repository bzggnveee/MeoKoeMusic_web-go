package module

import "moekoe-go/util"

func init() { Register("/youth/channel/detail", YouthChannelDetail) }
func YouthChannelDetail(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	var data []map[string]interface{}
	for _, id := range splitCSV(toString(params["global_collection_id"])) { data = append(data, map[string]interface{}{"global_collection_id":id}) }
	return requestFn(util.RequestConfig{URL:"/youth/api/channel/v1/channel_list_by_id",Method:"POST",Data:map[string]interface{}{"data":data},EncryptType:"android",Cookie:cookies})
}