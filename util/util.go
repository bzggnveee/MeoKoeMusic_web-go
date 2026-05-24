package util

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString 生成指定长度的随机大写字母数字串
func RandomString(length int) string {
	const keyString = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = keyString[rand.Intn(len(keyString))]
	}
	return string(b)
}

// RandomNumber 生成指定长度的随机数字串
func RandomNumber(length int) string {
	const keyString = "1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = keyString[rand.Intn(len(keyString))]
	}
	return string(b)
}

// ParseCookieString 解析 HTTP Set-Cookie 头，去掉 Domain/path/expires 属性
func ParseCookieString(cookie string) string {
	for _, attr := range []string{"Domain=", "domain=", "path=", "expires="} {
		for {
			idx := findAttr(cookie, attr)
			if idx < 0 {
				break
			}
			end := strings.IndexByte(cookie[idx:], ';')
			if end < 0 {
				cookie = cookie[:idx]
			} else {
				cookie = cookie[:idx] + cookie[idx+end+1:]
			}
		}
	}
	cookie = strings.ReplaceAll(cookie, ";HttpOnly", "")
	cookie = strings.TrimRight(cookie, "; ")
	return strings.TrimSpace(cookie)
}

func findAttr(s, attr string) int {
	lower := strings.ToLower(s)
	return strings.Index(lower, strings.ToLower(attr))
}

// CookieToJSON 将 cookie 字符串转为 map
func CookieToJSON(cookie string) map[string]string {
	if cookie == "" {
		return map[string]string{}
	}
	result := map[string]string{}
	for _, pair := range strings.Split(cookie, ";") {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			result[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return result
}

// GetGUID 生成类似 UUID 的随机设备 GUID
func GetGUID() string {
	e := func() string {
		n := (65536 * (1 + rand.Float64()))
		return fmt.Sprintf("%04x", int(n))[0:4]
	}
	return fmt.Sprintf("%s%s-%s-%s-%s-%s%s%s", e(), e(), e(), e(), e(), e(), e(), e())
}

// CalculateMid 基于 GUID 的 MD5 计算 big integer MID
func CalculateMid(str string) string {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	bigInt := big.NewInt(0)
	bigInt16 := big.NewInt(16)
	for i := 0; i < len(hash); i++ {
		charVal := big.NewInt(0)
		fmt.Sscanf(string(hash[i]), "%x", charVal)
		powerVal := new(big.Int).Exp(bigInt16, big.NewInt(int64(len(hash)-1-i)), nil)
		bigInt.Add(bigInt, new(big.Int).Mul(charVal, powerVal))
	}
	return bigInt.String()
}
