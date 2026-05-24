package module

import "moekoe-go/util"

func init() { Register("/login/qr/key", LoginQRKey) }

func LoginQRKey(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	a := 1001; if toString(params["type"]) == "web" { a = 1014 }
	return requestFn(util.RequestConfig{
		BaseURL: "https://login-user.kugou.com", URL: "/v2/qrcode", Method: "GET",
		Params: map[string]interface{}{"appid": a, "type": 1, "plat": 4, "qrcode_txt": "https://h5.kugou.com/apps/loginQRCode/html/index.html?appid=" + toString(util.GetAppID()) + "&", "srcappid": util.SrcAppID},
		EncryptType: "web", Cookie: cookies,
	})
}