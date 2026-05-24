.PHONY: all build clean dev dist help

# 项目名称
APP_NAME = moekoe-go

# 构建目录
BUILD_DIR = build

# Go 编译参数
GO = go
GOFLAGS = -ldflags="-s -w"
CGO_ENABLED ?= 0

# 源文件
SOURCES = $(shell find . -name '*.go' -not -path './build/*')

all: dist build

help:
	@echo "MoeKoe Music Go 版构建系统"
	@echo ""
	@echo "目标:"
	@echo "  make build       编译 Go 服务端 (单一 ELF)"
	@echo "  make build:prod  CGO_ENABLED=0 静态编译（无 glibc 依赖）"
	@echo "  make dist        构建前端 dist（需要 Node.js）"
	@echo "  make clean       清理构建产物"
	@echo "  make dev         本地开发模式（不嵌入前端）"
	@echo "  make run         编译并运行"
	@echo ""
	@echo "首次构建:"
	@echo "  1. make dist       # 构建前端"
	@echo "  2. make build      # 编译 Go"

# 构建前端（需要 Node.js）
dist:
	@echo "=== 构建前端 ==="
	cd ../MoeKoeMusic && npm run build:docker
	@echo "=== 复制前端文件 ==="
	rm -rf dist
	cp -r ../MoeKoeMusic/dist ./dist
	@echo "=== 前端构建完成 ==="

# 编译 Go 服务端（单一 ELF）
build:
	@echo "=== 编译 Go 服务端 ==="
	$(GO) build $(GOFLAGS) -o $(APP_NAME) .
	@echo "=== 编译完成: $(APP_NAME) ==="
	@ls -lh $(APP_NAME)
	@file $(APP_NAME)

# 静态编译（无 C 依赖，完全静态链接）
build:prod:
	@echo "=== 静态编译 Go 服务端 ==="
	CGO_ENABLED=0 $(GO) build $(GOFLAGS) -a -installsuffix cgo -o $(APP_NAME) .
	@echo "=== 静态编译完成: $(APP_NAME) ==="
	@ls -lh $(APP_NAME)
	@file $(APP_NAME)

# 清理
clean:
	rm -f $(APP_NAME)
	rm -rf dist

# 开发模式（不嵌入前端）
dev:
	@echo "=== 开发模式启动（前端需单独启动 vite） ==="
	$(GO) run -tags nodev .

# 编译并运行
run: build
	./$(APP_NAME) --port=8080 --platform=lite

# 运行测试
test:
	$(GO) test ./...

# 格式化代码
fmt:
	$(GO) fmt ./...

# 代码检查
vet:
	$(GO) vet ./...
