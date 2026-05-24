package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"moekoe-go/module"
	"moekoe-go/util"
)

// FrontendDist 嵌入前端构建文件（需在 main.go 中声明 //go:embed）
var FrontendDist embed.FS

// Server 配置
type Server struct {
	Port       int
	Host       string
	Platform   string
	cache      *MemoryCache
	guid       string
	serverDev  string
}

// NewServer 创建新的服务器实例
func NewServer(port int, host, platform, statePath string) *Server {
	if statePath != "" {
		InitState(statePath)
	}
	guid := util.CryptoMD5(util.GetGUID())
	return &Server{
		Port:      port,
		Host:      host,
		Platform:  platform,
		cache:     NewMemoryCache(),
		guid:      guid,
		serverDev: strings.ToUpper(util.RandomString(10)),
	}
}

// Start 启动 HTTP 服务
func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 首页自动同步：浏览器无 token 时注入 state cookie，然后刷新页面
		cookies := parseCookiesFromHeader(r)
		if cookies["token"] == "" {
			if st := GetState(); st != nil {
				for k, v := range st.GetCookies() {
					if v != "" {
						w.Header().Add("Set-Cookie", fmt.Sprintf("%s=%s; PATH=/", k, v))
					}
				}
				log.Printf("[sync] 首页自动注入 %d cookies，刷新页面", len(st.GetCookies()))
				// 等浏览器接收 cookie 后自动刷新
				w.Header().Set("Refresh", "0; url=/")
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.Write([]byte(`<html><body>同步登录中...<script>location.replace("/")</script></body></html>`))
				return
			}
		}
		s.staticHandler(w, r)
	})

	// /tb 同样强制同步（兼容旧入口）
	mux.HandleFunc("/tb", func(w http.ResponseWriter, r *http.Request) {
		if st := GetState(); st != nil {
			for k, v := range st.GetCookies() {
				if v != "" {
					w.Header().Add("Set-Cookie", fmt.Sprintf("%s=%s; PATH=/", k, v))
				}
			}
			log.Printf("[sync] /tb → %d cookies", len(st.GetCookies()))
		}
		data, _ := FrontendDist.ReadFile("dist/index.html")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	})

	// 注册 API 路由
	s.registerAPIRoutes(mux)

	// 注册静态文件路由（前端 dist）

	addr := net.JoinHostPort(s.Host, strconv.Itoa(s.Port))
	log.Printf("MoeKoe Go server running @ http://%s", addr)
	log.Printf("Platform: %s", s.Platform)
	return http.ListenAndServe(addr, s.middleware(mux))
}

// middleware 全局中间件链
func (s *Server) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS
		s.corsMiddleware(w, r)

		// OPTIONS 预检
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Cookie 解析 + 平台标识注入
		s.cookieMiddleware(w, r)

		next.ServeHTTP(w, r)
	})
}

