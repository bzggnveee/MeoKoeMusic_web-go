package util

// RequestConfig HTTP 请求配置（对应 Node.js createRequest 参数）
type RequestConfig struct {
	BaseURL            string
	URL                string
	Method             string
	Params             map[string]interface{}
	Data               interface{} // map[string]interface{} or string (raw body)
	Headers            map[string]string
	Cookie             map[string]string
	EncryptType        string // "android", "web", "register"
	EncryptKey         bool
	NotSign            bool
	ClearDefaultParams bool
	RealIP             string
	ResponseType       string
}

// Response API 响应结构
type Response struct {
	Status  int
	Body    interface{}
	Cookie  []string
	Headers map[string]string
}

// HandlerFunc API 模块处理函数签名
// params: 合并后的请求参数（query + body + cookie）
// cookies: 请求携带的 cookie
// requestFn: HTTP 请求工厂函数
type HandlerFunc func(params map[string]interface{}, cookies map[string]string, requestFn func(RequestConfig) (*Response, error)) (*Response, error)
