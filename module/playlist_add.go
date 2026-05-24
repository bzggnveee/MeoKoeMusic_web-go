package module

import ("moekoe-go/util"; "time")

func init() { Register("/playlist/add", PlaylistAdd) }

func PlaylistAdd(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	uid := getValInt(params, cookies, "userid"); tk := getVal(params, cookies, "token"); ct := time.Now().Unix()
	dm := map[string]interface{}{"userid": uid, "token": tk, "total_ver": 0, "name": params["name"], "type": toIntDefault(params, "type", 0), "source": toIntDefault(params, "source", 1), "is_pri": 0, "list_create_userid": params["list_create_userid"], "list_create_listid": params["list_create_listid"], "list_create_gid": toStringDefault(params, "list_create_gid", ""), "from_shupinmv": 0}
	if toInt(params["type"]) == 0 { dm["is_pri"] = toIntDefault(params, "is_pri", 0) }
	extra := map[string]interface{}{}
	if toInt(params["type"]) == 0 { extra = map[string]interface{}{"last_time": ct, "last_area": "gztx", "userid": uid, "token": tk} }
	return requestFn(util.RequestConfig{
		URL: "/cloudlist.service/v5/add_list", Method: "POST", Data: dm, Params: extra, EncryptType: "android", Cookie: cookies,
	})
}