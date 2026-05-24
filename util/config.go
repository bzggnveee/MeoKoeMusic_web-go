package util

import "os"

// 酷狗 API 常量，来自 api/util/config.json
const (
	AppID          = 1005
	ApiVer         = 20
	ClientVer      = 20489
	LiteAppID      = 3116
	LiteClientVer  = 11440
	SrcAppID       = 2919
	WxAppID        = "wx79f2c4418704b4f8"
	WxLiteAppID    = "wx72b795aca60ad321"
	WxSecret       = "4efcab88b700769e376e3f6087b8abc9"
	WxLiteSecret   = "33e486041e5e25729a4e3d2da7502f9a"

	GatewayBaseURL = "https://gateway.kugou.com"
	LoginBaseURL   = "https://loginserviceretry.kugou.com"
)

// IsLite 返回当前是否为概念版平台模式
func IsLite() bool {
	return os.Getenv("platform") == "lite"
}

// GetAppID 根据平台返回对应的 appid
func GetAppID() int {
	if IsLite() {
		return LiteAppID
	}
	return AppID
}

// GetClientVer 根据平台返回对应的 clientver
func GetClientVer() int {
	if IsLite() {
		return LiteClientVer
	}
	return ClientVer
}
