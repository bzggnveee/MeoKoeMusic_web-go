package module

import ("moekoe-go/util"; "strings")

func init() { Register("/top/ip", TopIP) }
func TopIP(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	resp, _ := requestFn(util.RequestConfig{
		BaseURL: "http://musicadservice.kugou.com", URL: "/v1/daily_recommend", Method: "POST",
		Data: map[string]interface{}{"tags": map[string]interface{}{}}, Params: map[string]interface{}{"clientver": 12349, "area_code": 1},
		EncryptType: "android", Cookie: cookies,
	})
	if resp != nil && resp.Body != nil {
		if body, ok := resp.Body.(map[string]interface{}); ok {
			if s, _ := body["status"].(float64); s == 1 {
				if data, ok := body["data"].(map[string]interface{}); ok {
					if list, ok := data["list"].([]interface{}); ok {
						for _, item := range list {
							if m, ok := item.(map[string]interface{}); ok {
								if extra, ok := m["extra"].(map[string]interface{}); ok {
									if iu, ok := extra["inner_url"].(string); ok {
										if idx := strings.LastIndex(iu, "ip_id"); idx != -1 {
											extra["ip_id"] = toInt(strings.TrimSpace(iu[idx+6:]))
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return resp, nil
}