// corsMiddleware CORS 跨域处理
func (s *Server) corsMiddleware(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && !strings.Contains(r.URL.Path, ".") {
		corsOrigin := os.Getenv("CORS_ALLOW_ORIGIN")
		if corsOrigin == "" {
			corsOrigin = r.Header.Get("Origin")
		}
		if corsOrigin == "" {
			corsOrigin = "*"
		}
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", corsOrigin)
		w.Header().Set("Access-Control-Allow-Headers", "Authorization,X-Requested-With,Content-Type,Cache-Control")
		w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
}

// cookieMiddleware Cookie 解析和平台标识注入
func (s *Server) cookieMiddleware(w http.ResponseWriter, r *http.Request) {
	// 解析 Cookie
	cookies := make(map[string]string)
	if cookieHeader := r.Header.Get("Cookie"); cookieHeader != "" {
		for _, pair := range strings.Split(cookieHeader, ";") {
			pair = strings.TrimSpace(pair)
			if pair == "" {
				continue
			}
			idx := strings.IndexByte(pair, '=')
			if idx < 1 || idx == len(pair)-1 {
				continue
			}
			key := strings.TrimSpace(pair[:idx])
			val := strings.TrimSpace(pair[idx+1:])
			cookies[key] = val
		}
	}

	// 注入平台标识 cookie
	isHTTPS := r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https"
	cookieSuffix := "; PATH=/"
	if isHTTPS {
		cookieSuffix = "; PATH=/; SameSite=None; Secure"
	}

	ensureCookie := func(key, value string) {
		if _, ok := cookies[key]; !ok {
			cookies[key] = value
			w.Header().Add("Set-Cookie", fmt.Sprintf("%s=%s%s", key, value, cookieSuffix))
		}
	}

	mid := util.CalculateMid(util.GenerateMD5(s.guid))
	ensureCookie("KUGOU_API_PLATFORM", s.Platform)
	ensureCookie("KUGOU_API_MID", mid)
	ensureCookie("KUGOU_API_GUID", s.guid)
	ensureCookie("KUGOU_API_DEV", s.serverDev)
	ensureCookie("KUGOU_API_MAC", "02:00:00:00:00:00")

	// 存入 context
	r.Header.Set("X-PARSED-COOKIES", serializeCookies(cookies))
}

func parseCookiesFromHeader(r *http.Request) map[string]string {
	c := make(map[string]string)
	if cookieHeader := r.Header.Get("Cookie"); cookieHeader != "" {
		for _, pair := range strings.Split(cookieHeader, ";") {
			pair = strings.TrimSpace(pair)
			if idx := strings.IndexByte(pair, '='); idx > 0 && idx < len(pair)-1 {
				c[strings.TrimSpace(pair[:idx])] = strings.TrimSpace(pair[idx+1:])
			}
		}
	}
	return c
}

func serializeCookies(c map[string]string) string {
	b, _ := json.Marshal(c)
	return string(b)
}

func deserializeCookies(s string) map[string]string {
	c := make(map[string]string)
	json.Unmarshal([]byte(s), &c)
	return c
}

// staticHandler 静态文件服务（前端 dist）
func (s *Server) staticHandler(w http.ResponseWriter, r *http.Request) {
	// API 路由由 registerAPIRoutes 处理，这里只处理静态文件
	path := r.URL.Path

	// 尝试提供静态文件
	filePath := strings.TrimPrefix(path, "/")
	if filePath == "" {
		filePath = "index.html"
	}

	data, err := FrontendDist.ReadFile("dist/" + filePath)
	if err != nil {
		// SPA fallback: 返回 index.html
		data, err = FrontendDist.ReadFile("dist/index.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}
	}

	// 设置 Content-Type
	contentType := getContentType(filePath)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "public, max-age=3600")
	w.Write(data)
}

// registerAPIRoutes 注册所有 API 路由
func (s *Server) registerAPIRoutes(mux *http.ServeMux) {
	// 内建路由：session、settings
	const deviceSecret = "moekoe-sync-2024"

	// GET  /session/status → 查看当前 origin
	mux.HandleFunc("/session/status", func(w http.ResponseWriter, r *http.Request) {
		if st := GetState(); st != nil {
			writeJSON(w, 200, map[string]string{
				"origin":    st.GetOrigin(),
				"device_id": st.DeviceID,
				"created":   st.CreateAt,
			}, nil, nil)
		} else {
			writeJSON(w, 200, map[string]string{"origin": "none"}, nil, nil)
		}
	})

	// GET  /session/sync → 下载同步状态（供 synced 设备拉取）
	mux.HandleFunc("/session/sync", func(w http.ResponseWriter, r *http.Request) {
		if st := GetState(); st != nil {
			writeJSON(w, 200, map[string]interface{}{
				"cookies":   st.GetCookies(),
				"origin":    st.GetOrigin(),
				"device_id": st.DeviceID,
				"created":   st.CreateAt,
			}, nil, nil)
		} else {
			writeJSON(w, 404, map[string]string{"error": "no session"}, nil, nil)
		}
	})

	// POST /session/upload → 仅 real 可上传（X-Origin:real + X-Device-Secret）
	mux.HandleFunc("/session/upload", func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("X-Origin")
		secret := r.Header.Get("X-Device-Secret")

		if origin != "real" || secret != deviceSecret {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]string{"error": "forbidden: only real logins can upload"})
			return
		}

		var body struct {
			Token    string `json:"token"`
			UserID   string `json:"user_id"`
			DeviceID string `json:"device_id"`
			Cookies  map[string]string `json:"cookies"`
		}
		if json.NewDecoder(r.Body).Decode(&body) != nil {
			w.WriteHeader(400)
			return
		}

		if st := GetState(); st != nil {
			if body.Cookies != nil {
				st.Cookies = body.Cookies
			} else {
				st.Cookies = map[string]string{
					"token":  body.Token,
					"userid": body.UserID,
				}
			}
			st.MarkReal(body.DeviceID)
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "origin": "real"})
	})

	mux.HandleFunc("/api/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var body map[string]string
			if json.NewDecoder(r.Body).Decode(&body) == nil {
				for k, v := range body {
					if st := GetState(); st != nil { st.SetSetting(k, v) }
				}
			}
		}
		if st := GetState(); st != nil {
			writeJSON(w, 200, st.Settings, nil, nil)
		} else {
			writeJSON(w, 200, map[string]string{}, nil, nil)
		}
	})
	// 模块路由
	for route, handler := range module.Registry {
		handler := handler // capture
		mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			s.apiHandler(w, r, handler)
		})
		log.Printf("[Route] %s", route)
	}
}

