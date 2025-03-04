package logic

import (
	"Quantum/internal/model/douyin_model/douyin_query"
	"context"

	"Quantum/interface/greet/pb"
	"Quantum/service/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *__.Request) (*__.Response, error) {
	// todo: add your logic here and delete this line
	douyin_query.Comment.WithContext(l.ctx).First()

	return &__.Response{}, nil
}
