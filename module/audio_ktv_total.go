package module

import "moekoe-go/util"

func init() { Register("/audio/ktv/total", AudioKtvTotal) }

func AudioKtvTotal(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dm := map[string]interface{}{"isteen":0,"songId":toInt(params["songId"]),"usemkv":1,"platform":2,"singerName":toString(params["singerName"]),"songHash":params["songHash"],"version":12375,"appid":util.GetAppID()}
	dm["sign"] = kugouSign(dm)
	return requestFn(util.RequestConfig{
		BaseURL:"https://acsing.service.kugou.com",URL:"/sing7/listenguide/json/v2/cdn/listenguide/get_total_opus_num_v02.do",
		Params:dm,Method:"GET",EncryptType:"android",Cookie:cookies,ClearDefaultParams:true,NotSign:true,
	})
}
