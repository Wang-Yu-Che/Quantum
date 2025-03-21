// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: greet.proto

package server

import (
	"context"

	"Quantum/interface/greet/pb"
	"Quantum/service/greet/internal/logic"
	"Quantum/service/greet/internal/svc"
)

type GreetServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedGreetServer
}

func NewGreetServer(svcCtx *svc.ServiceContext) *GreetServer {
	return &GreetServer{
		svcCtx: svcCtx,
	}
}

func (s *GreetServer) Ping(ctx context.Context, in *__.Request) (*__.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
