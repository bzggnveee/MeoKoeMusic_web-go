package module

import "moekoe-go/util"

func init() { Register("/youth/channel/sub", YouthChannelSub) }
func YouthChannelSub(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	t := 1; if toInt(params["t"]) == 0 { t = 0 }
	m := "post"; if t == 0 { m = "delete" }
	return requestFn(util.RequestConfig{
		URL:"/youth/v1/channel" + ternary(t==0,"_un","") + "_subscribe",Method:m,
		Params: map[string]interface{}{"global_collection_id":params["global_collection_id"],"source":1},
		EncryptType:"android",Cookie:cookies,
	})
}
func ternary(cond bool, a, b string) string { if cond { return a }; return b }