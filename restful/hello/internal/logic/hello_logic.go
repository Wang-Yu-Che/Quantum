// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"Quantum/service/hello/hello_client"
	"context"

	"Quantum/restful/hello/internal/svc"
	"Quantum/restful/hello/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req *types.Req) (resp *types.Res, err error) {
	r, err := l.svcCtx.Hello.Hello(l.ctx, &hello_client.Request{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return nil, nil
	}
	logx.Debug(req.Action)

	return &types.Res{
		Message: r.Msg,
	}, nil
}
