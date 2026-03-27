// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"Quantum/restful/demo/internal/config"
	"Quantum/service/demo/demo_client"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Demo   demo_client.Demo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Demo:   demo_client.NewDemo(zrpc.MustNewClient(c.Demo)),
	}
}
