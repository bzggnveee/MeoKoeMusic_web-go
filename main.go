package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"moekoe-go/module"
	"moekoe-go/server"
)

//go:embed dist/*
var frontendDist embed.FS

func main() {
	port := flag.Int("port", 8080, "服务端口号")
	host := flag.String("host", "", "监听地址（默认所有接口）")
	platform := flag.String("platform", "lite", "平台类型（lite=概念版, 默认）")
	proxy := flag.String("proxy", "", "代理地址（如 http://127.0.0.1:7890）")
	guid := flag.String("guid", "", "设备 GUID")
	dev := flag.String("dev", "", "设备 DEV ID")
	mac := flag.String("mac", "", "设备 MAC 地址")
	cors := flag.String("cors", "", "CORS 允许的 Origin")
	stateFile := flag.String("state", "", "状态文件路径（默认二进制同目录 moekoe-state.json）")
	showHelp := flag.Bool("help", false, "显示帮助")

	flag.Parse()

	if *showHelp {
		fmt.Println("MoeKoe Music Go 版 - 酷狗第三方客户端服务端")
		fmt.Println("")
		fmt.Println("用法:")
		fmt.Println("  ./moekoe-go [选项]")
		fmt.Println("")
		fmt.Println("选项:")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("环境变量:")
		fmt.Println("  PORT                  服务端口号")
		fmt.Println("  platform              平台类型（lite=概念版）")
		fmt.Println("  KUGOU_API_PROXY       代理地址")
		fmt.Println("  KUGOU_API_GUID        设备 GUID")
		fmt.Println("  KUGOU_API_DEV         设备 DEV ID")
		fmt.Println("  KUGOU_API_MAC         设备 MAC 地址")
		fmt.Println("  CORS_ALLOW_ORIGIN     CORS 允许的 Origin")
		return
	}

	// 环境变量覆盖
	if p := os.Getenv("PORT"); p != "" && *port == 8080 {
		fmt.Sscanf(p, "%d", port)
	}
	if p := os.Getenv("platform"); p != "" && *platform == "lite" {
		*platform = p
	}
	if p := os.Getenv("KUGOU_API_PROXY"); p != "" && *proxy == "" {
		*proxy = p
	}
	if p := os.Getenv("KUGOU_API_GUID"); p != "" && *guid == "" {
		*guid = p
	}
	if p := os.Getenv("KUGOU_API_DEV"); p != "" && *dev == "" {
		*dev = p
	}
	if p := os.Getenv("KUGOU_API_MAC"); p != "" && *mac == "" {
		*mac = p
	}
	if p := os.Getenv("CORS_ALLOW_ORIGIN"); p != "" && *cors == "" {
		*cors = p
	}

	// 设置环境变量（供 util 包使用）
	os.Setenv("platform", *platform)
	if *proxy != "" { os.Setenv("KUGOU_API_PROXY", *proxy) }
	if *guid != "" { os.Setenv("KUGOU_API_GUID", *guid) }
	if *dev != "" { os.Setenv("KUGOU_API_DEV", *dev) }
	if *mac != "" { os.Setenv("KUGOU_API_MAC", *mac) }
	if *cors != "" { os.Setenv("CORS_ALLOW_ORIGIN", *cors) }

	// 计算 state 文件路径（默认二进制同目录）
	statePath := *stateFile
	if statePath == "" {
		execPath, err := os.Executable()
		if err == nil {
			statePath = filepath.Join(filepath.Dir(execPath), "moekoe-state.json")
		} else {
			statePath = "moekoe-state.json"
		}
	}

	// 初始化前端嵌入文件系统
	server.FrontendDist = frontendDist

	// 打印启动信息
	log.Printf("========================================")
	log.Printf("  MoeKoe Music Go Server")
	log.Printf("  端口:     %d", *port)
	log.Printf("  平台:     %s", *platform)
	log.Printf("  状态文件: %s", statePath)
	log.Printf("  路由数:   %d", len(module.Registry))
	log.Printf("========================================")

	// 启动服务器
	srv := server.NewServer(*port, *host, *platform, statePath)
	if err := srv.Start(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}