// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"
	"path"

	"Quantum/restful/hello/internal/config"
	"Quantum/restful/hello/internal/handler"
	"Quantum/restful/hello/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

//go:embed assets/*
var content embed.FS

var configFile = flag.String("f", "etc/hello-dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// ==========================================
	// 1. 原有的首页 (第一个 index)
	// ==========================================
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 这里返回你最原始的那个 index.html
			file, _ := content.ReadFile("assets/index.html")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(file)
		}),
	})

	// ==========================================
	// 2. 新增的 Flow 实验室首页 (第二个 index)
	// ==========================================
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/flow",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 指向 dist 目录下的 index.html
			file, err := content.ReadFile("assets/dist/index.html")
			if err != nil {
				http.Error(w, "FocusFlow index not found", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(file)
		}),
	})

	// ==========================================
	// 其他页面 (贪吃蛇, 佳琪)
	// ==========================================
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/snake",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			file, _ := content.ReadFile("assets/snake.html")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(file)
		}),
	})

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/little-jiaqi",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			file, _ := content.ReadFile("assets/520.html")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(file)
		}),
	})

	// ==========================================
	// 3. 静态资源统一处理中心
	// ==========================================
	// 这里的逻辑是：不管请求来自哪个页面，
	// 只要请求路径是 /assets/xxx，就去 assets/dist/assets/ 找
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/assets/:file",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fileName := path.Base(r.URL.Path)
			// 重点：FocusFlow 打包出来的 JS/CSS 放在 dist/assets 文件夹下
			fullPath := path.Join("assets/dist/assets", fileName)

			file, err := content.ReadFile(fullPath)
			if err != nil {
				// 如果 dist 里找不到，尝试去 assets 根目录找（兼容老页面）
				fullPath = path.Join("assets", fileName)
				file, err = content.ReadFile(fullPath)
				if err != nil {
					http.NotFound(w, r)
					return
				}
			}

			// 手动处理 MIME 类型确保 JS/CSS 正常工作
			ext := path.Ext(fileName)
			switch ext {
			case ".js":
				w.Header().Set("Content-Type", "application/javascript")
			case ".css":
				w.Header().Set("Content-Type", "text/css")
			case ".svg":
				w.Header().Set("Content-Type", "image/svg+xml")
			}
			w.Write(file)
		}),
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
