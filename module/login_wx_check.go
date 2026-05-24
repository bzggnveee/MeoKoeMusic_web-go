package module

import ("encoding/json"; "io"; "net/http"; "fmt"; "moekoe-go/util")

func init() { Register("/login/wx/check", LoginWxCheck) }

func LoginWxCheck(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	resp, _ := http.Get(fmt.Sprintf("https://long.open.weixin.qq.com/connect/l/qrconnect?f=json&uuid=%s", toString(params["uuid"])))
	if resp == nil { return &util.Response{Status:502, Body:map[string]interface{}{"status":0,"msg":"wx request failed"}}, nil }
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	return &util.Response{Status:200, Body: data}, nil
}