package logic

import (
	"context"
	"errors"
	"strconv"

	"Quantum/interface/demo/pb/demo"
	"Quantum/service/demo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *UserInfoLogic) UserInfo(in *demo.UserInfoReq) (*demo.UserInfoResp, error) {
	// Mock 逻辑：根据 UserId 返回对应的模拟数据
	// 实际开发中这里会调用 Model 层查询数据库
	if in.UserId <= 0 {
		return nil, errors.New("无效的用户ID")
	}

	return &demo.UserInfoResp{
		UserId:   in.UserId,
		Username: "GoZeroUser_" + strconv.FormatInt(in.UserId, 10),
		Email:    "mock_user@example.com",
	}, nil
}
