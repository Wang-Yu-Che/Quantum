// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package handler

import (
	"net/http"

	g1 "Quantum/restful/demo/internal/handler/g1"
	"Quantum/restful/demo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthInterceptor},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/from/:name",
					Handler: g1.DemoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
	)
}
