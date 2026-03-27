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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	rpcRes, err := l.svcCtx.Demo.Login(l.ctx, &demo.UserLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// 2. 将 RPC 的响应结构体 转换为 API 的响应结构体
	return &types.UserLoginResp{
		AccessToken: rpcRes.AccessToken,
		Expire:      rpcRes.Expire,
	}, nil
}
