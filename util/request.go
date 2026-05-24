package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// CreateRequest 创建并发送 HTTP 请求到酷狗网关
// 自动注入签名、设备标识等参数
func CreateRequest(config RequestConfig) (*Response, error) {
	// 从 cookie 提取设备标识
	dfid := "-"
	token := ""
	userid := 0
	mid := ""
	if config.Cookie != nil {
		if v, ok := config.Cookie["dfid"]; ok && v != "" {
			dfid = v
		}
		if v, ok := config.Cookie["token"]; ok {
			token = v
		}
		if v, ok := config.Cookie["userid"]; ok {
			userid, _ = strconv.Atoi(v)
		}
		if v, ok := config.Cookie["KUGOU_API_MID"]; ok {
			mid = v
		}
	}

	uuid := "-"
	clienttime := time.Now().Unix()

	// 请求头
	headers := map[string]string{
		"dfid":       dfid,
		"clienttime": strconv.FormatInt(clienttime, 10),
		"mid":        mid,
		"kg-rc":      "1",
		"kg-thash":   "5d816a0",
		"kg-rec":     "1",
		"kg-rf":      "B9EDA08A64250DEFFBCADDEE00F8F25F",
	}

	ip := config.RealIP
	if ip != "" {
		headers["X-Real-IP"] = ip
		headers["X-Forwarded-For"] = ip
	}

	// 默认参数
	defaultParams := map[string]interface{}{
		"dfid":       dfid,
		"mid":        mid,
		"uuid":       uuid,
		"appid":      GetAppID(),
		"clientver":  GetClientVer(),
		"clienttime": clienttime,
	}

	if token != "" {
		defaultParams["token"] = token
	}
	if userid != 0 {
		defaultParams["userid"] = userid
	}

	// 合并参数
	var params map[string]interface{}
	if config.ClearDefaultParams {
		params = config.Params
	} else {
		params = make(map[string]interface{})
		for k, v := range defaultParams {
			params[k] = v
		}
		for k, v := range config.Params {
			params[k] = v
		}
	}

	// encryptKey: sign key for song URL
	if config.EncryptKey {
		hash, _ := params["hash"].(string)
		midStr, _ := params["mid"].(string)
		uid, _ := strconv.Atoi(fmt.Sprintf("%v", params["userid"]))
		aid, _ := strconv.Atoi(fmt.Sprintf("%v", params["appid"]))
		params["key"] = SignKey(hash, midStr, uid, aid)
	}

	// 签名
	if _, hasSig := params["signature"]; !hasSig && !config.NotSign {
		switch config.EncryptType {
		case "register":
			params["signature"] = SignatureRegisterParams(params)
		case "web":
			params["signature"] = SignatureWebParams(params)
		default: // android
			dataStr := ""
			if config.Data != nil {
				b, _ := json.Marshal(config.Data)
				dataStr = string(b)
			}
			params["signature"] = SignatureAndroidParams(params, dataStr)
		}
	}

	// 基础 URL
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = GatewayBaseURL
	}

	// 用户代理
	ua := "Android15-1070-11083-46-0-DiscoveryDRADProtocol-wifi"
	if config.Headers != nil {
		if v, ok := config.Headers["User-Agent"]; ok {
			ua = v
		}
	}

	// 构建完整 URL
	fullURL := baseURL + config.URL

	// openapicdn 特殊处理
	if strings.Contains(baseURL, "openapicdn") {
		values := url.Values{}
		for k, v := range params {
			values.Set(k, fmt.Sprintf("%v", v))
		}
		fullURL = fullURL + "?" + values.Encode()
	}

	// 请求体
	var bodyReader io.Reader
	var bodyStr string
	if config.Data != nil {
		switch d := config.Data.(type) {
		case string:
			bodyStr = d
			bodyReader = bytes.NewReader([]byte(d))
		case map[string]interface{}:
			b, _ := json.Marshal(d)
			bodyStr = string(b)
			bodyReader = bytes.NewReader(b)
		default:
			b, _ := json.Marshal(config.Data)
			bodyStr = string(b)
			bodyReader = bytes.NewReader(b)
		}
	}

	// 创建 HTTP 请求
	method := strings.ToUpper(config.Method)
	if method == "" {
		method = "GET"
	}
	req, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return &Response{Status: 502, Body: map[string]interface{}{"msg": err.Error(), "status": 0}}, nil
	}

	// 设置 headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("User-Agent", ua)
	if config.Headers != nil {
		for k, v := range config.Headers {
			if k != "User-Agent" {
				req.Header.Set(k, v)
			}
		}
	}
	if bodyStr != "" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	// 设置查询参数（非 openapicdn）
	if !strings.Contains(baseURL, "openapicdn") {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, fmt.Sprintf("%v", v))
		}
		req.URL.RawQuery = q.Encode()
	}

	// 设置 cookie
	if config.Cookie != nil {
		for k, v := range config.Cookie {
			req.AddCookie(&http.Cookie{Name: k, Value: v})
		}
	}

	// 代理支持
	client := &http.Client{Timeout: 30 * time.Second}
	if proxy := resolveProxyURL(); proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err == nil {
			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
		}
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return &Response{Status: 502, Body: map[string]interface{}{"status": 0, "msg": err.Error()}}, nil
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Response{Status: 502, Body: map[string]interface{}{"status": 0, "msg": err.Error()}}, nil
	}

	// 解析响应 cookie
	var cookies []string
	for _, c := range resp.Cookies() {
		cookies = append(cookies, ParseCookieString(c.String()))
	}
	// 也处理 Set-Cookie 头
	if setCookies := resp.Header["Set-Cookie"]; setCookies != nil {
		for _, c := range setCookies {
			cookies = append(cookies, ParseCookieString(c))
		}
	}

	// 解析响应体
	var body interface{}
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		body = string(bodyBytes)
	}

	// 构建响应头
	respHeaders := map[string]string{}
	if ssaCode := resp.Header.Get("ssa-code"); ssaCode != "" {
		respHeaders["ssa-code"] = ssaCode
	}

	answer := &Response{
		Status:  200,
		Body:    body,
		Cookie:  cookies,
		Headers: respHeaders,
	}

	// 检查业务状态码
	if bodyMap, ok := body.(map[string]interface{}); ok {
		if status, ok := bodyMap["status"]; ok {
			if statusNum, ok := status.(float64); ok && statusNum == 0 {
				answer.Status = 502
			}
		}
		if errCode, ok := bodyMap["error_code"]; ok {
			if errNum, ok := errCode.(float64); ok && errNum != 0 {
				answer.Status = 502
			}
		}
	}

	return answer, nil
}

func resolveProxyURL() string {
	raw := os.Getenv("KUGOU_API_PROXY")
	if raw == "" {
		return ""
	}
	return strings.TrimSpace(raw)
}

// 确保 json 被使用
var _ = json.Indent
