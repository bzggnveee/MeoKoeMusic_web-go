package module

import "moekoe-go/util"

func init() { Register("/artist/lists", ArtistLists) }

func ArtistLists(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/ocean/v6/singer/list",Method:"GET",
		Params:map[string]interface{}{"musician":toInt(params["musician"]),"sextype":toString(params["sextypes"]),"showtype":2,"type":toString(params["type"]),"hotsize":toIntDefault(params,"hotsize",30)},
		EncryptType:"android",Cookie:cookies,
	})
}