// apiHandler 处理 API 请求
func (s *Server) apiHandler(w http.ResponseWriter, r *http.Request, handler util.HandlerFunc) {
	cookies := deserializeCookies(r.Header.Get("X-PARSED-COOKIES"))

	// 解析 query 参数
	query := r.URL.Query()
	params := make(map[string]interface{})
	for k, v := range query {
		if len(v) == 1 {
			// 特殊处理 cookie 参数
			if k == "cookie" {
				continue
			}
			params[k] = v[0]
		} else {
			params[k] = v
		}
	}

	// 解析 query 中的 cookie
	if cookieStr := query.Get("cookie"); cookieStr != "" {
		cookieFromQuery := util.CookieToJSON(cookieStr)
		for k, v := range cookieFromQuery {
			cookies[k] = v
		}
	}

	// 解析 body（JSON）
	var body map[string]interface{}
	if r.Method == "POST" || r.Method == "PUT" {
		if r.Header.Get("Content-Type") == "application/json" || strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			if err := json.NewDecoder(r.Body).Decode(&body); err == nil {
				if bodyCookie, ok := body["cookie"]; ok {
					if cookieStr, ok := bodyCookie.(string); ok {
						for k, v := range util.CookieToJSON(cookieStr) {
							cookies[k] = v
						}
					}
					delete(body, "cookie")
				}
			}
		}
	}

	// Authorization header → cookie
	if auth := r.Header.Get("Authorization"); auth != "" {
		for k, v := range util.CookieToJSON(auth) {
			cookies[k] = v
		}
	}

	// 合并所有参数
	mergedParams := make(map[string]interface{})
	for k, v := range params {
		mergedParams[k] = v
	}
	if body != nil {
		mergedParams["body"] = body
	}

	// 获取客户端 IP
	clientIP := getClientIP(r)

	// 请求工厂函数
	requestFn := func(config util.RequestConfig) (*util.Response, error) {
		// 合并 cookie
		if config.Cookie == nil {
			config.Cookie = make(map[string]string)
		}
		for k, v := range cookies {
			if _, ok := config.Cookie[k]; !ok {
				config.Cookie[k] = v
			}
		}
		config.RealIP = clientIP
		return util.CreateRequest(config)
	}

	// API 缓存（2分钟）
	cacheKey := r.URL.String()
	if cached, ok := s.cache.Get(cacheKey); ok {
		if resp, ok := cached.(*util.Response); ok && resp.Status == 200 {
			writeJSON(w, resp.Status, resp.Body, resp.Headers, nil)
			return
		}
	}

	// 调用模块 handler
	resp, err := handler(mergedParams, cookies, requestFn)

	if err != nil {
		log.Printf("[ERR] %s: %v", r.URL.String(), err)
		errResp := map[string]interface{}{
			"code": 404,
			"data": nil,
			"msg":  "Not Found",
		}
		writeJSON(w, 404, errResp, nil, nil)
		return
	}

	if resp == nil {
		writeJSON(w, 404, map[string]interface{}{"code": 404, "data": nil, "msg": "Not Found"}, nil, nil)
		return
	}

	// 处理响应 Set-Cookie
	cookieSent := make(map[string]string)
	for _, c := range resp.Cookie {
		// 避免重复
		parts := strings.SplitN(c, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			if _, ok := cookieSent[key]; !ok {
				cookieSent[key] = c
				isHTTPS := r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https"
				if isHTTPS {
					w.Header().Add("Set-Cookie", c+"; PATH=/; SameSite=None; Secure")
				} else {
					w.Header().Add("Set-Cookie", c+"; PATH=/")
				}
			}
		}
	}

	// 持久化登录 cookie
	if len(resp.Cookie) > 0 {
		cm := make(map[string]string)
		for _, c := range resp.Cookie {
			parts := strings.SplitN(c, "=", 2)
			if len(parts) == 2 { cm[parts[0]] = parts[1] }
		}
		if st := GetState(); st != nil {
			st.MergeCookies(cm)
		}
	}

	// 缓存成功的响应（2分钟）
	if resp.Status == 200 {
		s.cache.Set(cacheKey, resp, 2*time.Minute)
	}

	writeJSON(w, resp.Status, resp.Body, resp.Headers, cookies)
	log.Printf("[OK] %s", r.URL.String())
}

