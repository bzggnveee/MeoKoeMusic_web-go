package module

import (
	"encoding/base64"
	"moekoe-go/util"

	qrcode "github.com/skip2/go-qrcode"
)

func init() { Register("/login/qr/create", LoginQRCreate) }

func LoginQRCreate(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	key := toString(params["key"])
	urlStr := "https://h5.kugou.com/apps/loginQRCode/html/index.html?qrcode=" + key
	base64Img := ""
	if toBool(params["qrimg"]) {
		png, err := qrcode.Encode(urlStr, qrcode.Medium, 256)
		if err == nil {
			base64Img = "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)
		}
	}
	return &util.Response{Status: 200, Body: map[string]interface{}{"code": 200, "data": map[string]interface{}{"url": urlStr, "base64": base64Img}}}, nil
}