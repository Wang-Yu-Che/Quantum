package logic

import (
	"context"

	"Quantum/interface/hello/pb/hello"
	"Quantum/service/hello/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HelloLogic) Hello(in *hello.Request) (*hello.Response, error) {
	msg := "Hello " + in.FirstName + " " + in.LastName
	logx.Debug(msg)

	return &hello.Response{
		Msg: msg,
	}, nil
}
