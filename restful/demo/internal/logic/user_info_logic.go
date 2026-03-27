// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"Quantum/interface/demo/pb/demo"
	"context"

	"Quantum/restful/demo/internal/svc"
	"Quantum/restful/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// 1. 调用 RPC
	rpcRes, err := l.svcCtx.Demo.UserInfo(l.ctx, &demo.UserInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	// 2. 转换并返回
	return &types.UserInfoResp{
		UserId:   rpcRes.UserId,
		Username: rpcRes.Username,
		Email:    rpcRes.Email,
	}, nil
}