// writeJSON 写入 JSON 响应
func writeJSON(w http.ResponseWriter, status int, body interface{}, headers, cookies map[string]string) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	if cookies != nil {
		// cookies 已通过 Set-Cookie 设置，这里不重复
	}
	w.WriteHeader(status)
	b, err := json.Marshal(body)
	if err != nil {
		http.Error(w, `{"status":0,"msg":"json marshal error"}`, http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

// getClientIP 获取客户端真实 IP
func getClientIP(r *http.Request) string {
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	// 去掉 IPv6-mapped IPv4 前缀
	if strings.HasPrefix(host, "::ffff:") {
		host = host[7:]
	}
	return host
}

// getContentType 获取 MIME 类型
func getContentType(path string) string {
	switch {
	case strings.HasSuffix(path, ".html"):
		return "text/html; charset=utf-8"
	case strings.HasSuffix(path, ".css"):
		return "text/css"
	case strings.HasSuffix(path, ".js"):
		return "application/javascript"
	case strings.HasSuffix(path, ".json"):
		return "application/json"
	case strings.HasSuffix(path, ".png"):
		return "image/png"
	case strings.HasSuffix(path, ".jpg"), strings.HasSuffix(path, ".jpeg"):
		return "image/jpeg"
	case strings.HasSuffix(path, ".svg"):
		return "image/svg+xml"
	case strings.HasSuffix(path, ".ico"):
		return "image/x-icon"
	case strings.HasSuffix(path, ".woff"):
		return "font/woff"
	case strings.HasSuffix(path, ".woff2"):
		return "font/woff2"
	case strings.HasSuffix(path, ".ttf"):
		return "font/ttf"
	default:
		return "application/octet-stream"
	}
}

// 确保 embed 被使用
var _ = fs.ReadDir
