package util

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ============================================================
// Signature 签名算法
// ============================================================

// SignatureWebParams Web 版本签名
func SignatureWebParams(params map[string]interface{}) string {
	str := "NVPh5oo715z5DIWAeQlhMDsWXXQV4hwt"
	paramsStr := sortedParams(params, "=", "")
	return CryptoMD5(fmt.Sprintf("%s%s%s", str, paramsStr, str))
}

// SignatureAndroidParams Android 版本签名
func SignatureAndroidParams(params map[string]interface{}, data string) string {
	var str string
	if IsLite() {
		str = "LnT6xpN3khm36zse0QzvmgTZ3waWdRSA"
	} else {
		str = "OIlwieks28dk2k092lksi2UIkp"
	}
	paramsStr := sortedParams(params, "=", "")
	return CryptoMD5(fmt.Sprintf("%s%s%s%s", str, paramsStr, data, str))
}

// SignatureRegisterParams Register 版本签名
func SignatureRegisterParams(params map[string]interface{}) string {
	paramsStr := sortedValues(params)
	return CryptoMD5(fmt.Sprintf("1014%s1014", paramsStr))
}

// SignKey URL sign key 加密
func SignKey(hash, mid string, userid, appid int) string {
	var str string
	if IsLite() {
		str = "185672dd44712f60bb1736df5a377e82"
	} else {
		str = "57ae12eb6890223e355ccfcb74edf70d"
	}

	useAppid := appid
	if useAppid == 0 {
		useAppid = GetAppID()
	}
	useUserid := userid
	if useUserid == 0 {
		useUserid = 0
	}

	return CryptoMD5(fmt.Sprintf("%s%s%d%s%d", hash, str, useAppid, mid, useUserid))
}

// SignCloudKey 云盘 key 加密
func SignCloudKey(hash, pid string) string {
	str := "ebd1ac3134c880bda6a2194537843caa0162e2e7"
	return CryptoMD5(fmt.Sprintf("musicclound%s%s%s", hash, pid, str))
}

// SignParamsKey 登录签名参数 key
func SignParamsKey(data int64) string {
	var str string
	if IsLite() {
		str = "LnT6xpN3khm36zse0QzvmgTZ3waWdRSA"
	} else {
		str = "OIlwieks28dk2k092lksi2UIkp"
	}
	appid := GetAppID()
	clientver := GetClientVer()
	return CryptoMD5(fmt.Sprintf("%d%s%d%d", appid, str, clientver, data))
}

// SignParams 通用签名参数
func SignParams(params map[string]interface{}, data string) string {
	str := "R6snCXJgbCaj9WFRJKefTMIFp0ey6Gza"
	paramsStr := sortedParams(params, "", "")
	return CryptoMD5(fmt.Sprintf("%s%s%s", paramsStr, data, str))
}

// ============================================================
// 排序工具
// ============================================================

// sortedParams 对 params 按 key 排序，拼接为 key=value 或 keyvalue 格式
func sortedParams(params map[string]interface{}, sep, joiner string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var parts []string
	for _, k := range keys {
		v := paramToString(params[k])
		if sep != "" {
			parts = append(parts, fmt.Sprintf("%s%s%s", k, sep, v))
		} else {
			parts = append(parts, fmt.Sprintf("%s%s", k, v))
		}
	}
	if joiner != "" {
		return strings.Join(parts, joiner)
	}
	return strings.Join(parts, "")
}

// sortedValues 对 params 按 value 排序后拼接
func sortedValues(params map[string]interface{}) string {
	values := make([]string, 0, len(params))
	for _, v := range params {
		values = append(values, paramToString(v))
	}
	sort.Strings(values)
	return strings.Join(values, "")
}

func paramToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		if val == float64(int64(val)) {
			return strconv.FormatInt(int64(val), 10)
		}
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		if val {
			return "true"
		}
		return "false"
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

// 确保 os 被使用
var _ = os.Getenv
