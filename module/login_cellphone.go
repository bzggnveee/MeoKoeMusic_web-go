package module

import (
	"fmt"
	"moekoe-go/util"
	"time"
)

func init() { Register("/login/cellphone", LoginCellphone) }

const (
	liteT2Key = "fd14b35e3f81af3817a20ae7adae7020"
	liteT2Iv  = "17a20ae7adae7020"
	liteT1Key = "5e4ef500e9597fe004bd09a46d8add98"
	liteT1Iv  = "04bd09a46d8add98"
)

func LoginCellphone(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	dateTime := time.Now().UnixMilli()
	mobile := toString(params["mobile"])
	code := toString(params["code"])
	isLite := util.IsLite()

	encrypt := util.CryptoAesEncrypt(fmt.Sprintf(`{"mobile":"%s","code":"%s"}`, mobile, code), "", "")

	maskedMobile := ""
	if len(mobile) >= 11 {
		maskedMobile = mobile[:2] + "*****" + mobile[10:11]
	}

	dfid := util.RandomString(24)
	if v, ok := cookies["dfid"]; ok {
		dfid = v
	}

	t2 := util.CryptoAesEncrypt(
		fmt.Sprintf("%s|0f607264fc6318a92b9e13c65db7cd3c|%s|%s|%d",
			cookies["KUGOU_API_GUID"],
			toUpper(cookies["KUGOU_API_MAC"]),
			cookies["KUGOU_API_DEV"],
			dateTime),
		liteT2Key, liteT2Iv)

	t1 := util.CryptoAesEncrypt(fmt.Sprintf("|%d", dateTime), liteT1Key, liteT1Iv)

	dataMap := map[string]interface{}{
		"plat": 1, "support_multi": 1, "t1": 0, "t2": 0,
		"clienttime_ms": dateTime, "mobile": maskedMobile,
		"key": util.SignParamsKey(dateTime),
		"pk": toUpper(util.CryptoRSAEncrypt(map[string]interface{}{"clienttime_ms": dateTime, "key": encrypt.Key})),
		"params": encrypt.Str,
	}

	if isLite {
		dataMap["t1"] = t1.Str; dataMap["t2"] = t2.Str
		dataMap["dfid"] = dfid; dataMap["dev"] = cookies["KUGOU_API_DEV"]
		dataMap["gitversion"] = "5f0b7c4"
	} else {
		dataMap["t3"] = "MCwwLDAsMCwwLDAsMCwwLDA="
	}
	if v, ok := params["userid"]; ok {
		dataMap["userid"] = v
	}

	mergedCookies := make(map[string]string)
	for k, v := range cookies {
		mergedCookies[k] = v
	}

	resp, err := requestFn(util.RequestConfig{
		BaseURL: "https://loginserviceretry.kugou.com",
		URL:     "/v7/login_by_verifycode", Method: "POST", Data: dataMap,
		EncryptType: "android",
		Headers:     map[string]string{"support-calm": "1", "User-Agent": "Android16-1070-11440-130-0-LOGIN-wifi"},
		Cookie:      mergedCookies,
	})

	if err != nil || resp == nil {
		return resp, err
	}

	if body, ok := resp.Body.(map[string]interface{}); ok {
		if s, _ := body["status"].(float64); s == 1 {
			if data, ok := body["data"].(map[string]interface{}); ok {
				if sp, ok := data["secu_params"]; ok {
					dec := util.CryptoAesDecrypt(toString(sp), encrypt.Key, "")
					if obj, ok := dec.(map[string]interface{}); ok {
						for k, v := range obj {
							data[k] = v
							resp.Cookie = append(resp.Cookie, fmt.Sprintf("%s=%v", k, v))
						}
					} else if str, ok := dec.(string); ok {
						data["token"] = str
					}
				}
				saveLoginCookies(resp, data)
				resp.Status = 200
			}
		}
	}
	return resp, nil
}

func toFloat(v interface{}) float64 {
	if v == nil { return 0 }
	switch val := v.(type) {
	case float64: return val
	case int: return float64(val)
	case int64: return float64(val)
	default: return 0
	}
}