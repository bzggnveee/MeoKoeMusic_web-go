package module

import ("moekoe-go/util"; "net/url"; "strings")

func init() { Register("/ip/zone", IPZone) }

func IPZone(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	resp, _ := requestFn(util.RequestConfig{
		URL: "/v1/zone/index", Method: "GET", EncryptType: "android", Cookie: cookies,
		Headers: map[string]string{"x-router": "yuekucategory.kugou.com"},
	})
	if resp != nil && resp.Body != nil {
		if body, ok := resp.Body.(map[string]interface{}); ok {
			if status, ok := body["status"].(float64); ok && status == 1 {
				if data, ok := body["data"].(map[string]interface{}); ok {
					if list, ok := data["list"].([]interface{}); ok {
						for _, item := range list {
							if m, ok := item.(map[string]interface{}); ok {
								if sl, ok := m["special_link"].(string); ok {
									if u, err := url.Parse(sl); err == nil {
										if path := u.Query().Get("path"); path != "" {
											if u2, err := url.Parse(path); err == nil {
												m["ip_id"] = toInt(strings.TrimSpace(u2.Query().Get("ip_id")))
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
	}
	return resp, nil
}