package module

import "moekoe-go/util"

func init() { Register("/user/video/love", UserVideoLove) }
func UserVideoLove(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/m.comment.service/v1/get_user_like_video",Method:"GET",
		Params: map[string]interface{}{"kugouid":getVal(params,cookies,"userid"),"pagesize":toIntDefault(params,"pagesize",30),"load_video_info":1,"p":1,"plat":1},
		EncryptType:"android",Cookie:cookies,
	})
}