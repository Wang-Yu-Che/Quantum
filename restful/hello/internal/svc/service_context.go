// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"Quantum/restful/hello/internal/config"
	"Quantum/service/hello/hello_client"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Hello  hello_client.Hello
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Hello:  hello_client.NewHello(zrpc.MustNewClient(c.Hello)),
	}
}
