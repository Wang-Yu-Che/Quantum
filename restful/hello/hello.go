// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"

	"Quantum/restful/hello/internal/config"
	"Quantum/restful/hello/internal/handler"
	"Quantum/restful/hello/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

//go:embed assets/index.html
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

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 从 embed 中读取并返回 index.html
			file, _ := content.ReadFile("assets/index.html")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(file)
		}),
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
