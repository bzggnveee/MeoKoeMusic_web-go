package module

import "moekoe-go/util"

func init() { Register("/user/follow/message", UserFollowMessage) }
func UserFollowMessage(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/msg.mobile/v3/msgtag/history",Method:"GET",
		Params: map[string]interface{}{"filter":1,"maxid":0,"pagesize":toIntDefault(params,"pagesize",30),"tag":"chat:"+getVal(params,cookies,"userid")+"_"+toString(params["id"])},
		EncryptType:"android",Cookie:cookies,
	})
}