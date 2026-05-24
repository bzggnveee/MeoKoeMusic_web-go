package module

import ("moekoe-go/util"; "strings")

func init() { Register("/register/dev", RegisterDev) }
func RegisterDev(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	uid := getValInt(params, cookies, "userid"); tk := getVal(params, cookies, "token")
	dm := map[string]interface{}{
		"availableRamSize": toIntDefault(params, "availableRamSize", 4983533568),
		"availableRomSize": toIntDefault(params, "availableRomSize", 48114719),
		"availableSDSize": toIntDefault(params, "availableSDSize", 48114717),
		"basebandVer": toStringDefault(params, "basebandVer", ""),
		"batteryLevel": toIntDefault(params, "batteryLevel", 100),
		"batteryStatus": toIntDefault(params, "batteryStatus", 3),
		"brand": toStringDefault(params, "brand", "Redmi"),
		"buildSerial": toStringDefault(params, "buildSerial", "unknown"),
		"device": toStringDefault(params, "device", "marble"),
		"imei": toStringDefault(params, "imei", cookies["KUGOU_API_GUID"]),
		"imsi": toStringDefault(params, "imsi", ""),
		"manufacturer": toStringDefault(params, "manufacturer", "Xiaomi"),
		"uuid": toStringDefault(params, "uuid", cookies["KUGOU_API_GUID"]),
		"accelerometer": toBool(params["accelerometer"]),
		"gravity": toBool(params["gravity"]),
		"gyroscope": toBool(params["gyroscope"]),
		"light": toBool(params["light"]),
		"magnetic": toBool(params["magnetic"]),
		"orientation": toBool(params["orientation"]),
		"pressure": toBool(params["pressure"]),
		"step_counter": toBool(params["step_counter"]),
		"temperature": toBool(params["temperature"]),
	}
	aesEnc := playlistAesEncrypt(dm)
	p := strings.ToUpper(util.RsaEncrypt2(map[string]interface{}{"aes": aesEnc.Key, "uid": uid, "token": tk}))
	resp, _ := requestFn(util.RequestConfig{
		BaseURL: "https://userservice.kugou.com", URL: "/risk/v2/r_register_dev", Method: "POST",
		Data: aesEnc.Str, Params: map[string]interface{}{"part": 1, "platid": 1, "p": p},
		EncryptType: "android", Cookie: cookies,
	})
	if resp != nil {
		resp.Body = playlistAesDecrypt(toString(resp.Body), aesEnc.Key)
		if body, ok := resp.Body.(map[string]interface{}); ok {
			if s, _ := body["status"].(float64); s == 1 {
				if data, ok := body["data"].(map[string]interface{}); ok {
					if dfid, ok := data["dfid"]; ok { resp.Cookie = append(resp.Cookie, "dfid="+toString(dfid)) }
				}
			}
		}
	}
	return resp, nil
}