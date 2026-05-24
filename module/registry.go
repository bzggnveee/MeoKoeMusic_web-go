package module

import "moekoe-go/util"

// Registry 全局 API 模块注册表
// key 为路由路径（如 "/search", "/user/detail"），与 Node.js 版保持一致
var Registry = make(map[string]util.HandlerFunc)

// Register 注册一个 API 模块
func Register(route string, handler util.HandlerFunc) {
	Registry[route] = handler
}
