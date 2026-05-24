package module

import "moekoe-go/util"

func init() { Register("/audio/accompany/matching", AudioAccompanyMatching) }

func AudioAccompanyMatching(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dm := map[string]interface{}{"isteen":0,"mixId":toInt(params["mixId"]),"usemkv":1,"platform":2,"fileName":toString(params["fileName"]),"hash":params["hash"],"version":12375,"appid":util.GetAppID()}
	dm["sign"] = kugouSign(dm)
	return requestFn(util.RequestConfig{
		BaseURL:"https://nsongacsing.kugou.com",URL:"/sing7/accompanywan/json/v2/cdn/optimal_matching_accompany_2_listen.do",
		Params:dm,Method:"GET",EncryptType:"android",Cookie:cookies,ClearDefaultParams:true,NotSign:true,
	})
}
