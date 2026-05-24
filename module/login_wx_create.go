package module

import ("encoding/json"; "io"; "net/http"; "moekoe-go/util"; "fmt")

func init() { Register("/login/wx/create", LoginWxCreate) }

func LoginWxCreate(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	wxAppID := util.WxAppID; wxSecret := util.WxSecret
	if util.IsLite() { wxAppID = util.WxLiteAppID; wxSecret = util.WxLiteSecret }
	resp, _ := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?appid=%s&secret=%s&grant_type=client_credential", wxAppID, wxSecret))
	if resp == nil { return &util.Response{Status:502, Body:map[string]interface{}{"status":0,"msg":"wx request failed"}}, nil }
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var tr map[string]interface{}
	json.Unmarshal(body, &tr)
	if at, _ := tr["access_token"].(string); at != "" {
		resp2, _ := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=2", at))
		if resp2 != nil {
			defer resp2.Body.Close()
			body2, _ := io.ReadAll(resp2.Body)
			var tr2 map[string]interface{}
			json.Unmarshal(body2, &tr2)
			if ec, _ := tr2["errcode"].(float64); ec == 0 {
				if ticket, _ := tr2["ticket"].(string); ticket != "" {
					return &util.Response{Status:200, Body: map[string]interface{}{"ticket": ticket, "appid": wxAppID}}, nil
				}
			}
			return &util.Response{Status:502, Body:map[string]interface{}{"status":0,"msg":tr2}}, nil
		}
	}
	return &util.Response{Status:502, Body:map[string]interface{}{"status":0,"msg":tr}}, nil
}